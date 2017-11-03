package errr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
)

type Error struct {
	Code       int
	Tag        string
	Message    string
	Position   string
	StackTrace string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Detail() string {
	return fmt.Sprintf("(%s) %s", e.Position, e.Message)
}

func (e *Error) Stack() string {
	return fmt.Sprintf("(%s) %s\tStack: %s", e.Position, e.Message, e.StackTrace)
}

func NewError(err interface{}, code ...int) *Error {
	return newError(err, code...)
}

func NewErrorWithTag(tag string, err interface{}, code ...int) *Error {
	e := newError(err, code...)
	if len(tag) > 0 {
		e.Tag = tag
	}
	return e
}

func newError(err interface{}, code ...int) *Error {
	e := new(Error)
	// e.Code
	if len(code) > 0 {
		e.Code = code[0]
	}

	// e.Message
	var message string
	switch e := err.(type) {
	case *Error:
		return e
	case string:
		message = e
	case error:
		message = e.Error()
	default:
		message = fmt.Sprintf("%v", e)
	}
	e.Message = message

	// e.Position
	_, file, line, ok := runtime.Caller(2)
	if ok {
		e.Position = filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	// e.StackTrace
	const size = 1 << 12
	buf := make([]byte, size)
	n := runtime.Stack(buf, false)
	e.StackTrace = string(buf[:n])

	return e
}
