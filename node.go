package oracled

import (
	"errors"
)

const (
	// config defaults
	UseLightweightKDF = true
	NoUSB             = true
)

var (
	ErrActiveNode              = errors.New("node is already running")
	ErrLiteRegistrationFailure = errors.New("failed to register the lite protocol")
)

type Node interface {
	Stop() error
	Start() error
}

type node struct {
	*node.Node
}

/*
type




// Node represents the kowala's oracle blockchain gateway
type Node struct {
	*node.Node
}

func NewNode() *Node {
	return &Node{}
}

func (node *Node) isActive() bool {
	if node.Node == nil || node.Node.Server() == nil {
		return false
	}
	return true
}

func (node *Node) Start(config *Config) error {
	if node.isActive() {
		return ErrActiveNode
	}

	return node.start(config)
}

func (node *Node) start(config *Config) error {
	kcoinNode, err := MakeNode(config)
	if err != nil {
		return err
	}

	node.Node = kcoinNode

	return nil
}

func registerKowalaService(registry *node.Node, config *Config) error {
	protocolConfig := kcoin.DefaultConfig
	protocolConfig.Genesis =


	if err := registry.Register(func(ctx *node.ServiceContext) (node.Service, error) {


		return kcoin.New(ctx, cfg)
	}); err != nil {
		return fmt.Errorf("%v: %v", ErrLiteRegistrationFailure, err)
	}

	return nil
}

func MakeNode(config *Config) (*node.Node, error) {
	if config == nil {
		config = NewConfig()
	}

	if err := os.MkdirAll(filepath.Join(config.DataDir), os.ModePerm); err != nil {
		return nil, err
	}

	nodeConfig := getNodeConfig(config)
	node, err := node.New(nodeConfig)
	if err != nil {
		return nil, err
	}

	if config.light.Enabled {
		if err := registerKowalaService(node, config); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func getNodeConfig(config *Config) *node.Config {
	return &node.Config{
		DataDir: config.DataDir,
		UseLightweightKDF: UseLightweightKDF,
		NoUSB: NoUSB,
		P2P: p2p.Config{
			NAT: nat.Any(),
			MaxPeers: config.Maxpeers,
			MaxPendingPeers: config.MaxPendingPeers,
			BootstrapNodes: getBootstrapNodes(config.NetworkID),
		}
	}
}

*/
