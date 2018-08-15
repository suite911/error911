package onfail

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

// Helper function for configurable fail behavior
func Fail (err, arg interface{}, calleeConf OnFail, callerArgs interface{}) error {
	return fail(InterfaceToError(err), arg, calleeConf, callerArgs)
}

// Helper function for configurable fail behavior -- error plus string
func FailEx (err error, msg string, arg interface{}, calleeConf OnFail, callerArgs interface{}) error {
	return fail(errors.WithMessage(err, msg), arg, calleeConf, callerArgs)
}

// An internal function; exported in case it is needed in user code
func InterfaceToError(err interface{}) error {
	switch err.(type) {
	case error:
		return err.(error)
	case string:
		return errors.New(err.(string))
	default:
		panic("`err` must be either `error` or `string`!")
	}
}

// An internal function; exported in case it is needed in user code
func LogFatalln(err error, one, two *log.Logger) {
	switch {
	case one != nil:
		one.Fatalln(err)
	case two != nil:
		two.Fatalln(err)
	default:
		log.Fatalln()
	}
}

// An internal function; exported in case it is needed in user code
func LogPrintln(err error, one, two *log.Logger) {
	switch {
	case one != nil:
		one.Println(err)
	case two != nil:
		two.Println(err)
	default:
		log.Println()
	}
}

// Called by the `Fail` helper function when the behavior is not configured at the call site
var Default OnFail = Panic

const LogDefaultFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile
const LogDefaultPrefixFatal = "-FATAL-\t"
const LogDefaultPrefixPrint = "-WARN-\t"
var LogFatal *log.Logger = log.New(os.Stderr, LogDefaultPrefixFatal, LogDefaultFlags)
var LogFatalTrace *log.Logger = nil
var LogPrint *log.Logger = log.New(os.Stderr, LogDefaultPrefixPrint, LogDefaultFlags)
var LogPrintTrace *log.Logger = nil

// Interface for types which can configure fail behavior
type OnFail interface {
	Fail(error, interface{})
}

// Function type for configure fail behavior by calling the function
type OnFailCallFunction func(error, interface{})

// To satisfy the `OnFail` interface
func (onFail OnFailCallFunction) Fail(err error, arg interface{}) {
	onFail(err, arg)
}

// Built-in fail behavior configuration to log fatally
var Fatal OnFailCallFunction = func(err error, arg interface{}) {
	LogFatalln(err, LogFatal, nil)
}

// Built-in fail behavior configuration to log fatally; with stack trace
var FatalTrace OnFailCallFunction = func(err error, arg interface{}) {
	LogFatalln(errors.WithStack(err), LogFatalTrace, LogFatal)
}

// Built-in fail behavior configuration to ignore the error
var Ignore OnFailCallFunction = func(err error, arg interface{}) {
}

// Built-in fail behavior configuration to panic
var Panic OnFailCallFunction = func(err error, arg interface{}) {
	panic(err)
}

// Built-in fail behavior configuration to panic; with stack trace
var PanicTrace OnFailCallFunction = func(err error, arg interface{}) {
	panic(errors.WithStack(err))
}

// Built-in fail behavior configuration to log and continue
var Print OnFailCallFunction = func(err error, arg interface{}) {
	LogPrintln(err, LogPrint, nil)
}

// Built-in fail behavior configuration to log and continue; with stack trace
var PrintTrace OnFailCallFunction = func(err error, arg interface{}) {
	LogPrintln(errors.WithStack(err), LogPrintTrace, LogPrint)
}

func fail(err error, arg interface{}, calleeConf OnFail, args interface{}) error {
	var callerConf OnFail
	set := func(arg interface{}) {
		if callerOnFail, ok := arg.(OnFail); ok && callerOnFail != nil {
			callerConf = callerOnFail
		}
	}
	switch args.(type) {
	case []OnFail:
		for _, arg := range args.([]OnFail) {
			set(arg)
		}
	case OnFail:
		set(args.(OnFail))
	case []interface{}:
		for _, arg := range args.([]interface{}) {
			set(arg)
		}
	default:
		panic("Bad arguments to fail")
	}
	switch {
	case callerConf != nil:
		callerConf.Fail(err, arg)
	case calleeConf != nil:
		calleeConf.Fail(err, arg)
	case Default != nil: // but you should never set Default to nil
		Default.Fail(err, arg)
	case Panic != nil: // but you should never set Panic to nil
		Panic.Fail(err, arg)
	default:
		panic(err)
	}
	return err
}
