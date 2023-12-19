package main

import "C"
import (
	"syscall"
	"unsafe"
	"{{.PackageName}}/pkg/{{.ExtensionName}}"
)

func sendOutput(data string, callback uintptr) {
	// Convert Go string to byte slice pointer
	outDataPtr, _ := syscall.BytePtrFromString(data)

	// Send data back - discard return/error, yolo
	syscall.SyscallN(callback, uintptr(unsafe.Pointer(outDataPtr)), uintptr(len(data)))
}

func sendError(err error, callback uintptr) {
	sendOutput(err.Error(), callback)
}

func getArgs(data uintptr, dataLen uintptr) string {
	// Get args from Pointer
	argSlice := unsafe.Slice((*byte)(unsafe.Pointer(data)), int(dataLen))
	return string(argSlice)
}

// This is the entrypoint called by the Sliver implant at runtime.
// Arguments are passed in via the `data` parameter as a byte array of size `dataLen`.
// To send data back to the implant, we recommend using the `sendOutput` helper function.
// Errors can be sent back using the `sendError` helper.

//export Run
func Run(data uintptr, dataLen uintptr, callback uintptr) uintptr {
	output, err := {{.ExtensionName}}.DoStuff(getArgs(data, dataLen))
	if err != nil {
		sendError(err, callback)
		return 1
	}
	sendOutput(output, callback)
	return 0
}

func main() {}
