package metadata

import (
	"fmt"
	"strings"
)

// Type - Use this interface to define and encapsulate the behavior of validation and tranformation
type Type interface {
	// validate the configure value
	validate(str string) error
	// get the real type of current value, if it is int, return int, if it is string return string etc.
	get(str string) (interface{}, error)
}

// StringType ...
type StringType struct {
}

func (t *StringType) validate(str string) error {
	return nil
}

func (t *StringType) get(str string) (interface{}, error) {
	return str, nil
}

type NoneEmptyStringValue struct {
	StringType
}

func (t *NoneEmptyStringValue) validate(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return ErrStringValueIsEmpty
	}
	return nil
}

// AuthModeType ...
type AuthModeType struct {
	StringType
}

func (t *AuthModeType) validate(str string) error {
	return fmt.Errorf("invalid ")
}
