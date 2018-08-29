package error911

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	err error
}

// Get the previous error, which is assumed to have caused tthis one
func (err *Error) Cause() error {
	return err.err
}

// Get the text of the first error, which is assumed to have caused the others
func (err *Error) Error() string {
	return err.First().Error()
}

// Get the first error, which is assumed to have caused the others
func (err *Error) First() error {
	var e error = err
	if e != nil {
		for {
			if c, ok := e.(Causer); ok {
				if prev := c.Cause(); prev != nil {
					e = prev
					continue
				}
			}
			break
		}
	}
	return e
}

// [internal] Used internally to push an error onto the error stack
func (err *Error) Push_(title string, msg ...interface{}) {
	cause := err.err
	if cause == nil {
		cause = errors.New(title)
	}
	err.err = errors.Wrap(cause, "\uff62" + fmt.Sprint(msg...) + "\uff63")
}

// Return the error stacks
func (err *Error) Stacks() (first error, stack string, earliestStackTrace errors.StackTrace) {
	var est errors.StackTrace
	var es string
	var e error = err
	if e != nil {
		for {
			if len(es) > 0 {
				es += "\n"
			}
			ch := " "
			if st, ok := e.(StackTracer); ok {
				if t := st.StackTrace(); t != nil {
					est = t
					ch = "*"
				}
			}
			es += ch
			es += e.Error()
			if c, ok := e.(Causer); ok {
				if prev := c.Cause(); prev != nil {
					e = prev
					continue
				}
			}
			break
		}
	}
	return e, es, est
}
