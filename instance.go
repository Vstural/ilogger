package ilogger

import (
	"ilogger/formatter"
	"ilogger/logger"
	"ilogger/types"
	"sync"
)

var mutex = &sync.RWMutex{}
var globalLogger *Logger

func init() {
	globalLogger = NewLogger(
		types.LevelDebug,
		logger.NewSTDLogger(),
		formatter.NewDefaultFormatter())
}

func GetLogger() *Logger {
	mutex.RLock()
	defer mutex.RUnlock()
	return globalLogger
}

func Debug(args ...interface{})    { GetLogger().Debug(args...) }
func Info(args ...interface{})     { GetLogger().Info(args...) }
func Warning(args ...interface{})  { GetLogger().Warning(args...) }
func Error(args ...interface{})    { GetLogger().Error(args...) }
func Critical(args ...interface{}) { GetLogger().Critical(args...) }
func Fixed(args ...interface{})    { GetLogger().Fixed(args...) }

func SetLevel(level types.LogLevel) {
	g := *globalLogger
	g.f.currentLevel = level
	SetLogger(&g)
}

func SetLogger(newLogger *Logger) {
	mutex.Lock()
	defer mutex.Unlock()
	globalLogger = newLogger
}
