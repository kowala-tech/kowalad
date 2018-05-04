package main

import (
	"C"

	"github.com/kowala-tech/kowalad"
)

var (
	api = kowalad.NewAPI()
)

//export Error
type Error *C.char

//export StartNode
func StartNode() (C.int, Error) {
	if err := api.StartNode(nil); err != nil {
		return 0, C.CString(err.Error())
	}
	return 0, nil
}

//export StopNode
func StopNode() (C.int, Error) {
	if err := api.StopNode(); err != nil {
		return 0, C.CString(err.Error())
	}
	return 0, nil
}

// main is required by CGO
func main() {}
