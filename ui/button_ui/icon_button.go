package button_ui

import (
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func IconButton(props Props, nodes ...Node) Node {
	if len(props.Type) == 0 {
		props.Type = "button"
	}
	if len(props.Size) == 0 {
		props.Size = tempest.SizeMain
	}
	return Button(
		tempest.Class().Transition().Rounded().Grid().PlaceItemsCenter().
			Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()).
			Bg(palette.Primary, 200, tempest.Hover()).Bg(palette.Primary, 300, tempest.Dark(), tempest.Hover()).
			If(props.Size == tempest.SizeMain, tempest.Class().Size(10).Px(3)).
			If(props.Size == tempest.SizeSm, tempest.Class().Size(8).Px(2)).
			Extend(form_tempest.FocusShadow()),
		Type(props.Type),
		icon_ui.Icon(icon_ui.Props{Icon: props.Icon, Size: ui.Sm, Class: tempest.Class().TextWhite()}),
		Fragment(nodes...),
	)
}
