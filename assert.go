package assert

import (
	"fmt"
	"io"
	"runtime"
)

const defaultMsg = "assertion failed"

// Assert works like a traditional `assert` in other languages: it panics if the ok condition is false.
// Optionally it can be configured
func Assert(ok bool, ops ...option) {
	a := &assertion{}

	for _, o := range ops {
		o(a)
	}

	if !ok {
		a.fail()
	}
}

type assertion struct {
	msg    string
	writer io.Writer
}

func (a *assertion) fail() {
	if a.msg == "" {
		a.msg = defaultMsg
	}

	if a.writer != nil {
		// notice that we're using 2, so it will actually log the function where
		// the error happened, 0 = this function and 1 = Assert(). We don't want that.
		pc, filename, line, _ := runtime.Caller(2)

		out := fmt.Sprintf("panic! %s[%s:%d]: %s\n", runtime.FuncForPC(pc).Name(), filename, line, a.msg)
		a.writer.Write([]byte(out))
	}

	panic(a.msg)
}

// option to configure how an assertion should behave
type option func(a *assertion)

// WithMessage allows to provide a custom panic message.
// Default message is "assertion failed".
func WithMessage(s string) option {
	return func(a *assertion) { a.msg = s }
}

// WithWriter adds a writer to send the message and the code location where the assertion failed before panicking.
// Useful to log the failed assertion before crashing for example.
func WithWriter(w io.Writer) option {
	return func(a *assertion) { a.writer = w }
}
