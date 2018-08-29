package error911

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

var NeverOpenErrorsInBrowser bool

// Log references an error and contains a log
type Log struct {
	Entries []*LogEntry // The entries in the log

	pError  *error
	title   string
}

// Create a new Log and initialize it
func NewLog(title string, pErr *error) *Log {
	return new(Log).Init(title, pErr)
}

// Initialize the Log
func (l *Log) Init(title string, pErr *error) *Log {
	l.pError = pErr
	l.title = title
	return l
}

// Try to open the error in a browser window for debugging
func (l *Log) ErrorBrowser() {
	if NeverOpenErrorsInBrowser {
		return
	}
	s := `<!DOCTYPE html>
<html>
<body>` + l.ErrorHTML() + `
</body>
</html>`
	defer recover()
	browser.OpenReader(strings.NewReader(s))
}

// Get the full text of the error stacks as HTML
func (l *Log) ErrorHTML() string {
	if l == nil {
		panic("cannot get the error from a nil error911.Log")
	}
	if l.pError == nil {
		panic("cannot get the error from an uninitialized error911.Log")
	}
	_, es, est := l.pError.Stacks()
	s := "<h2>" + l.title + "</h2>\n"

	if len(l.Entries) > 0 {
		s += "<h3>Log</h3>\n"
		for _, e := range l.Entries {
			s += e.HTML()
			s += "\n"
		}
	}

	if len(es) > 0 {
		s += "<h3>Errors</h3>\n<code>"
		s += es
		s += "</code>\n"
	}

	if len(est) > 0 {
		s += "<h3>Earliest Stack Trace</h3>\n<code>"
		for _, f := range est {
			s += fmt.Sprintf("%s:%d %n\n", f, f, f)
		}
		s += "</code>\n"
	}

	return s
}

// Get the full text of the error stacks as markdown
func (l *Log) ErrorMarkDown() string {
	if l == nil {
		panic("cannot get the error from a nil error911.Log")
	}
	if l.pError == nil {
		panic("cannot get the error from an uninitialized error911.Log")
	}
	_, es, est := l.pError.Stacks()
	s := "## " + l.title + " ##\n"

	if len(l.Entries) > 0 {
		s += "### Log ###\n"
		for _, e := range l.Entries {
			s += e.MarkDown()
			s += "\n"
		}
	}

	if len(es) > 0 {
		s += "### Errors ###\n```\n"
		s += strings.Replace(es, "\n", "\n    ", -1)
		s += "\n```\n"
	}

	if len(est) > 0 {
		s += "### Earliest Stack Trace ###\n```\n"
		for _, f := range est {
			s += fmt.Sprintf("%s:%d %n\n", f, f, f)
		}
		s += "\n```\n"
	}

	return s
}

// Get the full text of the error stacks as plain text
func (l *Log) ErrorText() string {
	if l == nil {
		panic("cannot get the error from a nil error911.Log")
	}
	if l.pError == nil {
		panic("cannot get the error from an uninitialized error911.Log")
	}
	_, es, est := l.pError.Stacks()
	s := "=== " + l.title + " ===\n"

	if len(l.Entries) > 0 {
		s += "=== Log\n"
		for _, e := range l.Entries {
			s += e.Text()
			s += "\n"
		}
	}

	if len(es) > 0 {
		s += "=== Errors\n"
		s += strings.Replace(es, "\n", "\n    ", -1)
		s += "\n"
	}

	if len(est) > 0 {
		s += "=== Earliest Stack Trace\n"
		for _, f := range est {
			s += fmt.Sprintf("%s:%d %n\n", f, f, f)
		}
		s += "\n"
	}

	s += "... " + l.title + " ...\n"
	return s
}

// Append an entry to the log
func (l *Log) Log(language, subTitle string, msg ...interface{}) {
	if l == nil {
		panic("cannot log to a nil error911.Log")
	}
	if l.pError == nil {
		panic("cannot log to an uninitialized error911.Log")
		// actually we can but it's better to fail fast
	}
	l.Entries = append(l.Entries, NewLogEntry(language, subTitle, msg...))
}

// Push an error onto the error stack
func (l *Log) Push(msg ...interface{}) {
	if l == nil {
		panic("cannot push to nil error911.Log")
	}
	if l.pError == nil {
		panic("cannot push to an uninitialized error911.Log")
	}
	l.pError.Push(l.title, msg...)
}
