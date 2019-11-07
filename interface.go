package ilogger

import (
	"ilogger/types"
)

type Outputer interface {
	Log(s string)
}

type Formatter interface {
	Format(level types.LogLevel, args ...interface{}) string
}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Critical(args ...interface{})
	Fixed(args ...interface{})
}
