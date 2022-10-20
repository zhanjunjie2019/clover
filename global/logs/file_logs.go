package logs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func InitLogger() error {
	serverConfig := confs.GetServerConfig()
	logCore := getEncoderCore(serverConfig.LogConf.TransportLevel(), serverConfig)
	logger := zap.New(logCore)
	defs.SetLogger(logger)
	return nil
}

func getEncoder(level zapcore.LevelEncoder) zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    level,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	})
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006-01-02T15:04:05.000"))
}

func getEncoderCore(level zapcore.Level, serverConfig confs.ServerConfig) zapcore.Core {
	writer, err := getWriteSyncer(serverConfig)
	errs.Panic(err)
	return zapcore.NewCore(getEncoder(zapcore.CapitalLevelEncoder), writer, level)
}

func getWriteSyncer(serverConfig confs.ServerConfig) (zapcore.WriteSyncer, error) {
	logConf := serverConfig.LogConf
	svcConf := serverConfig.SvcConf
	fn := fmt.Sprintf("%s-%d.log", svcConf.SvcName, svcConf.SvcNum)
	fileWriter, err := rotatelogs.New(
		path.Join(logConf.Director, "%Y-%m-%d-"+fn),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(logConf.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if logConf.LogInConsole == 1 {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
