package enums

type Currency string

const (
	CurrencyNGN Currency = "NGN"
	CurrencyGHS Currency = "GHS"
	CurrencyUSD Currency = "USD"
)

func (c Currency) String() string {
	return string(c)
}
