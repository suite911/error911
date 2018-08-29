package myerror

import "github.com/suite911/error911"

type MyError struct {
	error911.Impl
}

func New(title string, cause error, msg ...interface{}) *MyError {
	return new(MyError).Init(title, cause, msg...)
}

func (err *MyError) Init(title string, cause error, msg ...interface{}) *MyError {
	err.Impl.Init(title, cause, msg...)
	return err
}
