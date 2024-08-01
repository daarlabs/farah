package tab_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func TabHeader(nodes ...Node) Node {
	return Div(
		tempest.Class().Flex().Gap(2).FlexWrap(),
		Fragment(nodes...),
	)
}
