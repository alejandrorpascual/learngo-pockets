package money

import "math"

// Convert applies the change rate to convert an amount to target currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	rate := ExchangeRate{2, 0}
	converedtValue := applyExchangeRate(amount, to, rate)
	if err := converedtValue.validate(); err != nil {
		return Amount{}, err
	}
	return converedtValue, nil
}

// ExchangeRate represents a rate to convert from one currency to another
type ExchangeRate Decimal

// applyExchangeRate returns  new Amount representing the input multiplied
// by the rate.
// The precision of the returned value is that od the target currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted := multiply(a.quantity, rate)
	// if err != nil {
	// 	return Amount{}
	// }

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:

		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}
}

func multiply(d Decimal, rate ExchangeRate) Decimal {
	dec := Decimal{
		subunits:  d.subunits * rate.subunits,
		precision: d.precision + rate.precision,
	}

	dec.simplify()

	return dec
}

// pow10 is a quick implementation of how to raide 10 to a given power.
// It's optimized for small powers, and slow for unusually high powers.
func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}
