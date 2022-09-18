package coinmarketcap

import (
	"net/http"

	cmc "github.com/fildenisov/coinconv/src/client/coinmarketcap"
)

const (
	userAgent = "go-coinconv"
	keyHeader = "X-CMC_PRO_API_KEY"
)

type Client struct {
	c *cmc.APIClient
}

func New(cfg Config) *Client {

	return &Client{
		c: cmc.NewAPIClient(&cmc.Configuration{
			BasePath: cfg.BaseURL,
			DefaultHeader: map[string]string{
				keyHeader: cfg.APIKey,
			},
			UserAgent: userAgent,
			HTTPClient: &http.Client{
				Timeout: cfg.Timeout,
			},
		}),
	}
}
