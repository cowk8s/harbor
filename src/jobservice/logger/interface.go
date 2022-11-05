package logger

type Interface interface {
	Debug(v ...interface{})

	Debugf(format string, v ...interface{})

	Info(v ...interface{})

	Infof(format string, v ...interface{})
}
