package middlewares

import (
	"jeu-de-bourse/internal/verifier"

	"github.com/gin-gonic/gin"
)

type CaptchaReq struct {
	CaptchaToken string `json:"captchaToken" binding:"required"`
}

func VerifyCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user token is always in the post form as "captchaToken"
		captcha := CaptchaReq{}
		c.ShouldBindJSON(&captcha)

		pass := verifier.Verify(captcha.CaptchaToken)

		if !pass {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "captcha failed",
			})
			return
		}
	}
}
