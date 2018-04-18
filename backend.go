package oracled

import (
	"fmt"
	"sync"

	"github.com/kowala-tech/kcoin/node"
	"github.com/kowala-tech/oracled/kcoin/params"
)

type Backend interface {
	Start(config *params.Config) error
	Stop() error
}

type backend struct {
	backendMu sync.Mutex
	node      *node.Node
}

// New returns a new backend instance
func New() *backend {
	return &backend{
		node: node.New(),
	}
}

func (b *backend) StartNode(config *params.Config) error {
	b.backendMu.Lock()
	defer b.backendMu.Unlock()
	return b.startNode(config)
}

func (b *backend) startNode(config *params.Config) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("node crashed on start: %v", err)
		}
	}()

	return b.node.Start(config)
}

func (b *backend) StopNode() error {
	b.backendMu.Lock()
	defer b.backendMu.Unlock()
	return b.stopNode()
}

func (b *backend) stopNode() error {
	return b.node.Stop()
}
