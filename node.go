package kowalad

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kowala-tech/kcoin/kcoin"
	"github.com/kowala-tech/kcoin/kcoin/downloader"
	"github.com/kowala-tech/kcoin/kcoinclient"
	"github.com/kowala-tech/kcoin/log"
	"github.com/kowala-tech/kcoin/node"
	"github.com/kowala-tech/kcoin/p2p"
	"github.com/kowala-tech/kcoin/p2p/nat"
)

const (
	// config defaults
	UseLightweightKDF = true
	NoUSB             = true
)

var (
	// errors
	ErrActiveNode               = errors.New("node is already running")
	ErrInactiveNode             = errors.New("node is not running")
	ErrLightRegistrationFailure = errors.New("failed to register the light protocol")
)

type Node interface {
	Start(config *Config) error
	Stop() error
	Client() (*kcoinclient.Client, error)
}

// lightNode represents the kowala's blockchain gateway
type lightNode struct {
	*node.Node
	networkID uint64
	log       log.Logger
}

func NewLightNode() *lightNode {
	return &lightNode{}
}

func (ln *lightNode) Start(config *Config) error {
	if ln.isActive() {
		return ErrActiveNode
	}

	kcoinNode, err := MakeNode(config)
	if err != nil {
		return err
	}

	ln.Node = kcoinNode
	ln.networkID = config.NetworkID

	return ln.Node.Start()
}

func (ln *lightNode) Stop() error {
	if !ln.isActive() {
		return ErrInactiveNode
	}

	return ln.Node.Stop()
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

	if config.LightService.Enabled {
		if err := registerLightKowalaService(node, config); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func registerLightKowalaService(registry *node.Node, config *Config) error {
	serviceConfig := &kcoin.DefaultConfig
	serviceConfig.Genesis = mapNetworkIDToGenesis[config.NetworkID]
	// @TODO (rgeraldes) - replace with light sync in the future
	serviceConfig.SyncMode = downloader.FullSync
	serviceConfig.NetworkId = config.NetworkID

	if err := registry.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return kcoin.New(ctx, serviceConfig)
	}); err != nil {
		return fmt.Errorf("%v: %v", ErrLightRegistrationFailure, err)
	}

	return nil
}

func (ln *lightNode) isActive() bool {
	if ln.Node == nil || ln.Node.Server() == nil {
		return false
	}
	return true
}

func (ln *lightNode) Client() (*kcoinclient.Client, error) {
	rpcClient, err := ln.Node.Attach()
	if err != nil {
		return nil, err
	}

	return kcoinclient.NewClient(rpcClient), nil
}

func getNodeConfig(config *Config) *node.Config {
	return &node.Config{
		// @TODO (rgeraldes) uint64 > int
		DataDir:           filepath.Join(config.DataDir, strconv.Itoa(int(config.NetworkID))),
		UseLightweightKDF: UseLightweightKDF,
		NoUSB:             NoUSB,
		P2P: p2p.Config{
			NAT:             nat.Any(),
			MaxPeers:        config.Maxpeers,
			MaxPendingPeers: config.MaxPendingPeers,
			BootstrapNodes:  getBootstrapNodesfromNetworkID(config.NetworkID),
		},
	}
}
