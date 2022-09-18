package app

import (
	"github.com/fildenisov/coinconv/adapter/coinmarketcap"
	"github.com/fildenisov/coinconv/internal/rep"
)

type App struct {
	Converter rep.Converter
}

func New(cfg Config) *App {
	return &App{
		Converter: coinmarketcap.New(cfg.CMC),
	}
}
