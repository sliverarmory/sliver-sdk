package parser

import (
	"encoding/binary"
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/text/encoding/unicode"
)

type DataParser struct {
	original []byte
	n        int
}

// NewParser takes a pointer to the data and the length of the data
// and returns a DataParser object
func NewParser(data, dataLen uintptr) (*DataParser, error) {
	if data == 0 || dataLen == 0 {
		return nil, fmt.Errorf("no data to parse")
	}
	//turn uintptrs into slices
	dp := DataParser{
		original: unsafe.Slice((*byte)(unsafe.Pointer(data)), dataLen),
		n:        4,
	}
	return &dp, nil
}

func (dp *DataParser) GetString() (string, error) {
	outStr, err := dp.GetData()
	if err != nil {
		return "", err
	}
	// strings are NULL terminated, so we remove the NULL byte here
	return string(outStr[:len(outStr)-1]), nil
}

func (dp *DataParser) GetWString() (string, error) {
	outStr, err := dp.GetData()
	if err != nil {
		return "", err
	}
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	return decoder.Bytes(outStr[:len(outStr)-2]) //strip trailing nulls before decoding
}

// GetData returns a slice of bytes. The underlying data could be a string, a file, etc.
func (dp *DataParser) GetData() ([]byte, error) {
	if dp.GetDataLength() < 4 {
		return nil, fmt.Errorf("no more data to return")
	}
	//extract the length
	l := binary.LittleEndian.Uint32(dp.original[dp.n:])
	//increment n
	dp.n += 4
	//copy to a new buffer to avoid mutating the underlying state
	if dp.GetDataLength() < int(l) {
		return nil, fmt.Errorf("no more data to return")
	}
	rb := make([]byte, l)
	c := copy(rb, dp.original[dp.n:])
	if c != int(l) {
		return nil, fmt.Errorf("expected to copy %x but copied %x", l, c)
	}
	//increment n
	dp.n += c

	return rb, nil
}

// GetInt returns a uint32 from the DataParser
func (dp *DataParser) GetInt() (uint32, error) {
	if dp.GetDataLength() < 4 {
		return 0, fmt.Errorf("no more data to return")
	}
	r := binary.LittleEndian.Uint32(dp.original[dp.n:])
	dp.n += 4
	return r, nil
}

// GetShort returns a uint16 from the DataParser
func (dp *DataParser) GetShort() (uint16, error) {
	if dp.GetDataLength() < 2 {
		return 0, fmt.Errorf("no more data to return")
	}
	r := binary.LittleEndian.Uint16(dp.original[dp.n:])
	dp.n += 2
	return r, nil
}

// GetDataLength returns the remaining length of unparsed data
func (dp *DataParser) GetDataLength() int {
	return len(dp.original) - dp.n
}

type OutputBuffer struct {
	b        strings.Builder
	done     bool
	callback uintptr
}

func NewOutBuffer(callback uintptr) *OutputBuffer {
	return &OutputBuffer{b: strings.Builder{}, callback: callback}
}

func (o *OutputBuffer) SendOutput(data string) {
	//errors not captured here, where would we send them?
	o.b.WriteString(data)
	o.b.WriteString("\n") //newline
}

func (o *OutputBuffer) SendError(err error) {
	o.SendOutput(fmt.Sprintf("error: %s", err.Error()))
}

func (o *OutputBuffer) Flush() {
	//write the buffer - checks if it's already been flushed to avoid crashing the process
	if o.done {
		//figure out a way to return a message to implant that something went wrong with flush
		return
	}
	_sendOutput(o.b.String(), o.callback)
	o.done = true
}

// data should only be sent once per call from implant (for now)
func _sendOutput(data string, callback uintptr) {
	outDataPtr, err := syscall.BytePtrFromString(data)
	if err != nil {
		return
	}
	// Send data back
	syscall.SyscallN(callback, uintptr(unsafe.Pointer(outDataPtr)), uintptr(len(data)))
}
