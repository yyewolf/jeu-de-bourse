package auth

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.RouterGroup) {
	r.Group("/auth")

	r.POST("/login", Login)
}
