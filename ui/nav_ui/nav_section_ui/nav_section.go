package nav_section_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"component/ui"
)

func NavSection(props Props, nodes ...Node) Node {
	return Div(
		Clsx{
			"flex items-center": true,
			"ml-auto":           props.AlignX == ui.Right,
			"mr-auto":           props.AlignX == ui.Left,
			"mx-auto":           props.AlignX == ui.Center,
		},
		Fragment(nodes...),
	)
}
