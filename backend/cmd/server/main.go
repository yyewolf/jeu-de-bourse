package main

import (
	"jeu-de-bourse/api/v1/router"
	"jeu-de-bourse/internal/database"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Route(app)

	database.Session.Close()

	app.Run(":8080")
}
