package middlewares

import (
	"jeu-de-bourse/internal/verifier"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CaptchaReq struct {
	CaptchaToken string `json:"captchaToken" binding:"required"`
}

func VerifyCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user token is always in the post form as "captchaToken"
		captcha := CaptchaReq{}
		c.ShouldBindBodyWith(&captcha, binding.JSON)

		pass := verifier.Verify(captcha.CaptchaToken)

		if !pass {
			c.AbortWithStatusJSON(400, gin.H{
				"success": false,
				"status":  "Captcha invalide",
			})
			return
		}
	}
}
