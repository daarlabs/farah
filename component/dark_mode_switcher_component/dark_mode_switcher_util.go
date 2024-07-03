package dark_mode_switcher_component

import "github.com/daarlabs/hirokit/hiro"

func Dark(c hiro.Ctx) bool {
	return c.Cookie().Get(DarkModeCookieKey) == "true"
}
