package serrors

import (
	"errors"
	"fmt"
)

type Error struct {
	error      error
	attributes []any
}

func New(message string, attributes ...any) *Error {
	return &Error{
		error:      errors.New(message),
		attributes: attributes,
	}
}

func Format(format string, arguments ...any) func(...any) *Error {
	return func(attributes ...any) *Error {
		return &Error{
			error:      fmt.Errorf(format, arguments...),
			attributes: attributes,
		}
	}
}

func (err *Error) Unwrap() error {
	return err.error
}

func (err *Error) Error() string {
	return err.error.Error()
}

func Attributes(err error) []any {
	var attributes []any

	for current := error(err); current != nil; current = errors.Unwrap(current) {
		if current, ok := current.(*Error); ok {
			attributes = append(attributes, current.attributes...)
		}
	}

	return attributes
}
