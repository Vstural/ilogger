package ilogger

import (
	"ilogger/formatter"
	"ilogger/types"
)

// type Logger interface {
// 	Debug(args... interface{})
// 	Info(args... interface{})
// 	Warning(args... interface{})
// 	Error(args... interface{})
// 	Critical(args... interface{})
// 	Fixed(args... interface{})
// }

type Outputer interface {
	Log(s string)
}

type Logger struct {
	f Filter

	formatter formatter.Formatter
	logger    Outputer
}

func NewLogger(
	level types.LogLevel,
	outputer Outputer,
	formatter formatter.Formatter,
) *Logger {
	return &Logger{
		f:         Filter{currentLevel: level},
		logger:    outputer,
		formatter: formatter,
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.f.filter(types.LevelDebug, func() {
		l.logger.Log(l.formatter.Format(types.LevelDebug, args))
	})
}
func (l *Logger) Info(args ...interface{}) {
	l.f.filter(types.LevelInfo, func() {
		l.logger.Log(l.formatter.Format(types.LevelInfo, args))
	})
}
func (l *Logger) Warning(args ...interface{}) {
	l.f.filter(types.LevelWarning, func() {
		l.logger.Log(l.formatter.Format(types.LevelWarning, args))
	})
}
func (l *Logger) Error(args ...interface{}) {
	l.f.filter(types.LevelError, func() {
		l.logger.Log(l.formatter.Format(types.LevelError, args))
	})
}
func (l *Logger) Critical(args ...interface{}) {
	l.f.filter(types.LevelCritical, func() {
		l.logger.Log(l.formatter.Format(types.LevelCritical, args))
	})
}
func (l *Logger) Fixed(args ...interface{}) {
	l.f.filter(types.LevelFixed, func() {
		l.logger.Log(l.formatter.Format(types.LevelFixed, args))
	})
}

// a filter to determine is log should be output
type Filter struct {
	currentLevel types.LogLevel
}

func (f Filter) filter(level types.LogLevel, lo func()) {
	if level < f.currentLevel {
		return
	}
	lo()
}
