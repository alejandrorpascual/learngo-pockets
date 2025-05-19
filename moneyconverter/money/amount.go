package money

type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the
	// currency
	ErrTooPrecise = Error("quantity is too precise")
)

// NewAmount return an amount of money
func NewAmount(qnt Decimal, curr Currency) (Amount, error) {
	if qnt.precision > curr.precision {
		return Amount{}, ErrTooPrecise
	}

	qnt.precision = curr.precision

	return Amount{quantity: qnt, currency: curr}, nil

}

func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}
