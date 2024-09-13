package menu_item_ui

import (
	"github.com/daarlabs/farah/palette"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func MenuItem(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Transition().Px(4).Py(2).TextXs().CursorPointer().
			Flex().ItemsCenter().Gap(2).
			BorderB(1).BorderSlate(200).BorderSlate(700, tempest.Dark()).
			If(
				!props.Selected,
				tempest.Class().TextSlate(800).TextWhite(tempest.Dark()).BgWhite().BgSlate(800, tempest.Dark()).
					BgSlate(100, tempest.Hover()).BgSlate(700, tempest.Dark(), tempest.Hover()),
			).
			If(
				props.Selected,
				tempest.Class().TextWhite().Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()),
			),
		Fragment(nodes...),
	)
}
