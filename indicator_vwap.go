package techan

import "github.com/sdcoffey/big"

type vwapIndicator struct {
	volume Indicator
	ctp    Indicator
}

// NewVwapIndicator returns an Indicator which returns the volume weighted average price of a candle for a given index
// It expects two indicators, one for volume and one for close price / cumulative trade price
func NewVwapIndicator(volume, ctp Indicator) Indicator {
	return vwapIndicator{
		volume,
		ctp,
	}
}

func (v vwapIndicator) Calculate(index int) big.Decimal {
	numerator := big.ZERO
	denominator := big.ZERO
	for i := 0; i <= index; i++ {
		volume := v.volume.Calculate(i)
		numerator = numerator.Add(volume.Mul(v.ctp.Calculate(i)))
		denominator = denominator.Add(volume)
	}

	if denominator.IsZero() {
		return big.ZERO
	}

	result := numerator.Div(denominator)

	return result
}
