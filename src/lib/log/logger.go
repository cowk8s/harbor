package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

var logger = New(os.Stdout, NewTextFormatter(), WarningLevel, 3)

type Logger struct {
	out    io.Writer
	fmtter Formatter
	lvl    Level
}

func New(out io.Writer, fmtter Formatter, lvl Level, options ...interface{}) *Logger {

	return &Logger{
		out:    out,
		fmtter: fmtter,
	}
}

func (l *Logger) output(record *Record) (err error) {
	return
}

// Info ...
func (l *Logger) Info(v ...interface{}) {
	if l.lvl <= InfoLevel {
		record := NewRecord(time.Now(), fmt.Sprint(v...), l.getLine(), InfoLevel)
		l.output(record)
	}
}

func (l *Logger) getLine() string {
	var str string

	return str
}
