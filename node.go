package oracled

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kowala-tech/kcoin/kcoin"
	"github.com/kowala-tech/kcoin/kcoin/downloader"
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
	ErrActiveNode              = errors.New("node is already running")
	ErrInactiveNode            = errors.New("node is not running")
	ErrLiteRegistrationFailure = errors.New("failed to register the lite protocol")
)

type Node interface {
	Start(config *Config) error
	Stop() error
	LiteKowalaService() (*kcoin.Kowala, error)
}

// liteNode represents the kowala's blockchain gateway
type liteNode struct {
	*node.Node
	networkID uint64
}

func NewLiteNode() *liteNode {
	return &liteNode{}
}

func (lite *liteNode) Start(config *Config) error {
	if lite.isActive() {
		return ErrActiveNode
	}

	kcoinNode, err := MakeNode(config)
	if err != nil {
		return err
	}

	lite.Node = kcoinNode
	lite.networkID = config.NetworkID

	return lite.Node.Start()
}

func (lite *liteNode) Stop() error {
	if !lite.isActive() {
		return ErrInactiveNode
	}

	return lite.Node.Stop()
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

	if config.Lite.Enabled {
		if err := registerLiteKowalaService(node, config); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func registerLiteKowalaService(registry *node.Node, config *Config) error {
	serviceConfig := &kcoin.DefaultConfig
	serviceConfig.Genesis = mapNetworkIDToGenesis[config.NetworkID]
	// @TODO (rgeraldes) - replace with light sync in the future
	serviceConfig.SyncMode = downloader.FullSync
	serviceConfig.NetworkId = config.NetworkID

	if err := registry.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return kcoin.New(ctx, serviceConfig)
	}); err != nil {
		return fmt.Errorf("%v: %v", ErrLiteRegistrationFailure, err)
	}

	return nil
}

func (lite *liteNode) isActive() bool {
	if lite.Node == nil || lite.Node.Server() == nil {
		return false
	}
	return true
}

// @TODO (rgeraldes) - *kcoin.Kowala to be replaced by liteKowalaService
func (lite *liteNode) LiteKowalaService() (l *kcoin.Kowala, err error) {
	if err := lite.Node.Service(&l); err != nil {
		return nil, fmt.Errorf("service unavailable: %v", err)
	}
	return
}

func (lite *liteNode) NetworkID() uint64 {
	return lite.networkID
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
