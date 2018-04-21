package oracled

/*
#cgo LDFLAGS: -L. -lenclave
#include "trusted/Library/library-bridge.h"
*/

import (
	"math/big"
	"C"
	"unsafe"
)

type Enclave interface {
	Free()
	Price() *big.Int
}

type enclave struct {
	ptr unsafe.Pointer
	log log.Logger
}

func NewEnclave() *enclave {
	return enclave{ptr : C.ENCLAVE_NewEnclave()}
}

func (enc *enclave) Free() {
	C.ENCLAVE_DestroyEnclave(enc.ptr)
}

func (enc *enclave) Price() *big.Int {
	C.ENCLAVE_Request()
}
