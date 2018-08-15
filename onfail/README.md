# onfail
The `onfail` package allows specifiable behavior when an error is encountered.

## Imports
```go
import "github.com/suite911/error911/onfail"
```

## Usage Examples
```go
package main

import "github.com/suite911/error911/onfail"

func gowrong(onFail ...onfail.OnFail) {
	onfail.Fail("Oh noes, something went wrong! >~<", nil, onfail.Fatal, onFail)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			gowrong() // logs the error fatally
		}
	}
	gowrong(onfail.Ignore) // ignores the error
	gowrong(onfail.PrintTrace) // logs the error, along with a stack trace
	gowrong(onfail.Panic) // panics the error
}
```
