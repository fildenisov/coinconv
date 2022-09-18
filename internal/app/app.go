package app

import (
	"github.com/fildenisov/coinconv/adapter/coinmarketcap"
	"github.com/fildenisov/coinconv/internal/rep"
)

// App is an intancee of the main application
type App struct {
	Converter rep.Converter
}

// New is a constructor for an App
func New(cfg Config) *App {
	return &App{
		Converter: coinmarketcap.New(coinmarketcap.Config{
			BaseURL: cfg.BaseURL,
			APIKey:  cfg.APIKey,
			Timeout: cfg.Timeout,
		}),
	}
}
