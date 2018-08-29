package error911

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

// E911 holds an error and a log; it is not an error itself
type E911 struct {
	E911Interface
	Log
}

func New(title string) *E911 {
	return new(E911).Init(title)
}

func (e *E911) Init(title string) *E911 {
	e.Log.Init(title)
	return e
}

// Try to open the error in a browser window for debugging
func (err *E911) ErrorBrowser() {
	s := `<!DOCTYPE html>
<html>
<body>` + err.ErrorHTML() + `
</body>
</html>`
	defer recover()
	browser.OpenReader(strings.NewReader(s))
}

// Get the full text of the error stacks as HTML
func (err *E911) ErrorHTML() string {
	_, es, est := err.Stacks()
	s := "<h2>" + err.log.Title + "</h2>\n"

	if len(err.log.Entries) > 0 {
		s += "<h3>Log</h3>\n"
		for _, e := range err.log.Entries {
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
func (err *E911) ErrorMarkDown() string {
	_, es, est := err.Stacks()
	s := "## " + err.log.Title + " ##\n"

	if len(err.log.Entries) > 0 {
		s += "### Log ###\n"
		for _, e := range err.log.Entries {
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
func (err *E911) ErrorText() string {
	_, es, est := err.Stacks()
	s := "=== " + err.log.Title + " ===\n"

	if len(err.log.Entries) > 0 {
		s += "=== Log\n"
		for _, e := range err.log.Entries {
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

	s += "... " + err.log.Title + " ...\n"
	return s
}

// Append an entry to the log
func (err *E911) Log(language, subTitle string, msg ...interface{}) {
	err.log.Append(language, subTitle, msg...)
}

// Push an error onto the error stack
func (err *E911) Push(msg ...interface{}) {
	err.Error.Push(err.log.Title, msg...)
}
