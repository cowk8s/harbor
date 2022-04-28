package log

import "time"

// Record holds information about log
type Record struct {
	Time time.Time // time when the log produced
	Msg  string    // content of the log
	Line string    // in which	file and line that line the log produced
}

func NewRecord(time time.Time, msg, line string) *Record {
	return &Record{
		Time: time,
		Msg:  msg,
		Line: line,
	}
}
