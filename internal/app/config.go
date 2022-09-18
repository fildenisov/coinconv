package app

import "time"

// Config is an app config
type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}
