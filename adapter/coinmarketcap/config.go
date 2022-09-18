package coinmarketcap

import "time"

type Config struct {
	BaseURL string
	Timeout time.Duration
	APIKey  string
}
