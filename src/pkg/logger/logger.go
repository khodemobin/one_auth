package logger

import (
	"errors"
	"fmt"
)

type ErrorType interface {
	string | error
}

type Logger interface {
	Error(msg ErrorType)
	Fatal(msg ErrorType)
	Warn(msg ErrorType)
	Info(msg ErrorType)
}

func GetError(message ErrorType) error {
	switch msg := message.(type) {
	case error:
		return msg
	case string:
		return errors.New(msg)
	default:
		panic(fmt.Sprintf("message %v has unknown type %v", message, msg))
	}
}
