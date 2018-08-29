package error911

type Log struct {
	Title   string
	Entries []*LogEntry
}

func NewLog(title string) *Log {
	return new(Log).Init(title)
}

func (l *Log) Init(title string) *Log {
	l.Title = title
	return l
}

func (l *Log) Log(language, subTitle string, msg ...interface{}) {
	l.Entries = append(l.Entries, NewLogEntry(language, subTitle, msg...))
}
