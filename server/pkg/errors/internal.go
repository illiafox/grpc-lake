package errors

import (
	"runtime"
	"strconv"
)

const Separator = ": "

type InternalError struct {
	Err   error
	Scope string
	Line  string
}

func NewInternal(scope string, err error) error {
	_, file, line, _ := runtime.Caller(1)
	return InternalError{
		Err:   err,
		Scope: scope,
		Line:  file + Separator + strconv.Itoa(line),
	}
}

func (i InternalError) Unwrap() error {
	return i.Err
}

func (i InternalError) Cause() error {
	return i.Err
}
