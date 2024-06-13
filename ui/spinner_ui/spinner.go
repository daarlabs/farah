package spinner_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
)

func Spinner(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Transition().Absolute().Inset(0).M("auto").Grid().PlaceItemsCenter().
			BgWhite(tempest.Opacity(0.7)).BgSlate(800, tempest.Dark(), tempest.Opacity(0.7)).
			If(props.Class != nil, props.Class),
		Div(
			tempest.Class().Spin().Size(5).RoundedFull().Border(3).BorderWhite().BorderSlate(600, tempest.Dark()).
				BorderTColor(palette.Primary, 400).BorderTColor(palette.Primary, 100, tempest.Dark()),
		),
		Fragment(nodes...),
	)
}
