package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
)

func Content(nodes ...gox.Node) gox.Node {
	return gox.Div(
		gox.Clsx{
			"px-6 pb-6": true,
		},
		gox.Fragment(nodes...),
	)
}
