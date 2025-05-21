package ecbank

import (
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
)

// Client can call the bank to retrieve exchange rates.
type Client struct {
	url string
}

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrClientSide         = ecbankError("client side error when contacting ECB")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
	ErrServerSide         = ecbankError("server side error when contacting ECB")
	ErrUnexpectedFormat   = ecbankError("unexpected response format")
	ErrUnknownStatusCode  = ecbankError("unknown status code contacting ECB")
)

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const euroxrefURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	if c.url == "" {
		c.url = euroxrefURL
	}

	resp, err := http.Get(c.url)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrServerSide, err.Error())
	}

	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err

	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil

}

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}
