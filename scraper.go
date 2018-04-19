package oracled

/*

#include "./kcoin/sgx/OracleEnclave.h"
#cgo LDFLAGS: -I./sgx -L. -oracle_enclave

*/

import (
	"C"
	"math/big"

	"github.com/kowala-tech/kcoin/event"
)
import "time"

// Scraper extracts data out of exchanges
type Scraper interface {
	Start()
	Stop()
	SubscribeFeeds(ch chan<- *big.Int) event.Subscription
}

func NewTrustedScraper() *trustedScraper {
	return &trustedScraper{}
}

// trustedScraper represents a SGX enclave that delivers authenticated price feeds
type trustedScraper struct {
	feed   event.Feed
	scope  event.SubscriptionScope
	doneCh chan struct{}
}

func (scraper *trustedScraper) Start(periodicity time.Duration) {
	scraper.doneCh = make(chan struct{})

	go scraper.update(periodicity)
}

func (scraper *trustedScraper) Stop() {
	close(scraper.doneCh)
}

// SubscribePriceFeeds registers a subscription for the price feeds
func (scraper *trustedScraper) SubscribeFeeds(ch chan<- *big.Int) event.Subscription {
	return scraper.scope.Track(scraper.feed.Subscribe(ch))
}

func (scraper *trustedScraper) update(peridiocity time.Duration) {
	ticker := time.NewTicker()

	for {
		select {
		case <-doneCh:
			return
		}
	}
}
