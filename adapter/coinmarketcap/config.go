package coinmarketcap

import "time"

// Config is a coinmarketcap adapter config
type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}
