package rep

import "context"

// Converter is an component to convert currency rates.
type Converter interface {
	ConvertPrice(ctx context.Context, amount float32, symbol, convert string) (float32, error)
}
