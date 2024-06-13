package nav_section_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	
	"github.com/daarlabs/farah/ui"
)

func NavSection(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Flex().ItemsCenter().
			If(props.AlignX == ui.Right, tempest.Class().Ml("auto")).
			If(props.AlignX == ui.Left, tempest.Class().Mr("auto")).
			If(props.AlignX == ui.Center, tempest.Class().Mx("auto")),
		Fragment(nodes...),
	)
}
