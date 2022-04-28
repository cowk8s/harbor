package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// NOTE: the default depth for the logger is 3 so that we can get the correct file and line when we use the logger to log message
var logger = New(os.Stdout, NewTextFormatter(), 3)

const srcSeparator = "harbor" + string(os.PathSeparator) + "src"

// Logger provides a struct with fileds that describe the details of logger.
type Logger struct {
	out       io.Writer
	fmtter    Formatter
	callDepth int
	fieldsStr string
}

func New(out io.Writer, fmtter Formatter, options ...interface{}) *Logger {
	// Default set to 3
	depth := 3
	// If passed in as option, then reset depth
	// Use index 0
	if len(options) > 0 {
		d, ok := options[0].(int)
		if ok && d > 0 {
			depth = d
		}
	}

	return &Logger{
		out:       out,
		fmtter:    fmtter,
		callDepth: depth,
	}
}

func (l *Logger) clone() *Logger {
	return &Logger{
		out:       l.out,
		fmtter:    l.fmtter,
		callDepth: l.callDepth,
		fieldsStr: l.fieldsStr,
	}
}

// WithDepth returns cloned logger with new depth
func (l *Logger) WithDepth(depth int) *Logger {
	r := l.clone()
	r.callDepth = depth

	return r
}

func (l *Logger) output(record *Record) (err error) {
	b, err := l.fmtter.Format(record)
	if err != nil {
		return
	}

	_, err = l.out.Write(b)

	return
}

// Info ...
func (l *Logger) Info(v ...interface{}) {
	record := NewRecord(time.Now(), fmt.Sprint(v...), l.getLine())
	l.output(record)
}

func (l *Logger) getLine() string {
	var str string
	str = line(l.callDepth)

	str = str + l.fieldsStr

	if str != "" {
		str = str + ":"
	}

	return str
}

func Info(v ...interface{}) {
	logger.WithDepth(4).Info(v...)
}

func line(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"

		line = 0
	}
	l := strings.SplitN(file, srcSeparator, 2)
	if len(l) > 1 {
		file = l[1]
	}
	return fmt.Sprintf("[%s:%d]", file, line)
}
