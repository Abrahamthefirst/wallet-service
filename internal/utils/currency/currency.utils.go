package currency

import (
	"fmt"
	"math"
)

var currencyMinorUnits = map[string]int{
	"NGN": 2,
	"USD": 2,
	"GBP": 2,
	"KWD": 3,
	"BHD": 3,
	"JPY": 0,
}

func ToMinorUnit(amount float64, currency string) (int64, error) {
	decimals, ok := currencyMinorUnits[currency]
	if !ok {
		return 0, fmt.Errorf("unsupported currency: %s", currency)
	}
	multiplier := math.Pow(10, float64(decimals))
	return int64(math.Round(amount * multiplier)), nil
}
