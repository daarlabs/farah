package web_layout_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	
	"github.com/daarlabs/farah/ui/menu_ui"
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
	
	"github.com/daarlabs/farah/feature/dark_mode_switcher_feature"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/layout_ui"
	"github.com/daarlabs/farah/ui/nav_ui"
	"github.com/daarlabs/farah/ui/nav_ui/nav_section_ui"
)

func Layout(c mirage.Ctx, nodes ...Node) Node {
	return layout_ui.Layout(
		c,
		Div(
			Class("grid grid-rows-[64px_1fr]"),
			nav_ui.NavBar(
				nav_ui.NavImgLink(c.Generate().Link("home"), c.Generate().PublicUrl("img/svg/logo.svg")),
				nav_section_ui.NavSection(
					nav_section_ui.Props{},
					nav_ui.NavLink(c.Generate().Link("ui"), Text("UI")),
				),
				nav_section_ui.NavSection(
					nav_section_ui.Props{},
					nav_ui.NavLink(c.Generate().Link("feature"), Text("Feature")),
				),
				nav_section_ui.NavSection(
					nav_section_ui.Props{},
					menu_ui.Menu(
						menu_ui.Props{},
						nav_ui.NavLink(c.Generate().Link("home"), nav_ui.NavMenuHandler(Text("Menu 1"))),
						menu_item_ui.MenuItem(menu_item_ui.Props{}, A(Href(c.Generate().Link("home")), Text("Menu item 1"))),
					),
				),
				nav_section_ui.NavSection(
					nav_section_ui.Props{AlignX: ui.Right},
					c.Create().Component(&dark_mode_switcher_feature.DarkModeSwtitcher{}),
				),
			),
			Div(
				Class("w-screen h-screen overflow-y-auto overflow-x-hidden"),
				Fragment(nodes...),
			),
		),
	)
}
