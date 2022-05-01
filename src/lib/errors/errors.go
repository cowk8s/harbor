package errors

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// As alias function of `errors.As`
	As = errors.As
	// Is alias function of `errors.Is`
	Is = errors.Is
)

// Error ...
type Error struct {
	Cause   error  `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error returns a human readable error, error.Error() will not contains the track information. Needs it? just call error.StackTrace()
// Code will not be in the error output.
func (e *Error) Error() string {
	out := e.Message
	if e.Cause != nil {
		out = out + ": " + e.Cause.Error()
	}
	return out
}

// StackTrace ...
func (e *Error) StackTrace() string {
	return ""
}

// WithMessage ...
func (e *Error) WithMessage(format string, v ...interface{}) *Error {
	e.Message = fmt.Sprintf(format, v...)
	return e
}

// WithCode ...
func (e *Error) WithCode(code string) *Error {
	e.Code = code
	return e
}

// WithCause ...
func (e *Error) WithCause(err error) *Error {
	e.Cause = err
	return e
}

// Errors ...
type Errors []error

var _ error = Errors{}

func (errs Errors) Error() string {
	var tmpErrs struct {
		Errors []Error `json:"errors,omitempty"`
	}

	for _, e := range errs {
		err, ok := e.(*Error)
		if !ok {
			err = UnknownError(e)
		}
		if err.Code == "" {
			err.Code = GeneralCode
		}

		tmpErrs.Errors = append(tmpErrs.Errors, *err)
	}

	msg, err := json.Marshal(tmpErrs)
	if err != nil {
		return "{}"
	}
	return string(msg)
}

// NewErrs ...
func NewErrs(err error) Errors {
	return Errors{err}
}

// New ...
func New(in interface{}) *Error {
	var err error
	switch in := in.(type) {
	case error:
		err = in
	case *Error:
		err = in.Cause
	default:
		err = fmt.Errorf("%v", in)
	}

	return &Error{
		Message: err.Error(),
	}
}

// Cause gets the root error
func Cause(err error) error {
	for err != nil {
		cause, ok := err.(*Error)
		if !ok {
			break
		}
		if cause.Cause == nil {
			break
		}
		err = cause.Cause
	}
	return err
}

// IsErr checks whether the err chain contains error matches the code
func IsErr(err error, code string) bool {
	var e *Error
	if As(err, &e) {
		return e.Code == code
	}
	return false
}

// ErrCode returns code of err
func ErrCode(err error) string {
	if err == nil {
		return ""
	}

	var e *Error
	if ok := As(err, &e); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Cause != nil {
		return ErrCode(e.Cause)
	}

	return GeneralCode
}
