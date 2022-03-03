package retry

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ErrRetryTimeout = errors.New("retry timeout")
)

type abort struct {
	cause error
}

func (a *abort) Error() string {
	if a.cause != nil {
		return fmt.Sprintf("retry abort, error: %v", a.cause)
	}

	return "retry abort"
}

func Abort(err error) error {
	return &abort{cause: err}
}

type Options struct {
}
