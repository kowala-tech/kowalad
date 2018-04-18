package oracled

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kowala-tech/kcoin/kcoin"
	"github.com/kowala-tech/kcoin/node"
	"github.com/kowala-tech/kcoin/p2p"
	"github.com/kowala-tech/kcoin/p2p/nat"
	"github.com/kowala-tech/kowalad/kcoin/params"
)

var (
	ErrActiveNode                 = errors.New("there's an active node")
	ErrInactiveNode               = errors.New("inactive node")
	ErrServiceRegistrationFailure = errors.New("failed to register a service")
)

var defaultKowalaNodeConfig = &node.Config{
	UseLightweightKDF: true,
	NoUSB:             true,
	P2P: p2p.Config{
		// @TODO (rgeraldes) - not sure about this one
		NoDiscovery: true,
		NAT:         nat.Any(),
	},
}

type Node struct {
	*node.Node
}

func New() *Node {
	return &Node{}
}

func (node *Node) isActive() bool {
	if node.Node == nil || node.Node.Server() == nil {
		return false
	}
	return true
}

func (node *Node) Start(config *params.Config) error {
	if node.isActive() {
		return ErrActiveNode
	}

	return node.start(config)
}

func registerKowalaService(registry *node.Node, config *params.Config) error {
	if err := registry.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		cfg := &kcoin.Config{}
		return kcoin.New(ctx, cfg)
	}); err != nil {
		return fmt.Errorf("%v: %v", ErrServiceRegistrationFailure, err)
	}

	return nil
}

func MakeNode(config *params.Config) (*node.Node, error) {
	if err := os.MkdirAll(filepath.Join(config.DataDir), os.ModePerm); err != nil {
		return nil, err
	}

	// @TODO (rgeraldes) - config

	node, err := node.New(&node.Config{})
	if err != nil {
		return nil, err
	}

	if err := registerKowalaService(node, config); err != nil {
		return nil, err
	}

	return node, nil
}

func (node *Node) start(config *params.Config) error {
	kcoinNode, err := MakeNode(config)
	if err != nil {
		return err
	}

	node.Node = kcoinNode

	return nil
}

func (node *Node) KowalaNode() *node.Node { return node.Node }
