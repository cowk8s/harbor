package log

import (
	"fmt"
	"time"
)

var defaultTimeFormat = time.RFC3339

type TextFormatter struct {
	timeFormat string
}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{
		timeFormat: defaultTimeFormat,
	}
}

func (t *TextFormatter) Format(r *Record) (b []byte, err error) {
	s := fmt.Sprintf("%s", r.Time.Format(t.timeFormat))

	if len(r.Line) != 0 {
		s = s + r.Line + " "
	}

	if len(r.Msg) != 0 {
		s = s + r.Msg
	}

	b = []byte(s)

	if len(b) == 0 || b[len(b)-1] != '\n' {
		b = append(b, '\n')
	}

	return
}

// SetTimeFormat sets time format of TextFormatter if the parameter fmt is not null
func (t *TextFormatter) SetTimeFormat(fmt string) {
	if len(fmt) != 0 {
		t.timeFormat = fmt
	}
}
