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

// code in here can't be tested because it relies on cgo. :(

import (
	"os"
	"unsafe"
	"os/exec"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>

char *string_from_argv(int, char**);
char *get_user(pam_handle_t *pamh);
int get_uid(char *user);
*/
import "C"

func init() {
	if !disablePtrace() {
		pamLog("unable to disable ptrace")
	}
}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	cUsername := C.get_user(pamh)
	pamLog("Init pam custom")
	if cUsername == nil {
		pamLog("No username provided")
		return C.PAM_USER_UNKNOWN
	}
	defer C.free(unsafe.Pointer(cUsername))

	uid := int(C.get_uid(cUsername))
	if uid < 0 {
		pamLog("User '%s' not exists in /etc/passwd. uid: %d \n", C.GoString(cUsername), uid)
		uid = 999999
		username := C.GoString(C.get_user(pamh))
		cmd := exec.Command("/sbin/useradd", "-m", username)
		err := cmd.Run()
		if err != nil {
			pamLog("Error adding user: '%s'", err)
		}
		pamLog("User added\n")
		// return C.PAM_USER_UNKNOWN
	}
	pamLog("User '%s' uid: %d \n", C.GoString(cUsername), uid)

	r := pamAuthenticate(os.Stderr, uid, C.GoString(cUsername))
	if r == AuthError {
		return C.PAM_AUTH_ERR
	}
	return C.PAM_SUCCESS
}

//export pam_sm_acct_mgmt
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("Calling pam_sm_acct_mgmt\n")
	return C.PAM_SUCCESS
	// return C.PAM_AUTHTOK_EXPIRED
}

//export pam_sm_chauthtok
func pam_sm_chauthtok(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("Calling pam_sm_chauthtok\n")
	return C.PAM_SUCCESS
}

//export pam_sm_open_session
func pam_sm_open_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("Calling pam_sm_open_session\n")
	return C.PAM_SUCCESS
}

//export pam_sm_close_session
func pam_sm_close_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("Calling pam_sm_close_session\n")
	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("Calling pam_sm_setcred\n")
	return C.PAM_SUCCESS
}
