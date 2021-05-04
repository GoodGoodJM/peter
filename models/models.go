package models

import "time"

type Ticker struct {
	ID       uint   `json:"id"`
	Symbol   string `json:"symbol" gorm:"type:varchar(11)"`
	Exchange string `json:"exchange" gorm:"type:varchar(11)"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
}

type Registration struct {
	ID          uint
	Symbol      string `gorm:"type:varchar(11)"`
	PortfolioID uint
	CreatedAt   time.Time
}

type RegistrationGroup struct {
	Symbol string `json:"symbol"`
	Count  uint   `json:"count"`
}

type Price struct {
	Symbol    string  `json:"symbol" gorm:"type:varchar(11)"`
	Price     float64 `json:"price"`
	CreatedAt time.Time
}
