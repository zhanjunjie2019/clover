package defs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func SetLogger(l *zap.Logger) {
	logger = l
}

func NewLogLayout(
	level zapcore.Level,
	svcMode uint8,
	svcName string,
	svcNum uint8,
	svcVersion string,
) *LogLayout {
	return &LogLayout{
		level: level,
		logFields: []zap.Field{
			zap.Uint8("svcMode", svcMode),
			zap.String("svcName", svcName),
			zap.Uint8("svcNum", svcNum),
			zap.String("svcVersion", svcVersion),
		},
	}
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

func (l *LogLayout) Warn(msg string, logFields ...zap.Field) {
	l.level = zapcore.WarnLevel
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
