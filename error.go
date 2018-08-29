package error911

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error *SError;

func New(title string, cause error, msg ...interface{}) Error {
	var err Error = new(SError)
	return err.Init(title, cause, msg...)
}

func (err Error) Init(title string, cause error, msg ...interface{}) Error {
	if cause == nil {
		cause = errors.New(title)
	}
	e := errors.Wrap(cause, fmt.Sprint(msg...))
	err.err = e
	return err
}

func (err *Error) New(title string, cause error, msg ...interface{}) {
	if *err == nil {
		*err = New(title, cause, msg...)
	} else {
		*err = errors.Wrap((*err).err, "\uff62" + fmt.Sprintf("%q: %q", fmt.Sprint(msg...), cause) + "\uff63")
	}
}
