package dark_mode_switcher_component

import "github.com/daarlabs/arcanum/mirage"

func Dark(c mirage.Ctx) bool {
	return c.Cookie().Get(DarkModeCookieKey) == "true"
}
