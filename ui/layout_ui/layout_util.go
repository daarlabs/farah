package layout_ui

import (
	"github.com/daarlabs/farah/component/dark_mode_switcher_component"
	"github.com/daarlabs/hirokit/hiro"
)

func Dark(c hiro.Ctx) bool {
	return c.Cookie().Get(dark_mode_switcher_component.DarkModeCookieKey) == "true"
}
