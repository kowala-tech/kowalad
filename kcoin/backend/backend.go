package backend

import (
	"fmt"
	"sync"

	"github.com/kowala-tech/kowalad/kcoin/node"
	"github.com/kowala-tech/kowalad/kcoin/params"
)

type Backend interface {
	StartNode(config *params.Config) error
	StopNode() error
	Node() *node.Node
}

type backend struct {
	nodeMu sync.Mutex
	node   *node.Node
}

// New returns a new backend instance
func New() *backend {
	return &backend{
		node: node.New(),
	}
}

func (b *backend) StartNode(config *params.Config) error {
	b.nodeMu.Lock()
	defer b.nodeMu.Unlock()
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
	b.nodeMu.Lock()
	defer b.nodeMu.Unlock()
	return b.stopNode()
}

func (b *backend) stopNode() error {
	return b.node.Stop()
}

func (b *backend) Node() *node.Node { return b.node }
