package oracled

import (
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/log"
	"github.com/kowala-tech/kcoin/p2p/discover"
	"github.com/kowala-tech/kcoin/params"
)

var (
	mapNetworkIDToGenesis = map[uint64]*core.Genesis{
		params.MainnetChainConfig.ChainID.Uint64(): core.DefaultGenesisBlock(),
		params.TestnetChainConfig.ChainID.Uint64(): core.DefaultTestnetGenesisBlock(),
	}

	mapNetworkIDToBootNodes = map[uint64][]string{
		params.MainnetChainConfig.ChainID.Uint64(): params.MainnetBootnodes,
		params.TestnetChainConfig.ChainID.Uint64(): params.TestnetBootnodes,
	}
)

func getBootstrapNodesfromNetworkID(networkID uint64) []*discover.Node {
	urls := mapNetworkIDToBootNodes[networkID]

	bootNodes := make([]*discover.Node, 0, len(urls))
	for _, url := range urls {
		node, err := discover.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "url", url, "err", err)
			continue
		}
		bootNodes = append(bootNodes, node)
	}
}
