package main

import (
	"jeu-de-bourse/api/v1/router"
	"jeu-de-bourse/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	router.Route(app)

	database.Session.Close()

	app.Run(":8080")
}
