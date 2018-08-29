package impl

import (
	"github.com/suite911/error911"

	"github.com/pkg/errors"
)

type Embed struct {
	cause error
}

// Init the error stack to a non-nil error value
func (emb *Embed) Init(title string, cause error, msg ...interface{}) *Embed {
	if emb == nil {
		emb = new(Embed)
	}
	if cause == nil {
		cause = errors.New(title)
	}
	emb.cause = errors.Wrap(cause, fmt.Sprint(msg...))
	return emb
}

// Return a new error stack with the current stack pushed beneath the new error value
func (emb *Embed) New(title string, immediateCause error, msg ...interface{}) *Embed {
	if emb == nil {
		return emb.Init(title, immediateCause, msg...)
	}
	message := fmt.Sprint(msg...)
	if immediateCause != nil {
		message = fmt.Sprintf("%q: %q", message, immediateCause)
	}
	emb.cause = errors.Wrap(emb.cause, "\uff62" + message + "\uff63")
	return emb
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
