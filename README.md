# ilogger

a go logger impl and interface define example

include functions below:
- local log cache with a selected cache size
- log level and output control
- default with std logger
- thread safe hot reload
- ~~easy custom your logger~~
- ~~a quick web ui to show latest log content~~
~~(require log cache enable)~~

## Level define

here is code about level define

you can also find define under
`types/log_levels.go`
```go
const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
	LevelFixed
)
```

## Interface Define

here is the interface define

outputer output log msg to the place you want
such as std(has impl) or es
```go
type Outputer interface {
	Log(s string)
}
``` 

formatter foramt input into target format
you can custom your own formatter according this interface
```go
type Formatter interface {
	Format(level types.LogLevel, args ...interface{}) string
}
```

maybe in some case, you want to custom more, so impl this
```go
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Critical(args ...interface{})
	Fixed(args ...interface{})
}
```

and set it

```go
ilogger.SetLogger(nil)
```

## usage

see `test/log_test.go`

### base usage

```go
func TestLog(t *testing.T) {
	ilogger.Debug("Debug")
	ilogger.Info("Info")
	ilogger.Warning("Warning")
	ilogger.Error("Error")
	ilogger.Critical("Critical")
	ilogger.Fixed("Fixed")
}
```

### use your custom formatter and logger

```go
func TestCustom(t *testing.T) {
	ilogger.SetLogger(ilogger.NewDefaultLogger(
		types.LevelDebug,
		logger.NewSTDLogger(),
		formatter.NewDefaultFormatter()))

	ilogger.Debug("Hello!")
}
```