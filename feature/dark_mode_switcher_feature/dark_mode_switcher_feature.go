package dark_mode_switcher_feature

import (
	"strconv"
	"time"
	
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

type DarkModeSwtitcher struct {
	mirage.Component
	dark bool
}

const (
	DarkModeCookieKey = "X-Dark-Mode"
)

var (
	darkModeCookieDuration = 30 * 24 * time.Hour
)

func (c *DarkModeSwtitcher) Name() string {
	return "dark-mode-switcher"
}

func (c *DarkModeSwtitcher) Mount() {
	c.dark = c.Cookie().Get(DarkModeCookieKey) == "true"
}

func (c *DarkModeSwtitcher) Node() Node {
	icon := icon_ui.Sun
	if c.dark {
		icon = icon_ui.Moon
	}
	return Div(
		Class("inline-flex"),
		A(
			Class("block"),
			Href(c.Generate().Action("HandleSwitch")),
			icon_ui.Icon(
				icon_ui.Props{
					Icon:  icon,
					Size:  ui.Sm,
					Class: "text-slate-900 dark:text-white",
				},
			),
		),
	)
}

func (c *DarkModeSwtitcher) HandleSwitch() error {
	c.Cookie().Set(DarkModeCookieKey, strconv.FormatBool(!c.dark), darkModeCookieDuration)
	return c.Response().Refresh()
}
