package oracled

import (
	"github.com/kowala-tech/kcoin/log"
	"github.com/kowala-tech/kcoin/p2p/discover"
	"github.com/kowala-tech/kcoin/params"
)

func getBootstrapNodes(networkID uint64) []*discover.Node {
	urls := params.MainnetBootnodes
	switch networkID {
	case TestnetChainConfig.ChainID.Uint64():
		urls = params.TestnetBootnodes
	}

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
