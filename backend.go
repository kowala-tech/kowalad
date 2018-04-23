package oracled

import (
	"fmt"

	"github.com/kowala-tech/kcoin/log"
)

// @TODO (rgeraldes) - add async operations
// @TODO (rgeraldes) - add mutex

type Backend interface {
	StartNode(config *Config) error
	StopNode()
	StartOracle() error
	StopOracle() error
}

// backend represents the oracle app backend
type backend struct {
	config        *Config
	node          Node
	enclave       Enclave
	oracleManager OracleManager
	log           log.Logger
}

func NewBackend() *backend {
	b := &backend{
		node:    NewLiteNode(),
		enclave: NewEnclave(),
		log:     log.New("package", "oracled/backend"),
	}
	b.node.
		b.oracleManager = NewOracleManager(b.node)

	return b
}

func (b *backend) StartNode(config *Config) error {
	b.config = config
	return b.startNode(config)
}

func (b *backend) startNode(config *Config) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("node crashed on start: %v", err)
		}
	}()
	return b.node.Start(config)
}

func (b *backend) StopNode() error {
	return b.node.Stop()
}

func (b *backend) StartOracle() {
	b.oracleManager.Start()
}

func (b *backend) StopOracle() {
	b.oracleManager.Stop()
}
