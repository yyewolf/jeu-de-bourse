package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "captcha passed",
	})
}
