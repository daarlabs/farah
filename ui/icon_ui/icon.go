package icon_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	
	"github.com/daarlabs/farah/ui"
)

func Icon(props Props) Node {
	if len(props.Size) == 0 {
		props.Size = ui.Main
	}
	return Svg(
		tempest.Class().Transition().FillCurrent().
			If(props.Size == tempest.SizeMain, tempest.Class().Size(6)).
			If(props.Size == tempest.SizeSm, tempest.Class().Size(4)).
			If(props.Size == tempest.SizeXs, tempest.Class().Size(3)).
			If(props.Class != nil, props.Class),
		ViewBox("0 0 24 24"),
		icons[props.Icon],
	)
}
