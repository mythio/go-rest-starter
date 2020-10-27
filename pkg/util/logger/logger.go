package logger

type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Panic(string)
	Fatal(string)
}
