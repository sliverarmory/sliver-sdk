package main

import "C"
import (
	"{{.PackageName}}/pkg/parser"


	"{{.PackageName}}/pkg/{{.ExtensionName}}"
)

const (
	Success = 0
	Error = 1
)

// This is the entrypoint called by the Sliver implant at runtime.
// Arguments are passed in via the `data` parameter as a byte array of size `dataLen`.
// Use the OutputBuffer.SendOutput() and OutputBuffer.SendError() methods to
// prepare the data to be sent back to the implant.
// Data is sent with a call to OutputBuffer.Flush().

//export Run
func Run(data uintptr, dataLen uintptr, callback uintptr) uintptr {
	// Prepare the output buffer used to send data back to the implant
	outBuff := parser.NewOutBuffer(callback)
	// Create a new argument parser
	dataParser, err := parser.NewParser(data, dataLen)
	if err != nil {
		outBuff.SendError(err)
		outBuff.Flush()
	}

	// Parse arguments 
	
	// Get a string argument
	stringArg, err := dataParser.GetString()
	if err != nil {
		outBuff.SendError(err)
		outBuff.Flush()
	}

	// Get an int argument
	intArg, err := dataParser.GetInt()
	if err != nil {
		outBuff.SendError(err)
		outBuff.Flush()
	}

	if intArg > 0 {
		// do stuff
	}

	output, err := {{.ExtensionName}}.DoStuff(stringArg)
	if err != nil {
		parser.SendError(err, callback)
		return Error
	}
	parser.SendOutput(output, callback)
	return Success
}

func main() {}
