package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
)

func Page(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		gox.Clsx{
			"grid gap-4 h-full grid-rows-[48px_1fr]": true,
		},
		gox.Fragment(nodes...),
	)
}
