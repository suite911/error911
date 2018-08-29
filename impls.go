package error911

import (
	"fmt"

	"github.com/pkg/errors"
)

func ImplNew(title string, cause error, msg ...interface{}) *SError {
	return ImplInit(new(SError), title, cause, msg...)
}

func ImplInit(err *SError, title string, cause error, msg ...interface{}) *SError {
	if cause == nil {
		cause = errors.New(title)
	}
	e := errors.Wrap(cause, fmt.Sprint(msg...))
	err.err = e
	return err
}

func ImplNewMethod(err *SError, title string, cause error, msg ...interface{}) *SError {
	if err == nil {
		err = New(title, cause, msg...)
	} else {
		err = errors.Wrap(err.err, "\uff62" + fmt.Sprintf("%q: %q", fmt.Sprint(msg...), cause) + "\uff63")
	}
	return err
}

func ImplPush(err **SError, title string, cause error, msg ...interface{}) {
	*err = ImplNewMethod(*err, title, cause, msg...)
}
