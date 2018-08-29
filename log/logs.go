package error911

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
)

var NeverOpenErrorsInBrowser bool

// Logs references an error and contains a log
type Logs struct {
	Entries []*LogEntry // The entries in the log

	pIError interface{}
	title   string
}

// Create a new Logs and initialize it
func NewLog(title string, pIError interface{}) *Logs {
	return new(Logs).Init(title, pIError)
}

// Initialize the Logs
func (l *Logs) Init(title string, pIError interface{}) *Logs {
	if len(title) < 1 {
		title = "Error"
	}
	if pIError == nil {
		panic("pIError is mandatory")
	}
	pInterface, ok := pIError.(*interface{})
	if !ok {
		panic(fmt.Sprintf("pIError must be pointer to type implementing " +
			"github.com/suite911/error911.IError, not %T", pIError))
	}
	if _, ok := (*pInterface).(IError); !ok {
		panic(fmt.Sprintf("pIError must be pointer to type implementing " +
			"github.com/suite911/error911.IError, not %T", pIError))
	}
	l.pIError = pIError
	l.title = title
	return l
}

// Try to open the error in a browser window for debugging
func (l *Logs) ErrorBrowser() {
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
func (l *Logs) ErrorHTML() string {
	if l == nil {
		panic("cannot get the error from a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot get the error from an uninitialized error911.Logs")
	}
	_, es, est := l.iError().Stacks()
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
func (l *Logs) ErrorMarkDown() string {
	if l == nil {
		panic("cannot get the error from a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot get the error from an uninitialized error911.Logs")
	}
	_, es, est := l.iError().Stacks()
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
func (l *Logs) ErrorText() string {
	if l == nil {
		panic("cannot get the error from a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot get the error from an uninitialized error911.Logs")
	}
	_, es, est := l.iError().Stacks()
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
func (l *Logs) Logs(language, subTitle string, msg ...interface{}) {
	if l == nil {
		panic("cannot log to a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot log to an uninitialized error911.Logs")
		// actually we can but it's better to fail fast
	}
	l.Entries = append(l.Entries, NewLogEntry(language, subTitle, msg...))
}

// Push an error onto the error stack
func (l *Logs) Push(msg ...interface{}) {
	if l == nil {
		panic("cannot push to a nil error911.Logs")
	}
	if l.pIError == nil {
		panic("cannot push to an uninitialized error911.Logs")
	}
	l.iError().Push_(l.title, msg...)
}

func (l *Logs) iError() IError {
	return (*l.pIError.(*interface{})).(IError)
}