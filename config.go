package oracled

const (
	// config defaults
	MaxPeers          = 15
	MaxPendingPeers   = 5
	NetworkID         = 1
	DataDir           = "kowala-oracle"
	LiteEnabled       = true
	LiteDatabaseCache = 16
)

// Config represents the backend config
type Config struct {
	// MaxPeers is the maximum number of peers that can be connected.
	Maxpeers uint64

	// MaxPendingPeers represents the maximum pending peers that a node has at a time.
	MaxPendingPeers uint64

	// NetworkID is the network that the node will join
	NetworkID uint64

	// DataDir is the default data directory for the oracle executable
	DataDir string

	// Lite represents lowala lite protocol configuration
	Lite ProtocolConfig
}

// ProtocolConfig contains the light client protocol related config
type ProtocolConfig struct {
	// Enabled specifies whether the protocol is enabled
	Enabled bool

	// DatabaseCache internal caching (min 16MB)
	DatabaseCache uint64
}

// NewConfig returns a new configuration, initialized with the default/foundation values
func NewConfig() *Config {
	return &Config{
		Maxpeers:        MaxPeers,
		MaxPendingPeers: MaxPendingPeers,
		NetworkID:       NetworkID,
		DataDir:         DataDir,
		Lite: ProtocolConfig{
			Enabled:       LiteEnabled,
			DatabaseCache: LiteDatabaseCache,
		},
	}
}
