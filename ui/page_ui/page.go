package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Page(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Grid().H("full").
			If(props.Header, tempest.Class().GridRows("48px 1fr")),
		gox.Fragment(nodes...),
	)
}
