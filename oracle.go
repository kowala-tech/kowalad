package oracled

import (
	"math/big"
	"time"

	
	contract "github.com/kowala-tech/kcoin/contracts/oracle"
	"github.com/kowala-tech/kcoin/event"
	"github.com/kowala-tech/kcoin/kcoin"
	"github.com/kowala-tech/kcoin/log"
)

const (
	priceChanSize = 1024
)

type Oracle interface {
	Start() error
	Stop()
}

type oracle struct {
	contract.OracleManagerContract
	node Node
	scraper Scraper
	log log.Logger
	
	priceCh  chan *big.Int
	priceSub event.Subscription
}

func NewOracle(node Node) *oracleManager {
	return &oracleManager{
		node: node,
		log: log.New("package", "oracled/oracle_manager"),
		scraper: NewTrustedScraper(),
	}
}

func (oracle *oracle) Register() 


func (oracle *oracle) Start() error {
	if err := manager.generateContractBindings(); err != nil {
		// @TODO (rgeraldes) - log error
		return err
	}

	// check if the user is already an oracle (process crash)
	isOracle := false

	if !isOracle {
		// register oracle
		if err := manager.contract.RegisterOracle(nil,0); err != nil 
		if err != nil {
			return err
		}
	}

	// we use the blockchain as the persistence in case of a node crash
	// we just need to check the latest transaction and period to initialize the scraper
	/* Initialize Scraper */
	// @TODO (rgeraldes) - contract contains sync period
	period := time.Minute
	manager.scraper.Start(period)

	// submit authenticated prices
	manager.priceCh = make(chan *big.Int, priceChanSize)
	manager.priceSub = manager.scraper.SubscribeFeeds(manager.priceCh)
	go manager.priceBroadcastLoop()

	// initialize the scraper
	manager.scraper.Start(period)

	_, err := tx.RegisterOracle(&bind.TransactOpts)
	if err != nil {
		tx.log.Error("")
	}	
}

func (manager *oracleManager) Stop() {

	manager.priceSub.Unsubscribe()
	manager.contract.DeregisterOracle(&bind.TransactOpts{})
}

func (manager *oracleManager) priceBroadcastLoop() {
	for {
		select {
		case data := <-manager.priceCh:
			manager.broadcastPrice(data)
		case <-manager.priceSub.Err(): // Unsubscribe
			return
		}
	}
}

func (manager *oracleManager) broadcastPrice(result *big.Int) {
	_, err := manager.contract.SubmitPrice(result)
	if err != nil {
		return
	}
}


func (oracle *oracle) generateContractBindings() error {
	service, err := oracle.node.LiteKowalaService()
	if err != nil {
		return err
	}
	bindings, err := oracle.NewOracleManagerContract(kcoin.NewContractBackend(service.ApiBackend), new(big.Int).SetUint64(service.NetVersion()))
	if err != nil {
		return err
	}
	manager.contract = bindings
}