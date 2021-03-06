package errors

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestABCInfo(t *testing.T) {
	cases := map[string]struct {
		err      error
		debug    bool
		wantCode uint32
		wantLog  string
	}{
		"plain weave error": {
			err:      ErrNotFound,
			debug:    false,
			wantLog:  "not found",
			wantCode: ErrNotFound.code,
		},
		"wrapped weave error": {
			err:      Wrap(Wrap(ErrNotFound, "foo"), "bar"),
			debug:    false,
			wantLog:  "bar: foo: not found",
			wantCode: ErrNotFound.code,
		},
		"nil is empty message": {
			err:      nil,
			debug:    false,
			wantLog:  "",
			wantCode: 0,
		},
		"nil weave error is not an error": {
			err:      (*Error)(nil),
			debug:    false,
			wantLog:  "",
			wantCode: 0,
		},
		"stdlib is generic message": {
			err:      io.EOF,
			debug:    false,
			wantLog:  "internal error",
			wantCode: 1,
		},
		"stdlib returns error message in debug mode": {
			err:      io.EOF,
			debug:    true,
			wantLog:  "EOF",
			wantCode: 1,
		},
		"wrapped stdlib is only a generic message": {
			err:      Wrap(io.EOF, "cannot read file"),
			debug:    false,
			wantLog:  "internal error",
			wantCode: 1,
		},
		// This is hard to test because of attached stacktrace. This
		// case is tested in an another test.
		//"wrapped stdlib is a full message in debug mode": {
		//	err:      Wrap(io.EOF, "cannot read file"),
		//	debug:    true,
		//	wantLog:  "cannot read file: EOF",
		//	wantCode: 1,
		//},
		"custom error": {
			err:      customErr{},
			debug:    false,
			wantLog:  "custom",
			wantCode: 999,
		},
		"custom error in debug mode": {
			err:      customErr{},
			debug:    true,
			wantLog:  "custom",
			wantCode: 999,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			code, log := ABCIInfo(tc.err, tc.debug)
			if code != tc.wantCode {
				t.Errorf("want %d code, got %d", tc.wantCode, code)
			}
			if log != tc.wantLog {
				t.Errorf("want %q log, got %q", tc.wantLog, log)
			}
		})
	}
}

func TestABCIInfoStacktrace(t *testing.T) {
	cases := map[string]struct {
		err            error
		debug          bool
		wantStacktrace bool
		wantErrMsg     string
	}{
		"wrapped weave error in debug mode provides stracktrace": {
			err:            Wrap(ErrNotFound, "wrapped"),
			debug:          true,
			wantStacktrace: true,
			wantErrMsg:     "wrapped: not found",
		},
		"wrapped weave error in non-debug mode does not have stracktrace": {
			err:            Wrap(ErrNotFound, "wrapped"),
			debug:          false,
			wantStacktrace: false,
			wantErrMsg:     "wrapped: not found",
		},
		"wrapped stdlib error in debug mode provides stracktrace": {
			err:            Wrap(fmt.Errorf("stdlib"), "wrapped"),
			debug:          true,
			wantStacktrace: true,
			wantErrMsg:     "wrapped: stdlib",
		},
		"wrapped stdlib error in non-debug mode does not have stracktrace": {
			err:            Wrap(fmt.Errorf("stdlib"), "wrapped"),
			debug:          false,
			wantStacktrace: false,
			wantErrMsg:     "internal error",
		},
	}

	const thisTestSrc = "github.com/iov-one/weave/errors.TestABCIInfoStacktrace"

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			_, log := ABCIInfo(tc.err, tc.debug)
			if tc.wantStacktrace {
				if !strings.Contains(log, thisTestSrc) {
					t.Errorf("log does not contain this file stack trace: %s", log)
				}

				if !strings.Contains(log, tc.wantErrMsg) {
					t.Errorf("log does not contain expected error message: %s", log)
				}
			} else {
				if log != tc.wantErrMsg {
					t.Fatalf("unexpected log message: %s", log)
				}
			}
		})
	}
}

func TestABCIInfoHidesStacktrace(t *testing.T) {
	err := Wrap(ErrNotFound, "wrapped")
	_, log := ABCIInfo(err, false)

	if log != "wrapped: not found" {
		t.Fatalf("unexpected message in non debug mode: %s", log)
	}
}

func TestRedact(t *testing.T) {
	if err := Redact(ErrPanic, false); ErrPanic.Is(err) {
		t.Error("in non-debug mode, reduct must not pass through panic error")
	}
	if err := Redact(ErrPanic, true); !ErrPanic.Is(err) {
		t.Error("in debug mode, reduct should pass through panic error")
	}

	if err := Redact(ErrNotFound, true); !ErrNotFound.Is(err) {
		t.Error("in debug mode, reduct should pass through weave error")
	}
	if err := Redact(ErrNotFound, false); !ErrNotFound.Is(err) {
		t.Error("in non-debug mode, reduct should pass through weave error")
	}

	var cerr customErr
	if err := Redact(cerr, true); err != cerr {
		t.Error("in debug mode, reduct should pass through ABCI code error")
	}
	if err := Redact(cerr, false); err != cerr {
		t.Error("in non-debug mode, reduct should pass through ABCI code error")
	}

	serr := fmt.Errorf("stdlib error")
	if err := Redact(serr, false); err == serr {
		t.Error("in non-debug mode, reduct must not pass through a stdlib error")
	}
	if err := Redact(serr, true); err != serr {
		t.Error("in debug mode, reduct should pass through a stdlib error")
	}
}

// customErr is a custom implementation of an error that provides an ABCICode
// method.
type customErr struct{}

func (customErr) ABCICode() uint32 { return 999 }

func (customErr) Error() string { return "custom" }
