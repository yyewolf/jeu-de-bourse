package stocks

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.RouterGroup) {
	sg := r.Group("/stocks")

	sg.POST("/", GetStocksPage)
}
