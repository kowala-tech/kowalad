package oracled

import (
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/event"
)

// Scraper extracts data out of exchanges
type Scraper interface {
	Start(periodicity time.Duration)
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
	// Unsubscribe all subscriptions
	scraper.scope.Close()
	close(scraper.doneCh)
}

// SubscribePriceFeeds registers a subscription for the price feeds
func (scraper *trustedScraper) SubscribeFeeds(ch chan<- *big.Int) event.Subscription {
	return scraper.scope.Track(scraper.feed.Subscribe(ch))
}

func (scraper *trustedScraper) update(peridiocity time.Duration) {
	ticker := time.NewTicker(peridiocity)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// @TODO (rgeraldes)
		case <-scraper.doneCh:
			return
		}
	}
}
