package batch

import (
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddJob("*/5 * * * * *", fetchQuotesJob{db})

	if err != nil {
		panic(err)
	}

	c.Run()
}
