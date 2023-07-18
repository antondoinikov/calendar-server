package logger

import "fmt"

type Logger int

const (
	NoMessage Logger = iota
	Info
	Warning
	Error
	Debug
)

func New(level string) Logger {
	switch level {
	case "Info", "info":
		return Info
	case "Warning", "warning":
		return Warning
	case "Error", "error":
		return Error
	case "Debug", "debug":
		return Debug
	}
	return NoMessage
}

func (l Logger) Info(msg string) {
	if l == Info || l == Debug {
		fmt.Println(msg)
	}
}

func (l Logger) Warning(msg string) {
	if l == Warning || l == Debug {
		fmt.Println(msg)
	}
}

func (l Logger) Error(msg string) {
	if l == Error || l == Debug {
		fmt.Println(msg)
	}
}

func (l Logger) Debug(msg string) {
	if l == Debug {
		fmt.Println(msg)
	}
}
