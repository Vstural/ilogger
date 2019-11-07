package test

import (
	"ilogger"
	"ilogger/formatter"
	"ilogger/types"
	"testing"
)

func TestLog(t *testing.T) {
	ilogger.Debug("Debug")
	ilogger.Info("Info")
	ilogger.Warning("Warning")
	ilogger.Error("Error")
	ilogger.Critical("Critical")
	ilogger.Fixed("Fixed")
}

type levelTestLog struct {
	outputCount int
}

func (l *levelTestLog) Log(s string) {
	l.outputCount++
}

func TestLevelControl(t *testing.T) {
	outputer := &levelTestLog{}

	lg := ilogger.NewDefaultLogger(
		types.LevelInfo,
		outputer,
		formatter.NewDefaultFormatter())

	lg.Debug("")
	if outputer.outputCount != 0 {
		t.Error()
	}
	lg.Info("")
	if outputer.outputCount != 1 {
		t.Error()
	}
	lg.Fixed("")
	if outputer.outputCount != 2 {
		t.Error()
	}
}
