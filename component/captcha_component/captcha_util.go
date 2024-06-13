package captcha_component

import (
	"time"
	
	"github.com/daarlabs/arcanum/mirage"
)

const (
	captchaCacheKey = "captcha"
)

func Valid(c mirage.Ctx) (bool, error) {
	if c.Request().Is().Get() {
		return true, nil
	}
	var r captcha
	challengeToken := c.Request().Form().Get(captchaChallenge)
	successToken := c.Request().Form().Get(captchaSuccess)
	if err := c.Cache().Get(captchaCacheKey+":"+challengeToken, &r); err != nil {
		return false, err
	}
	isValid := len(r.SuccessToken) > 0 && len(successToken) > 0 && r.SuccessToken == successToken
	if err := c.Cache().Set(captchaCacheKey+":"+challengeToken, "", time.Millisecond); err != nil {
		return isValid, err
	}
	return isValid, nil
}

func MustValid(c mirage.Ctx) bool {
	valid, err := Valid(c)
	if err != nil {
		panic(err)
	}
	return valid
}

func save(c mirage.Ctx, challengeToken, successToken string) error {
	return c.Cache().Set(captchaCacheKey+":"+challengeToken, captcha{SuccessToken: successToken}, 5*time.Minute)
}
