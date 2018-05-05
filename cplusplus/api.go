package main

import (
	"C"
	"unsafe"

	"github.com/kowala-tech/kowalad"
)

var (
	api = kowalad.NewAPI()
)

//export StartNode
func StartNode() C.int {
	if err := api.StartNode(nil); err != nil {
		return -1
	}
	return 0
}

//export StopNode
func StopNode() C.int {
	if err := api.StopNode(); err != nil {
		return -1
	}
	return 0
}

//export SendRawTransaction
func SendRawTransaction(data *C.uchar, len C.int) C.int {
	godata := C.GoBytes(unsafe.Pointer(data), len)
	if err := api.SendRawTransaction(godata); err != nil {
		return -1
	}
	return 0
}

// main is required by CGO
func main() {}
