package impl

import (
	"fmt"

	"github.com/suite911/error911"

	"github.com/pkg/errors"
)

type Embed struct {
	cause error
}

// Init the error stack to a non-nil error value
func (emb *Embed) Init(title string, cause error, msg ...interface{}) {
	if emb == nil {
		panic("Please handle the case of nil Embed in your wrapper; see the MyError example code for details.")
	}
	if cause == nil {
		cause = errors.New(title)
	}
	emb.cause = errors.Wrap(cause, fmt.Sprint(msg...))
}

// Get the previous error, which is assumed to have caused this one
func (emb Embed) Cause() error {
	return emb.cause
}

// Get the text of the first error, which is assumed to have caused the others
func (emb Embed) Error() string {
	return emb.First().Error()
}

// Get the first error, which is assumed to have caused the others
func (emb Embed) First() error {
	var e error = emb
	if e != nil {
		for {
			if c, ok := e.(error911.Causer); ok {
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

// Push a new error onto the error stack
func (emb *Embed) Push(title string, immediateCause error, msg ...interface{}) {
	if emb == nil {
		panic("Please handle the case of nil Embed in your wrapper; see the MyError example code for details.")
	}
	message := fmt.Sprint(msg...)
	if immediateCause != nil {
		message = fmt.Sprintf("%q: %q", message, immediateCause)
	}
	emb.cause = errors.Wrap(emb.cause, "\uff62" + message + "\uff63")
}

// Return the error stacks
func (emb Embed) Stacks() (first error, stack string, earliestStackTrace errors.StackTrace) {
	var est errors.StackTrace
	var es string
	var e error = emb
	if e != nil {
		for {
			if len(es) > 0 {
				es += "\n"
			}
			ch := " "
			if st, ok := e.(error911.StackTracer); ok {
				if t := st.StackTrace(); t != nil {
					est = t
					ch = "*"
				}
			}
			es += ch
			es += e.Error()
			if c, ok := e.(error911.Causer); ok {
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
