package error911

type Error *SError;

func New(title string, cause error, msg ...interface{}) Error {
	return ImplNew(title, cause, msg...)
}

func (err Error) Init(title string, cause error, msg ...interface{}) Error {
	return ImplInit(err, title, cause, msg...)
}

func (err Error) New(title string, cause error, msg ...interface{}) Error {
	return ImplNewMethod(err, title, cause, msg...)
}
func (err *Error) Push(title string, cause error, msg ...interface{}) {
	ImplPush(err, title, cause, msg...)
}
