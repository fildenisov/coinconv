package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fildenisov/coinconv/adapter/coinmarketcap"
	"github.com/fildenisov/coinconv/internal/app"
)

const (
	defaultBaseURL = "https://sandbox-api.coinmarketcap.com"
	defaultAPIKey  = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c" // nolint: gosec
	defaultTimeout = 10 * time.Second

	baseURLEnv = "COINCONV_BASE_URL"
	apiKeyEnv  = "COINCONV_API_KEY"
)

var (
	baseURL string
	apiKey  string
	amount  float32
	from    string
	to      string
)

func init() {
	if len(os.Args[1:]) < 3 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	amountParsed, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		fmt.Println("first argument must be a valid float32")
		os.Exit(1)
	}
	amount = float32(amountParsed)

	from = os.Args[2]
	to = os.Args[3]

	baseURL = os.Getenv(baseURLEnv)
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	apiKey = os.Getenv(apiKeyEnv)
	if apiKey == "" {
		apiKey = defaultAPIKey
	}
}

func main() {
	a := app.New(app.Config{
		CMC: coinmarketcap.Config{
			BaseURL: baseURL,
			Timeout: defaultTimeout,
			APIKey:  apiKey,
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := a.Converter.ConvertPrice(ctx, amount, from, to)
	cancel()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
