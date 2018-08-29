package myerror

import "github.com/suite911/error911/impl"

type MyError struct {
	impl.Embed
}

func New(title string, cause error, msg ...interface{}) *MyError {
	err := new(MyError)
	err.Init(title, cause, msg...)
	return err
}

func (err *MyError) Init(title string, cause error, msg ...interface{}) *MyError {
	if err == nil {
		return New(title, cause, msg...)
	}
	err.Embed.Init(title, cause, msg...)
	return err
}

func (err *MyError) Push(title string, immediateCause error, msg ...interface{}) *MyError {
	if err == nil {
		return New(title, immediateCause, msg...)
	}
	err.Embed.Push(title, immediateCause, msg...)
	return err
}
