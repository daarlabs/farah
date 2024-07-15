package icon_ui

import (
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui"
)

func Icon(props Props) gox.Node {
	if len(props.Size) == 0 {
		props.Size = ui.Main
	}
	return CreateIcon(
		tempest.Class().Transition().FillCurrent().
			If(props.Size == tempest.SizeMain, tempest.Class().Size(6)).
			If(props.Size == tempest.SizeSm, tempest.Class().Size(4)).
			If(props.Size == tempest.SizeXs, tempest.Class().Size(3)).
			If(props.Class != nil, props.Class),
		icons[props.Icon],
	)
}

func CreateIcon(nodes ...gox.Node) gox.Node {
	return gox.Svg(
		gox.ViewBox("0 0 24 24"),
		gox.Fragment(nodes...),
	)
}
