package oracled

/*

#include "./kcoin/sgx/OracleEnclave.h"
#cgo LDFLAGS: -I./sgx -L. -oracle_enclave

*/

import "C"
import "math/big"

type DataFeeder interface {
	Feed() *big.Int
}

// enclave represents a SGX enclave
type enclave struct{}

func NewSGXEnclave() *enclave {
	return &enclave{}
}

// Feed provides a price
func (e *enclave) Feed() {}
