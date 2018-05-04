package kowalad

import "github.com/kowala-tech/kcoin/log"

type API struct {
	backend Backend
	log     log.Logger
}

func NewAPI() *API {
	return &API{
		backend: NewBackend(),
		log:     log.New("package", "kowalad/api"),
	}
}

func (api *API) StartNode(config *Config) error {
	return api.backend.StartNode(config)
}

func (api *API) StopNode() error {
	return api.backend.StopNode()
}

func (api *API) SendRawTransaction(data []byte) error {
	return api.backend.SendRawTransaction(data)
}
