package errors

import (
	"fmt"
	"log"
	"runtime"

	"github.com/pkg/errors"
)

type Label string

func Labelf(format string, args ...interface{}) Label {
	return Label(fmt.Sprintf(format, args...))
}

func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}
	err := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case uint32: // That will be a ABCIResponseCode type instead of uint32
			err.Code = arg
		case Label:
			err.Labels = append(err.Labels, string(arg))
		case string:
			err.Err = S(arg)
		case error:
			err.Err = arg
		case *Error:
			copy := *arg
			err.Err = &copy
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.E: bad call from %s:%d: %v", file, line, args)
			return S("unknown type %T, value %v in error call", arg, arg)
		}
	}

	// If previous error is the same, merge.
	if prev, ok := err.Err.(*Error); ok && prev.Code == err.Code {
		prev.Labels = append(prev.Labels, err.Labels...)
		err = prev
	}

	return err
}

type Error struct {
	// Code represents the kind of operation that failed.
	Code uint32

	// Labels carry abitrary string data that might be helpful in
	// debugging.
	Labels []string

	// The underlying error that triggered this one, if any.
	Err error
}

var _ TMError = (*Error)(nil)

func (e *Error) Error() string {
	// TODO: format nicely
	return fmt.Sprintf("%d: %v: %v", e.Code, e.Labels, e.Err)
}

func (e *Error) ABCICode() uint32 {
	return e.Code
}

func (e *Error) ABCILog() string {
	return e.Err.Error()
}

func (e *Error) Cause() error {
	panic("no one is using this")
}

func (e *Error) StackTrace() errors.StackTrace {
	panic("no one is using this")
}

// S returns an error that formats as the given text. It is intended to be used
// as the error-typed argument to the E function. Is equivalent to
// fmt.Errorf, but allows clients to import only this package for all error
// handling.
// TODO: rename this to New once the old implementation is removed
func S(msg string, args ...interface{}) error {
	return &errorString{msg: fmt.Sprintf(msg, args...)}
}

type errorString struct {
	msg string
}

func (e *errorString) Error() string {
	return e.msg
}
