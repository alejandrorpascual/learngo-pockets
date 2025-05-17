package money

// Currency defines the code of the currency.
type Currency struct {
	code      string
	precision byte
}

// ErrInvalidCurrencyCode is returned when the currency to parse is not a
// standard 3-letter code.
const ErrInvalidCurrencyCode = Error("invalid currency code")

// ParseCurrency returns the currency associated to a name and may return
// ErrInvalidCurrencyCode
func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}
