package batch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	QuoteResponse struct {
		Quotes []Quote `json:"result"`
		Error  *Error  `json:"error"`
	} `json:"quoteResponse"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

type Quote struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"regularMarketPrice"`
}

func FetchQuotes(symbols string) (_ []Quote, err error) {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&symbols=%s", symbols)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = res.Body.Close()
	}()

	body := response{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}

	if err := body.QuoteResponse.Error; err != nil {
		return nil, err
	}

	return body.QuoteResponse.Quotes, nil
}
