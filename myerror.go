package error911

type MyError struct {
	*SError
}

func New(title string, cause error, msg ...interface{}) MyError {
	return ImplNew(title, cause, msg...)
}

func (err MyError) Init(title string, cause error, msg ...interface{}) MyError {
	return ImplInit(err, title, cause, msg...)
}

func (err MyError) New(title string, cause error, msg ...interface{}) MyError {
	return ImplNewMethod(err, title, cause, msg...)
}
func (err *MyError) Push(title string, cause error, msg ...interface{}) {
	ImplPush(err, title, cause, msg...)
}
