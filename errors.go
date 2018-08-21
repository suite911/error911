package error911

import (
	"errors"
	"fmt"
	pkgErrors "github.com/pkg/errors"
)

// error911.Cancel

type Cancel struct {
	err error
}

func NewCancel(msg ...interface{}) error {
	return new(Cancel).Init(msg...)
}

func (err *Cancel) Init(msg ...interface{}) error {
	err.err = pkgErrors.Wrap(
		errors.New("Attempted to cancel an operation that does not support cancelation"),
		fmt.Sprint(msg...),
	)
	return err
}

func (err Cancel) Error() string {
	return err.err.Error()
}

// error911.Email

type Email struct {
	err error
}

func NewEmail(msg ...interface{}) error {
	return new(Email).Init(msg...)
}

func (err *Email) Init(msg ...interface{}) error {
	err.err = pkgErrors.Wrap(
		errors.New("An e-mail address was invalid"),
		fmt.Sprint(msg...),
	)
	return err
}

func (err Email) Error() string {
	return err.err.Error()
}

// error911.NotSupported

type NotSupported struct {
	err error
}

func NewNotSupported(msg ...interface{}) error {
	return new(NotSupported).Init(msg...)
}

func (err *NotSupported) Init(msg ...interface{}) error {
	err.err = pkgErrors.Wrap(errors.New("Not supported"), fmt.Sprint(msg...))
	return err
}

func (err NotSupported) Error() string {
	return err.err.Error()
}
