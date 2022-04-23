package lib

import "io"

// nopCloser is just like ioutil's, but here to let us re-read the same
// buffer inside by moving position to the start every time we done with reading
type nopCloser struct {
	io.ReadSeeker
}
