package captcha_component

import (
	"time"
	
	"github.com/daarlabs/hirokit/hiro"
)

const (
	captchaCacheKey = "captcha"
)

func Valid(c hiro.Ctx) (bool, error) {
	if c.Request().Is().Get() {
		valid := true
		if err := c.State().Get(captchaCacheKey, &valid); err != nil {
			return valid, err
		}
		return valid, nil
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
	return isValid, c.State().Save(captchaCacheKey, isValid)
}

func MustValid(c hiro.Ctx) bool {
	valid, err := Valid(c)
	if err != nil {
		panic(err)
	}
	return valid
}

func save(c hiro.Ctx, challengeToken, successToken string) error {
	return c.Cache().Set(captchaCacheKey+":"+challengeToken, captcha{SuccessToken: successToken}, 5*time.Minute)
}
