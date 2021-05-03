package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Ticker{},
		&Registration{},
		&Price{},
	)

	if err != nil {
		panic(err)
	}
}
