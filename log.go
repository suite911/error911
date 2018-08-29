package error911

// Log references an error and contains a log
type Log struct {
	Entries []*LogEntry
	pError  *error
	Title   string
}

// Create a new Log and initialize it
func NewLog(title string, pErr *error) *Log {
	return new(Log).Init(title, pErr)
}

// Initialize the Log
func (l *Log) Init(title string, pErr *error) *Log {
	l.pError = pErr
	l.Title = title
	return l
}

// Append an entry to the log
func (l *Log) Log(language, subTitle string, msg ...interface{}) {
	l.Entries = append(l.Entries, NewLogEntry(language, subTitle, msg...))
}

// Push an error onto the error stack
func (l *Log) Push(msg ...interface{}) {
	if l == nil {
		panic("cannot push to nil error911.Log")
	}
	if l.pError == nil {
		panic("cannot push to uninitialized error911.Log")
	}
	l.pError.Push(l.Title, msg...)
}
