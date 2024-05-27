package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
)

func HeaderSection(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		gox.Clsx{
			"flex items-center h-12": true,
			"ml-auto":                props.AlignX == ui.Right,
			"mr-auto":                props.AlignX == ui.Left,
			"mx-auto":                props.AlignX == ui.Center,
		},
		gox.Fragment(nodes...),
	)
}
