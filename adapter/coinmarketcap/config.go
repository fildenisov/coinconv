package coinmarketcap

import "time"

type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}
