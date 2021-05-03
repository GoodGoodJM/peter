package database

import (
	"github.com/goodgoodjm/peter/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("peter.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	models.Migrate(db)
	return db, err
}
