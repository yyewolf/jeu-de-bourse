package router

import (
	"jeu-de-bourse/api/v1/handlers"
	"jeu-de-bourse/api/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	path := engine.Group("/api/v1")

	path.POST("/captcha", middlewares.VerifyCaptcha(), handlers.TestCaptcha)
}
