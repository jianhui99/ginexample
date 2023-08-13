package server_error

import (
	"fmt"
)

// error that will be returned to the client for handling
type ServerError struct {
	msg  string
	code int
}

func (e ServerError) Code() int {
	return e.code
}

func (e ServerError) Msg() string {
	return e.msg
}

func (e ServerError) Error() string {
	return fmt.Sprintf("AppError{msg: %s, code: %v}", e.msg, e.code)
}

type BadRequestError struct {
	msg  string
	code int
}

func (e BadRequestError) Code() int {
	return e.code
}

func (e BadRequestError) Msg() string {
	return e.msg
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("AppError{msg: %s, code: %v}", e.msg, e.code)
}
