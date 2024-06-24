package tab_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func TabHeader(nodes ...Node) Node {
	return Div(
		tempest.Class().Flex().Gap(4).FlexWrap(),
		Fragment(nodes...),
	)
}
