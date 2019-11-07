package test

import (
	"ilogger"
	"ilogger/formatter"
	"ilogger/logger"
	"ilogger/types"
	"testing"
)

func TestCustom(t *testing.T) {
	ilogger.SetLogger(ilogger.NewDefaultLogger(
		types.LevelDebug,
		logger.NewSTDLogger(),
		formatter.NewDefaultFormatter()))

	ilogger.Debug("Hello!")
}
