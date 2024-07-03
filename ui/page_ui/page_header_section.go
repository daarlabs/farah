package page_ui

import (
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui"
)

func HeaderSection(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Flex().ItemsCenter().H(12).
			If(props.AlignX == ui.Right, tempest.Class().Ml("auto")).
			If(props.AlignX == ui.Left, tempest.Class().Mr("auto")).
			If(props.AlignX == ui.Center, tempest.Class().Mx("auto")),
		gox.Fragment(nodes...),
	)
}
