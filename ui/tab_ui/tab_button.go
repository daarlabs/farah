package tab_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
)

func TabButton(link, title string, active bool) Node {
	return A(
		tempest.Class().Transition().Px(2.5).Py(1.5).Rounded().TextSize("10px").
			Border(1).
			If(
				!active,
				tempest.Class().TextSlate(900).Text(palette.Primary, 100, tempest.Dark()).
					BorderSlate(300).BorderSlate(600, tempest.Dark()).
					BorderColor(palette.Primary, 400, tempest.Hover()).
					BorderColor(palette.Primary, 100, tempest.Hover(), tempest.Dark()),
			).
			If(
				active,
				tempest.Class().TextWhite().
					Bg(palette.Primary, 400).
					Bg(palette.Primary, 200, tempest.Dark()).
					Bg(palette.Primary, 500, tempest.Hover()).
					Bg(palette.Primary, 300, tempest.Dark(), tempest.Hover()),
			),
		Href(link),
		Text(title),
	)
}
