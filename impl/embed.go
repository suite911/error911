package impl

import (
	"github.com/suite911/error911"

	"github.com/pkg/errors"
)

type Embed struct {
	cause error
}

func (emb *Embed) Init(title string, immediateCause error, msg ...interface{}) string {
	if immediateCause == nil {
		immediateCause = errors.New(title)
	}
	text := "\uff62" + fmt.Sprintf("%q: %q", fmt.Sprint(msg...), cause) + "\uff63"
	if emb.cause == nil {
		emb.cause = errors.Wrap(immediateCause, text)
	} else {
		emb.cause = errors.Wrap(emb.cause, text)
	}
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
