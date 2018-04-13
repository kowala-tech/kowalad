package main

import (
	"C"

	"github.com/kowala-tech/kowalad/kcoin/backend"
	"github.com/kowala-tech/kowalad/kcoin/params"
)

var (
	api backend.Backend
)

func init() {
	api = backend.New()
}

//export Error
type Error *C.char

//export Kowala_StartNode
func Kowala_StartNode() (C.int, Error) {
	// @TODO (rgeraldes) - export config ob
	config := &params.Config{}

	if err := api.StartNode(config); err != nil {
		return 0, C.CString(err.Error())
	}
	return 0, nil
}

//export Kowala_StopNode
func Kowala_StopNode() (C.int, Error) {
	if err := api.StopNode(); err != nil {
		return 0, C.CString(err.Error())
	}
	return 0, nil
}

// main is required by CGO
func main() {}
