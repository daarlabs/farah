package dark_mode_switcher_component

import (
	"strconv"
	"time"
	
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

type DarkModeSwtitcher struct {
	hiro.Component
	Menu bool `json:"-"`
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
	c.dark = Dark(c)
}

func (c *DarkModeSwtitcher) Node() Node {
	icon := icon_ui.Sun
	if c.dark {
		icon = icon_ui.Moon
	}
	if c.Menu {
		return c.createMenuItemLink(icon)
	}
	return c.createIconLink(icon)
}

func (c *DarkModeSwtitcher) createMenuItemLink(icon string) Node {
	return A(
		Href(c.Generate().Action("HandleSwitch")),
		menu_item_ui.MenuItem(
			menu_item_ui.Props{},
			Div(
				tempest.Class().Flex().ItemsCenter().Gap(2),
				c.createIcon(icon),
				Span(Text(c.getMenuItemTitle())),
			),
		),
	)
}

func (c *DarkModeSwtitcher) createIconLink(icon string) Node {
	return Div(
		tempest.Class().InlineFlex(),
		A(
			tempest.Class().Block().Size(4),
			Href(c.Generate().Action("HandleSwitch")),
			c.createIcon(icon),
		),
	)
}

func (c *DarkModeSwtitcher) createIcon(icon string) Node {
	return icon_ui.Icon(
		icon_ui.Props{
			Icon:  icon,
			Size:  ui.Sm,
			Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
		},
	)
}

func (c *DarkModeSwtitcher) HandleSwitch() error {
	c.Cookie().Set(DarkModeCookieKey, strconv.FormatBool(!c.dark), darkModeCookieDuration)
	return c.Response().Redirect(c.Generate().Current())
}

func (c *DarkModeSwtitcher) getMenuItemTitle() string {
	if c.dark {
		return c.Translate("component.dark-mode-switcher.dark")
	}
	return c.Translate("component.dark-mode-switcher.light")
}
