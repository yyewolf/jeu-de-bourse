package router

import (
	"jeu-de-bourse/api/v1/handlers/auth"
	"jeu-de-bourse/api/v1/handlers/profile"
	"jeu-de-bourse/api/v1/handlers/stocks"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	path := engine.Group("/api/v1")

	auth.LoadRoutes(path)
	stocks.LoadRoutes(path)
	profile.LoadRoutes(path)
}
