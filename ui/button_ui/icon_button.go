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
	return Button(
		tempest.Class().Transition().Size(10).Rounded().Grid().PlaceItemsCenter().
			Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()).
			Bg(palette.Primary, 200, tempest.Hover()).Bg(palette.Primary, 300, tempest.Dark(), tempest.Hover()).
			Extend(form_tempest.FocusShadow()),
		Type(props.Type),
		icon_ui.Icon(icon_ui.Props{Icon: props.Icon, Size: ui.Sm, Class: tempest.Class().TextWhite()}),
		Fragment(nodes...),
	)
}
