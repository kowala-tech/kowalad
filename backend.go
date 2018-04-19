package oracled

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrActiveBackend = errors.New("backend is already running")
)

// @TODO (rgeraldes) - add async operations
// @TODO (rgeraldes) - add mutex

type Backend interface {
	StartNode(config *Config) error // node starts upon application startup to sync
	StopNode()                      // stop node is called as soon as the user exits the application
	StartOracle() error
	StopOracle() error
}

type backend struct {
	node    Node
	scraper Scraper
	sender  Transactor
}

// New returns a new backend instance
func New() *backend {
	return &backend{
		started: false,
		node:    NewNode(),
		scraper: NewScraper(),
	}
}

func (b *backend) StartNode() error {
	return b.startNode(config)
}

func (b *backend) startNode(config *Config) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("node crashed on start: %v", err)
		}
	}()
	return b.node.Start(config)
}

func (b *backend) StopNode() error {
	// @TODO (rgeraldes) - ErrInactiveBackend
	return b.node.Stop()
}

func (b *backend) StartOracle() {

	// @TODO (rgeraldes) - contract contains sync period
	period := 1 * time.Hour

	/* Initialize Scraper */
	b.scraper.Start(period)

}

/*
func (b *backend) Start(config *Config) error {
	b.backendMu.Lock()
	defer b.backendMu.Unlock()

	if b.started {
		return ErrActiveBackend
	}
	b.started = true

	if err := transactor

	// @TODO (rgeraldes) - start cron job based on the latest
	// transaction posted
	return nil
}





func (b *backend) Stop() error {
	b.backendMu.Lock()
	defer b.backendMu.Unlock()

	b.scraper.Stop()
	b.node.Stop()
	return b.transactor.Stop()
}

*/
