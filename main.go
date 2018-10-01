// go-shutdownmonitor project main.go
package main

import (
	"fmt"
	"syscall"
)

const (
	wm_syscommand   = 0x0112
	sc_monitorpower = 0xf170
)

var (
	user32, _       = syscall.LoadLibrary("user32.dll")
	findwindowex, _ = syscall.GetProcAddress(user32, "FindWindowExA")
	sendmessage, _  = syscall.GetProcAddress(user32, "SendMessageA")
)

func abort(funcname string, err syscall.Errno) {
	panic(funcname + " failed: " + err.Error())
}

func FindWindowEx() uintptr {
	ret, _, callErr := syscall.Syscall6(findwindowex, 4, 0, 0, 0, 0, 0, 0)
	if callErr != 0 {
		abort("Call FindWindowEx", callErr)
	}

	return uintptr(ret)
}

func SendMessage(handle uintptr) (result int) {
	ret, _, callErr := syscall.Syscall6(sendmessage, 4, handle, wm_syscommand, sc_monitorpower, 2, 0, 0)
	if callErr != 0 {
		abort("Call SendMessage", callErr)
	}
	result = int(ret)
	return
}

func main() {
	defer syscall.FreeLibrary(user32)
	HWND_BROADCAST := FindWindowEx()
	SendMessage(HWND_BROADCAST)
	fmt.Println("Shutdown monitor!")
}
