package error911

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	err error
}

func New(err error, msg ...interface{}) *Error {
	return new(Error).Init(err, msg...)
}

func (err *Error) Init(err error, msg ...interface{}) *Error {
	err.err = errors.Wrap(err, fmt.Sprint(msg...))
	return err
}

func (ppError **Error) Push(title string, msg ...interface{}) {
	if *ppError == nil {
		*ppError = New(errors.New(title), msg...)
	} else {
		*ppError = errors.Wrap((*ppError).err, "\uff62" + fmt.Sprint(msg...) + "\uff63")
	}
}

// Get the previous error, which is assumed to have caused this one
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
