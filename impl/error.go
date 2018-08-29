package impl

import "github.com/pkg/errors"

type Error interface {
	Init(title string, immediateCause error, msg ...interface{}) string

	// Get the previous error, which is assumed to have caused this one
	Cause() error

	// Get the text of the first error, which is assumed to have caused the others
	Error() string

	// Get the first error, which is assumed to have caused the others
	First() error

	// Return the error stacks
	Stacks() (first error, stack string, earliestStackTrace errors.StackTrace)
}
