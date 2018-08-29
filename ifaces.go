package error911

import "github.com/pkg/errors"

// Causer is defined as "causer" in "github.com/pkg/errors" but unexported
type Causer interface {
	Cause() error
}

/*
type IError interface {
	Cause() error
	Error() string
	First() error
	Stacks() (error, string, errors.StackTrace)
}
*/

// StackTracer is defined as "stackTracer" in "github.com/pkg/errors" but unexported
type StackTracer interface {
	StackTrace() errors.StackTrace
}
