package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/ui/menu_ui"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func NavMenuHandler(nodes ...Node) Node {
	return Button(
		Type("button"),
		menu_ui.Open(),
		tempest.Class().Flex().ItemsCenter().Gap(1),
		Fragment(nodes...),
		Div(
			tempest.Class().Transition().Rotate(-180, tempest.Hover(tempest.Group)),
			menu_ui.Chevron(),
			icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.ChevronDown, Size: ui.Sm,
					Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()).
						Text(palette.Primary, 400, tempest.Hover(tempest.Group)).
						Text(palette.Primary, 100, tempest.Hover(tempest.Group), tempest.Dark()),
				},
			),
		),
	)
}
