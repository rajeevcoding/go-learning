package logging

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
}
