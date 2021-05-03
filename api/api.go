package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/goodgoodjm/peter/models"
	"gorm.io/gorm"
)

func Apply(r *gin.Engine) {
	r.GET("/helloa", func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		var registrationsGroups []models.RegistrationGroup
		exec := db.
			Model(&models.Registration{}).
			Select("ticker, count(*) as count").
			Group("ticker").
			Order("count DESC").
			Find(&registrationsGroups)

		if err := exec.Error; err != nil {
			c.AbortWithStatus(500)
			log.Fatal(err.Error())
			return

		}
		c.JSON(200, registrationsGroups)
	})
}
