package defs

import (
	"github.com/zhanjunjie2019/clover/global/confs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func SetLogger(l *zap.Logger) {
	logger = l
}

func NewLogLayout(level zapcore.Level) *LogLayout {
	return &LogLayout{level: level}
}

func ErrorLog(msg string, logFields ...zap.Field) {
	layout := NewLogLayout(zapcore.ErrorLevel)
	layout.msg = msg
	layout.logFields = logFields
	layout.Println()
}

func InfoLog(msg string, logFields ...zap.Field) {
	layout := NewLogLayout(zapcore.InfoLevel)
	layout.msg = msg
	layout.logFields = logFields
	layout.Println()
}

// LogLayout 日志layout
type LogLayout struct {
	level     zapcore.Level
	msg       string
	logFields []zap.Field
}

func (l *LogLayout) Error(msg string, logFields ...zap.Field) {
	l.level = zapcore.ErrorLevel
	l.msg = msg
	l.logFields = append(l.logFields, logFields...)
}

func (l *LogLayout) Info(msg string, logFields ...zap.Field) {
	l.level = zapcore.InfoLevel
	l.msg = msg
	l.logFields = append(l.logFields, logFields...)
}

// Println 打印日志
func (l LogLayout) Println() {
	serverConfig := confs.GetServerConfig().SvcConf
	l.logFields = append(l.logFields,
		zap.String("svcName", serverConfig.SvcName),
		zap.String("svcVersion", serverConfig.SvcVersion),
		zap.Uint8("svcNum", serverConfig.SvcNum),
	)
	switch l.level {
	case zapcore.DebugLevel:
		logger.Debug(l.msg, l.logFields...)
	case zapcore.InfoLevel:
		logger.Info(l.msg, l.logFields...)
	case zapcore.WarnLevel:
		logger.Warn(l.msg, l.logFields...)
	case zapcore.ErrorLevel:
		logger.Error(l.msg, l.logFields...)
	case zapcore.DPanicLevel:
		logger.DPanic(l.msg, l.logFields...)
	case zapcore.PanicLevel:
		logger.Panic(l.msg, l.logFields...)
	case zapcore.FatalLevel:
		logger.Fatal(l.msg, l.logFields...)
	default:
		logger.Info(l.msg, l.logFields...)
	}
}

func (l *LogLayout) AppendLogsFields(logFields ...zap.Field) {
	l.logFields = append(l.logFields, logFields...)
}
