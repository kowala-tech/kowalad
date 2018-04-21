package main

import (
	"C"

	"github.com/kowala-tech/oracled/"
)

var (
	api = oracled.NewBackend()
)

//export Error
type Error *C.char

//export StartNode
func Oracle_StartNode() (C.int, Error) {
	// @TODO (rgeraldes) - export config ob
	config := &Config{}

	if err := api.StartNode(config); err != nil {
		return 0, C.CString(err.Error())
	}
	return 0, nil
}

//export Oracle_StopNode
func Oracle_StopNode() (C.int, Error) {
	if err := api.StopNode(); err != nil {
		return 0, C.CString(err.Error())
	}
	return 0, nil
}

//export Oracle_Start
func Oracle_Start() (C.int, Error) {
	return 0, nil
}

//export Oracle_Stop
func Oracle_Stop() (C.int, Error) {
	return 0, nil
}

// main is required by CGO
func main() {}
