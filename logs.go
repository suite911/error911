package error911

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
)

var NeverOpenErrorsInBrowser bool

// Logs represents a log of events and a possible error stack
type Logs struct {
	Entries []*LogEntry // The entries in the log
	Error   Error       // The error stack

	title   string
}

// Create a new Logs and initialize it
func NewLog(title string) *Logs {
	return new(Logs).Init(title)
}

// Initialize the Logs
func (l *Logs) Init(title string) *Logs {
	if len(title) < 1 {
		title = "Error"
	}
	l.title = title
	return l
}

// Try to open the error in a browser window for debugging
func (l Logs) ErrorBrowser() {
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
func (l Logs) ErrorHTML() string {
	if l.Error == nil {
		return ""
	}
	_, es, est := l.Error.Stacks()
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
func (l Logs) ErrorMarkDown() string {
	if l.Error == nil {
		return ""
	}
	if l == nil {
		panic("cannot get the error from a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot get the error from an uninitialized error911.Logs")
	}
	_, es, est := l.Error.Stacks()
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
func (l Logs) ErrorText() string {
	if l.Error == nil {
		return ""
	}
	if l == nil {
		panic("cannot get the error from a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot get the error from an uninitialized error911.Logs")
	}
	_, es, est := l.Error.Stacks()
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
func (l *Logs) Log(language, subTitle string, msg ...interface{}) {
	if l == nil {
		panic("cannot log to a nil error911.Logs")
	}
	l.Entries = append(l.Entries, NewLogEntry(language, subTitle, msg...))
}
