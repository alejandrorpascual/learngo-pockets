package money

import (
	"errors"
	"testing"
)

func TestParseCurrency_Success(t *testing.T) {
	tt := map[string]struct {
		code     string
		expected Currency
	}{
		"majority EUR":   {code: "EUR", expected: Currency{code: "EUR", precision: 2}},
		"thousandth BHD": {code: "BHD", expected: Currency{code: "BHD", precision: 3}},
		"tenth VND":      {code: "VND", expected: Currency{code: "VND", precision: 1}},
		"integer IRR":    {code: "IRR", expected: Currency{code: "IRR", precision: 0}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.code)
			if err != nil {
				t.Errorf("expected no error, got %s", err.Error())
			}

			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}

}

func TestParseCurrency_UnknownCurrency(t *testing.T) {
	_, err := ParseCurrency("INVALID")
	if !errors.Is(err, ErrInvalidCurrencyCode) {
		t.Errorf("expected error %s, got %v", ErrInvalidCurrencyCode, err)
	}
}
