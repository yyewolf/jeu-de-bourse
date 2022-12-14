package verifier

import (
	"fmt"
	"jeu-de-bourse/internal/env"
	"time"

	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

var session *recaptcha.ReCAPTCHA

func init() {
	captcha, err := recaptcha.NewReCAPTCHA(env.GlobalConfig.RecaptchaSecret, recaptcha.V2, 10*time.Second)
	if err != nil {
		panic(err)
	}

	session = &captcha
}

func Verify(token string) bool {
	err := session.VerifyWithOptions(token, recaptcha.VerifyOption{})
	fmt.Println(err)
	return err == nil
}
