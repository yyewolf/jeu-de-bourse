package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
