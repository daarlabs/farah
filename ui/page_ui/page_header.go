package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Header(nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Flex().ItemsCenter().Justify("between").H(12).Px(6),
		gox.Fragment(nodes...),
	)
}
