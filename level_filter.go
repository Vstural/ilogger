package ilogger

import (
	"ilogger/types"
)

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
