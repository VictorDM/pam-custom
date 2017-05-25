// +build darwin linux

/*
Copyright (c) 2017 Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"io"
	"log/syslog"
	"runtime"
)

// AuthResult is the result of the authentcate function.
type AuthResult int

const (
	// AuthError is a failure.
	AuthError AuthResult = iota
	// AuthSuccess is a success.
	AuthSuccess
)

func pamLog(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-custom")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

func pamAuthenticate(w io.Writer, uid int, username string) AuthResult {
	runtime.GOMAXPROCS(1)

	if username == "victor" || username == "pepe" {
		pamLog("Authentication %s: %s\n", "OK", username)
		return AuthSuccess
	} else {
		pamLog("Authentication %s: %s\n", "KO", username)
		return AuthError
	}
}

func main() {}
