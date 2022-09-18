package coinmarketcap

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	cmc "github.com/fildenisov/coinconv/src/client/coinmarketcap"
)

func (c *Client) ConvertPrice(ctx context.Context, amount float32, from, to string) (float32, error) {
	resp, httpResp, err := c.c.ToolsApi.GetV2ToolsPriceconversion(
		ctx,
		amount,
		&cmc.ToolsApiGetV2ToolsPriceconversionOpts{
			Symbol:  optional.NewString(from),
			Convert: optional.NewString(to),
		})

	if err != nil {
		return 0, fmt.Errorf("got error on conversion request: %v", err)
	}

	if httpResp.StatusCode != 200 {
		return 0, fmt.Errorf("got %v status got on request", httpResp.StatusCode)
	}

	if resp.Status != nil {
		if resp.Status.ErrorCode != 0 || resp.Status.ErrorMessage != "" {
			return 0, fmt.Errorf(
				"error! Code: %v, message %v",
				resp.Status.ErrorCode,
				resp.Status.ErrorMessage,
			)
		}
	}

	symbolData, ok := resp.Data[from]
	if !ok {
		return 0, fmt.Errorf("no symbol data returned for %v", from)
	}

	quoteData, ok := symbolData.Quote[to]
	if !ok {
		return 0, fmt.Errorf("no quote data returned for %v", to)
	}

	return quoteData.Price, nil
}
