package ilogger

import (
	"ilogger/types"
)

type DefaultLogger struct {
	f Filter

	formatter Formatter
	logger    Outputer
}

func NewDefaultLogger(
	minimumLevel types.LogLevel,
	outputer Outputer,
	formatter Formatter,
) *DefaultLogger {
	return &DefaultLogger{
		f:         Filter{currentLevel: minimumLevel},
		logger:    outputer,
		formatter: formatter,
	}
}

func (l *DefaultLogger) Debug(args ...interface{}) {
	l.f.filter(types.LevelDebug, func() {
		l.logger.Log(l.formatter.Format(types.LevelDebug, args...))
	})
}
func (l *DefaultLogger) Info(args ...interface{}) {
	l.f.filter(types.LevelInfo, func() {
		l.logger.Log(l.formatter.Format(types.LevelInfo, args...))
	})
}
func (l *DefaultLogger) Warning(args ...interface{}) {
	l.f.filter(types.LevelWarning, func() {
		l.logger.Log(l.formatter.Format(types.LevelWarning, args...))
	})
}
func (l *DefaultLogger) Error(args ...interface{}) {
	l.f.filter(types.LevelError, func() {
		l.logger.Log(l.formatter.Format(types.LevelError, args...))
	})
}
func (l *DefaultLogger) Critical(args ...interface{}) {
	l.f.filter(types.LevelCritical, func() {
		l.logger.Log(l.formatter.Format(types.LevelCritical, args...))
	})
}
func (l *DefaultLogger) Fixed(args ...interface{}) {
	l.f.filter(types.LevelFixed, func() {
		l.logger.Log(l.formatter.Format(types.LevelFixed, args...))
	})
}
