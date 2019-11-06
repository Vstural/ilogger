package ilogger


type Logger interface {
	Debug(args... interface{})
	Info(args... interface{})
	Warning(args... interface{})
	Error(args... interface{})
	Critical(args... interface{})
	Fixed(args... interface{})
}



