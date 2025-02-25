package techan

import "github.com/sdcoffey/big"

// import "fmt"

type vwapIndicator struct {
	*TimeSeries
}

// NewBollingerUpperBandIndicator a a derivative indicator which returns the upper bound of a bollinger band
// on the underlying indicator
func NewVwapIndicator(series *TimeSeries) Indicator {
	return vwapIndicator{
		series,
	}
}

// func (bbi vwapIndicator) Calculate(index int) big.Decimal {
// 	numerator := big.NewDecimal(0)
// 	denominator := big.NewDecimal(0)
// 	for candle := range bbi.Candles[:index] {
// 		fmt.Println("In vwapcalculate, numbers are", bbi.Candles[candle].ClosePrice, bbi.Candles[candle].MaxPrice, bbi.Candles[candle].MinPrice, bbi.Candles[candle].Volume)
// 		cumulativeTypicalPrice := bbi.Candles[candle].ClosePrice.Add(bbi.Candles[candle].MaxPrice).Add(bbi.Candles[candle].MinPrice).Div(big.NewDecimal(3.0))
// 		numerator = numerator.Add(cumulativeTypicalPrice.Mul(bbi.Candles[candle].Volume))
// 		denominator = denominator.Add(bbi.Candles[candle].Volume)
// 		fmt.Println("In vwapcalculate, numerator is", numerator, "denominator is", denominator)
// 	}
// 	fmt.Println("In vwapcalculate, numerator is", numerator, "denominator is", denominator)
//
// 	return numerator.Div(denominator)
// }

func (bbi vwapIndicator) Calculate(index int) big.Decimal {
	numerator := big.ZERO
	denominator := big.ZERO
	result := big.ZERO

	for _, candle := range bbi.Candles[index:] {
		numerator = numerator.Add(candle.ClosePrice.Add(candle.MaxPrice).Add(candle.MinPrice).Div(big.NewDecimal(3.0)).Mul(candle.Volume))
		denominator = denominator.Add(candle.Volume)
	}

	result = numerator.Div(denominator)
	// fmt.Println("In Calculate, result is", result)

	return result
}
