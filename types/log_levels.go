package types

type LogLevel int

const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
	LevelFixed
)

var levelMapping = [...]string{
	"debug",
	"info",
	"warning",
	"error",
	"critical",
	"fixed",
}

func (l LogLevel) String() string {
	return levelMapping[l]
}
