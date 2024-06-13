package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func NavBar(nodes ...Node) Node {
	return Div(
		tempest.Class().Transition().Flex().ItemsCenter().Px(6).Gap(6).W("full").H(16).
			BgWhite().BgSlate(800, tempest.Dark()).
			BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
		Fragment(nodes...),
	)
}
