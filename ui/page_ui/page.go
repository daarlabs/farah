package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Page(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Grid().H("full").
			If(props.Header && !props.Footer, tempest.Class().GridRows("48px 1fr")).
			If(!props.Header && props.Footer, tempest.Class().GridRows("1fr 64px")).
			If(props.Header && props.Footer, tempest.Class().GridRows("48px 1fr 64px")),
		gox.Fragment(nodes...),
	)
}
