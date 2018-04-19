package oracled

import (
	"math/big"

	"github.com/kowala-tech/kcoin/event"
)

const (
	priceChanSize = 1024
)

// Transactor represents a blockchain transactor
type Transactor interface {
	Start()
	Stop()
	SubmitPrice(*big.Int) error
}

type transactor struct {
	scraper  Scraper
	priceCh  chan *big.Int
	priceSub event.Subscription
}

func NewTransactor(scraper Scraper) {
	transactor := &transactor{
		scraper: scraper,
	}
}

func (tx *transactor) Start() {
	// submit authenticated prices
	tx.priceCh = make(chan *big.Int, priceChanSize)
	tx.priceSub = tx.scraper.SubscribeFeeds(tx.priceCh)
	go tx.priceBroadcastLoop()
}

func (tx *transactor) Stop() {
	tx.priceSub.Unsubscribe()
}

func (tx *transactor) priceBroadcastLoop() {
	for {
		select {
		case data := <-tx.priceCh:
			tx.broadcastPrice(data)
		case <-tx.priceSub.Err(): // Unsubscribe
			return
		}
	}
}

// @TODO (rgeraldes)
func (tx *transactor) broadcastPrice(data *big.Int) {}
