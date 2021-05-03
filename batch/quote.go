package batch

import (
	"encoding/json"
	"fmt"
	"log"
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

func FetchQuotes(symbols string) ([]Quote, error) {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&symbols=%s", symbols)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body := response{}
	json.NewDecoder(res.Body).Decode(&body)

	if err := body.QuoteResponse.Error; err != nil {
		log.Println(err)
		log.Println(body)
		return nil, err
	}

	return body.QuoteResponse.Quotes, nil
}
