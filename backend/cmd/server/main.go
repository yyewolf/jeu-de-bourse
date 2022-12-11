package main

import (
	"jeu-de-bourse/api/v1/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	router.Route(app)

	app.Run(":8080")
}
