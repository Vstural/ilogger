package formatter

import (
	"fmt"
	"ilogger/types"
	"time"
)

type Formatter interface {
	Format(level types.LogLevel, args ...interface{}) string
}

type DefaultFormatter struct{}

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

func (f *DefaultFormatter) Format(level types.LogLevel, args ...interface{}) string {
	return fmt.Sprintf("%s %s : %s", time.Now().String(), level, args)
}
