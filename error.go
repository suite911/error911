package error911

import "github.com/pkg/errors"

type Error interface {
	// Get the previous error, which is assumed to have caused this one
	Cause() error

	// Get the text of the first error, which is assumed to have caused the others
	Error() string

	// Get the first error, which is assumed to have caused the others
	First() error

	// Initialize the Error with an error
	Init(title string, cause error, msg ...interface{}) *MyError

	// Push an error onto the Error's stack
	Push(title string, immediateCause error, msg ...interface{}) *MyError

	// Return the error stacks
	Stacks() (first error, stack string, earliestStackTrace errors.StackTrace)
}
