package runtime

import "time"

const (
	dialConnectionTimeout = 30 * time.Second
	dialReadTimeout = 10 * time.Second
	dialWriteTimeout = 10 * time.Second
)

// JobService ...
var JobService 