package coinmarketcap

import "time"

// Config is an coinmarketcap adapter config
type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}
