package auth

import (
	"jeu-de-bourse/api/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.RouterGroup) {
	sg := r.Group("/auth")

	sg.POST("/register", middlewares.DeauthRequired(), middlewares.VerifyCaptcha(), Register)
	sg.POST("/login", middlewares.DeauthRequired(), middlewares.VerifyCaptcha(), Login)
	sg.GET("/logout", middlewares.AuthRequired(), Logout)
	sg.GET("/debug", middlewares.AuthRequired(), Connected)
}
