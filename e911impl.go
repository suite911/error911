package error911

import (
	"fmt"

	"github.com/pkg/errors"
)

type E911Impl struct {
	err error
}

// Initialize the E911 implementation with a title
func (err *E911Impl) Init(title string) *E911Impl {
	err.log.Title = title
	return err
}

// Get the previous error, which is assumed to have caused tthis one
func (err *E911Impl) Cause() error {
	return err.err
}

// Get the text of the first error, which is assumed to have caused the others
func (err *E911Impl) Error() string {
	return err.First().Error()
}

// Get the first error, which is assumed to have caused the others
func (err *E911Impl) First() error {
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

// Push an error onto the error stack
func (err *E911Impl) Push(msg ...interface{}) {
	cause := err.err
	if cause == nil {
		cause = errors.New(err.log.Title)
	}
	err.err = errors.Wrap(cause, "\uff62" + fmt.Sprint(msg...) + "\uff63")
}

// Return the error stacks
func (err *E911Impl) Stacks() (first error, stack string, earliestStackTrace errors.StackTrace) {
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
