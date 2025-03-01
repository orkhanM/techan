package techan

import "github.com/sdcoffey/big"

// Cumulative Typical Price
type ctpIndicator struct {
	high  Indicator
	low   Indicator
	close Indicator
}

// NewCtpIndicator returns an Indicator which returns the Cumulative Typical Price for the last
// candle in our series, the cumulative typical price is the average of the high, low, and close prices
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
