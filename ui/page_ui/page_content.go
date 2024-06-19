package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Content(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Px(6).Pb(6).Pt(1).H("full").If(props.Class != nil, props.Class),
		gox.Fragment(nodes...),
	)
}
