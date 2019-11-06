package logger

import (
	"fmt"
)

type STDLogger struct{}

func NewSTDLogger() *STDLogger {
	return &STDLogger{}
}

func (s *STDLogger) Log(ss string) { fmt.Println(ss) }

// func (s *STDLogger) Debug(args ...interface{})    { fmt.Println(args...) }
// func (s *STDLogger) Info(args ...interface{})     { fmt.Println(args...) }
// func (s *STDLogger) Warning(args ...interface{})  { fmt.Println(args...) }
// func (s *STDLogger) Error(args ...interface{})    { fmt.Println(args...) }
// func (s *STDLogger) Critical(args ...interface{}) { fmt.Println(args...) }
// func (s *STDLogger) Fixed(args ...interface{})    { fmt.Println(args...) }
