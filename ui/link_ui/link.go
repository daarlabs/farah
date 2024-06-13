package link_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
)

func Link(props Props, nodes ...gox.Node) gox.Node {
	return gox.A(
		tempest.Class().Transition().TextXs().
			Text(palette.Primary, 400).Text(palette.Primary, 100, tempest.Dark()).
			Underline().NoUnderline(tempest.Hover()).
			If(props.Class != nil, props.Class),
		gox.If(len(props.Url) > 0, gox.Href(props.Url)),
		gox.Fragment(nodes...),
		gox.If(len(props.Title) > 0, gox.Text(props.Title)),
	)
}
