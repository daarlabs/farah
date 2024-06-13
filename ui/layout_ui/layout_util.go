package layout_ui

import (
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/farah/component/dark_mode_switcher_component"
)

func Dark(c mirage.Ctx) bool {
	return c.Cookie().Get(dark_mode_switcher_component.DarkModeCookieKey) == "true"
}
