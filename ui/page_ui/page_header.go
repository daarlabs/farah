package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
)

func Header(nodes ...gox.Node) gox.Node {
	return gox.Div(
		gox.Clsx{
			"flex items-center h-12 px-6": true,
		},
		gox.Fragment(nodes...),
	)
}
