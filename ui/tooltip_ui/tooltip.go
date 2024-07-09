package tooltip_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Tooltip(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Transition().
			Absolute().Inset(0).M("auto").Z(9999).
			Hidden().Grid(tempest.Hover(tempest.Group)).PlaceItemsCenter().TextCenter().
			BgSlate(900, tempest.Opacity(80)).Rounded().
			TextSize("10px").TextWhite(),
		Fragment(nodes...),
	)
}

func Wrapper() tempest.Tempest {
	return tempest.Class().Relative().Group()
}
