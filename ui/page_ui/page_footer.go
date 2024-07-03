package page_ui

import (
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Footer(nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Flex().ItemsCenter().Justify("end").H(16).Px(6).
			BorderT(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
		gox.Fragment(nodes...),
	)
}
