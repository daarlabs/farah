package box_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Box(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Transition().Rounded().ShadowMain().Overflow("hidden").
			Grid().
			If(len(props.Title) > 0, tempest.Class().GridRows("48px 1fr")).
			If(len(props.Title) == 0, tempest.Class().GridRows(1)).
			BgWhite().BgSlate(800, tempest.Dark()).
			Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
			If(props.Class != nil, props.Class),
		Fragment(nodes...),
	)
}
