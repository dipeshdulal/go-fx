package main

import (
	"fmt"
	"log"

	"go.uber.org/fx"
)

// CustomString is custom string type
type CustomString string

// MyAwesomeLogger is custom logger for go fx
type MyAwesomeLogger struct {
	Prefix string
}

// LogMe logs given string
func (lg *MyAwesomeLogger) LogMe(data string) string {
	return fmt.Sprintf("%v: %v\n", lg.Prefix, data)
}

// LoggerInterface interfaces logging
type LoggerInterface interface {
	LogMe(str string) string
}

// ProvideString provides a string
func ProvideString() CustomString {
	var str CustomString = "hello world"
	return str
}

// ProvideLogger provides a logger
func ProvideLogger() LoggerInterface {
	return &MyAwesomeLogger{Prefix: "Konichiwa"}
}

// RunMe should be run
func RunMe(logger LoggerInterface, str CustomString) {
	logData := logger.LogMe("Hello Dependency Injection Works")
	fmt.Println(logData)
	fmt.Println("Received string from DI", str)
}

func main() {
	log.Println("Starting Application")
	fx.New(
		fx.Provide(ProvideString),
		fx.Provide(ProvideLogger),
		fx.Invoke(RunMe),
	).Run()
}
