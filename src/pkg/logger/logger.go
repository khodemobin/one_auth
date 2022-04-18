package logger

type Logger interface {
	Error(err error)
	Fatal(err error)
	Warn(msg string)
	Info(msg string)
}
