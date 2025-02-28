package techan

import "github.com/sdcoffey/big"

// Cumulative Typical Price
type ctpIndicator struct {
	high  Indicator
	low   Indicator
	close Indicator
}

func NewCtpIndicator(high, low, close Indicator) Indicator {
	return ctpIndicator{
		high:  high,
		low:   low,
		close: close,
	}
}

func (v ctpIndicator) Calculate(index int) big.Decimal {
	high := v.high.Calculate(index)
	low := v.low.Calculate(index)
	close := v.close.Calculate(index)

	return high.Add(low).Add(close).Div(big.NewFromInt(3))
}
