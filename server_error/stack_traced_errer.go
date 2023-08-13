package server_error

import (
	"fmt"
	"runtime/debug"
)

func NewStackTracedError(err error) *StackTracedError {
	return &StackTracedError{
		err:   err,
		stack: debug.Stack(),
	}
}

func NewStackTracedErrorf(format string, args ...interface{}) *StackTracedError {
	return &StackTracedError{
		err:   fmt.Errorf(format, args...),
		stack: debug.Stack(),
	}
}

type StackTracedError struct {
	err   error
	stack []byte
}

func (ste *StackTracedError) Error() string {
	return ste.err.Error() + "\n" + string(ste.stack)
}
