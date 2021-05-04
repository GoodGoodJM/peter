package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goodgoodjm/peter/api"
	"github.com/goodgoodjm/peter/batch"
	"github.com/goodgoodjm/peter/database"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}

	go batch.Start(db)

	app := gin.Default()
	app.Use(database.Inject(db))
	api.Apply(app)

	if err = app.Run(); err != nil {
		panic(err)
	}
}
