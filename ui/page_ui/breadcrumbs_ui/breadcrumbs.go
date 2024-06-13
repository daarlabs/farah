package breadcrumbs_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Breadcrumbs(mainLink string, nodes ...Node) Node {
	return Div(
		tempest.Class().Flex().ItemsCenter().Gap(2),
		A(
			Href(mainLink), icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.Home, Size: ui.Sm,
					Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()).
						Text(palette.Primary, 400, tempest.Hover()).
						Text(palette.Primary, 100, tempest.Dark(), tempest.Hover()),
				},
			),
		),
		Fragment(nodes...),
	)
}
