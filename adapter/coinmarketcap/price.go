package coinmarketcap

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	cmc "github.com/fildenisov/coinconv/src/client/coinmarketcap"
)

func (c *Client) ConvertPrice(ctx context.Context, amount float32, symbol, convert string) (float32, error) {
	resp, httpResp, err := c.c.ToolsApi.GetV2ToolsPriceconversion(ctx, amount, &cmc.ToolsApiGetV2ToolsPriceconversionOpts{
		Symbol:  optional.NewString(symbol),
		Convert: optional.NewString(convert),
	})

	if err != nil {
		return 0, fmt.Errorf("Got error on conversion request: %v", err)
	}

	if httpResp.StatusCode != 200 {
		return 0, fmt.Errorf("Got %v status got on request", httpResp.StatusCode)
	}

	if resp.Status != nil {
		if resp.Status.ErrorCode != 0 || resp.Status.ErrorMessage != "" {
			return 0, fmt.Errorf("Error! Code: %v, message %v", resp.Status.ErrorCode, resp.Status.ErrorMessage)
		}
	}

	symbolData, ok := resp.Data[symbol]
	if !ok {
		return 0, fmt.Errorf("No symbol data returned for %v", symbol)
	}

	quoteData, ok := symbolData.Quote[convert]
	if !ok {
		return 0, fmt.Errorf("No quote data returned for %v", convert)
	}

	return quoteData.Price, nil
}
