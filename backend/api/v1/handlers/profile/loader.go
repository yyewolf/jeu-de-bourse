package profile

import (
	"jeu-de-bourse/api/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.RouterGroup) {
	sg := r.Group("/profile")

	sg.GET("/infos", middlewares.AuthRequired(), GetProfileInfos)
}
