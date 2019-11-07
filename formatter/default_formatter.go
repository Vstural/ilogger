package formatter

import (
	"fmt"
	"ilogger/types"
	"time"
)

type DefaultFormatter struct{}

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

func (f *DefaultFormatter) Format(level types.LogLevel, args ...interface{}) string {
	return fmt.Sprintf("%s %s : %s", time.Now().String(), level, fmt.Sprint(args...))
}
