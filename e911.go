package error911

import (
	"fmt"
	"strings"
)

type E911 interface {
	Cause() error
	Error() string
	ErrorHTML() string
	ErrorMarkDown() string
	ErrorText() string
	First() error
	Log(string, string, ...interface{})
	LogHTML() string
	LogMarkDown() string
	LogText() string
	Stacks() (error, string, errors.StackTrace)
}

type E911Impl struct {
	err error
	log Log
}

// Get the previous error, which is assumed to have caused tthis one
func (err *E911Impl) Cause() error {
	return err.err
}

// Get the text of the first error, which is assumed to have caused the others
func (err *E911Impl) Error() string {
	return err.First().Error()
}

// Get the full text of the error stacks as HTML
func (err *E911Impl) ErrorHTML() string {
	_, es, est := err.Stacks()
	s := "<h2>" + l.Title + "</h2>\n"

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
func (err *E911Impl) ErrorMarkDown() string {
	_, es, est := err.Stacks()
	s := "## " + l.Title + " ##\n"

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
func (err *E911Impl) ErrorText() string {
	_, es, est := err.Stacks()
	s := "=== " + l.Title + " ===\n"

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

	s := "... " + l.Title + " ...\n"
	return s
}

// Get the first error, which is assumed to have caused the others
func (err *E911Impl) First() error {
	e := err
	if e != nil {
		for {
			if c, ok := e.(Causer); ok {
				if prev := c.Cause(); prev != nil {
					e = prev
					continue
				}
			}
			break
		}
	}
	return e
}

// Append an entry to the log
func (err *E911Impl) Log(language, subTitle string, msg ...interface{}) {
	err.log.Append(language, subTitle, msg...)
}

// Return the log as HTML
func (err *E911Impl) LogHTML() string {
	return err.log.HTML()
}

// Return the log as markdown
func (err *E911Impl) LogMarkDown() string {
	return err.log.MarkDown()
}

// Return the log as text
func (err *E911Impl) LogText() string {
	return err.log.Text()
}

// Return the error stacks
func (err *E911Impl) Stacks() (first error, stack string, earliestStackTrace errors.StackTrace) {
	var est errors.StackTrace
	var es string
	e := err
	if e != nil {
		for {
			if len(es) > 0 {
				es += "\n"
			}
			ch := " "
			if st, ok := e.(StackTracer); ok {
				if t := st.StackTrace(); t != nil {
					est = t
					ch = "*"
				}
			}
			es += ch
			es += e.Error()
			if c, ok := e.(Causer); ok {
				if prev := c.Cause(); prev != nil {
					e = prev
					continue
				}
			}
			break
		}
	}
	return e, es, est
}
