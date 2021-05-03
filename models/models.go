package models

import "time"

type Ticker struct {
	ID       uint   `json:"id"`
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Name     string `json:"name"`
}

type Registration struct {
	ID          uint
	Symbol      string
	PortfolioID uint
	CreatedAt   time.Time
}

type RegistrationGroup struct {
	Symbol string `json:"symbol"`
	Count  uint   `json:"count"`
}

type Price struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	CreatedAt time.Time
}
