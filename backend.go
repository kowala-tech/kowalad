package kowalad

import (
	"context"
	"fmt"
	"time"

	"github.com/kowala-tech/kcoin/kcoinclient"
	"github.com/kowala-tech/kcoin/log"
)

const (
	timeout = 300 * time.Second
)

type Backend interface {
	StartNode(config *Config) error
	StopNode() error
	SendRawTransaction(data []byte) error
}

type backend struct {
	config     *Config
	node       Node
	client     *kcoinclient.Client
	log        log.Logger
	rpcTimeout time.Duration
}

func NewBackend() *backend {
	b := &backend{
		node: NewLightNode(),
		log:  log.New("package", "kowalad/backend"),
	}

	return b
}

func (b *backend) StartNode(config *Config) error {
	b.config = config
	return b.startNode(config)
}

func (b *backend) startNode(config *Config) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("node crashed on start: %v", err)
		}
	}()

	b.client, err = b.node.Client()
	if err != nil {
		log.Error("Failed to get the client", "err", err)
		return err
	}

	return b.node.Start(config)
}

func (b *backend) StopNode() error {
	return b.node.Stop()
}

func (b *backend) SendRawTransaction(data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), b.rpcTimeout)
	defer cancel()
	return b.client.SendRawTransaction(ctx, data)
}
