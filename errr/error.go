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

func NewError(e interface{}, a ...interface{}) *Error {
	err := new(Error)

	// err.Message
	var msg string
	switch t := e.(type) {
	case *Error:
		return t
	case error:
		msg = t.Error()
	case string:
		if len(a) > 0 {
			msg = fmt.Sprintf(t, a...)
		} else {
			msg = t
		}
	default:
		msg = fmt.Sprintf("%v", t)
	}
	err.Message = msg

	// err.Position
	_, file, line, ok := runtime.Caller(2)
	if ok {
		err.Position = filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	// err.StackTrace
	const size = 1 << 12
	buf := make([]byte, size)
	n := runtime.Stack(buf, false)
	err.StackTrace = string(buf[:n])

	return err
}

func NewTagError(tag string, e interface{}, a ...interface{}) *Error {
	err := NewError(e, a...)
	err.Tag = tag
	return err
}

func NewTagCodeError(tag string, code int, e interface{}, a ...interface{}) *Error {
	err := NewError(e, a...)
	err.Tag = tag
	err.Code = code
	return err
}
