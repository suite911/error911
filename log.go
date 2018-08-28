package error911

type Log struct {
	Title   string
	Entries []LogEntry
}

func (l Log) HTML() string {
	s := "<h2>" + l.Title + "</h2>\n"
	for _, e := range l.Entries {
		s += e.HTML()
		s += "\n"
	}
	return s
}

func (l Log) MarkDown() string {
	s := "## " + l.Title + " ##\n"
	for _, e := range l.Entries {
		s += e.MarkDown()
		s += "\n"
	}
	return s
}

func (l Log) Text() string {
	s := "=== \"" + l.Title + "\"\n"
	for _, e := range l.Entries {
		s += e.Text()
		s += "\n"
	}
	s += "...\n"
	return s
}
