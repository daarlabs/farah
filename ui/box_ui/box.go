package box_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Box(props Props, nodes ...Node) Node {
	titleExists := len(props.Title) > 0
	return Div(
		tempest.Class().Transition().Rounded().ShadowMain().Overflow("hidden").
			Grid().
			If(titleExists, tempest.Class().GridRows("48px 1fr")).
			If(!titleExists, tempest.Class().GridRows(1)).
			BgWhite().BgSlate(800, tempest.Dark()).
			Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
			If(props.Class != nil, props.Class),
		If(
			titleExists,
			Div(
				tempest.Class().P(4).TextSm().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()),
				Text(props.Title),
			),
		),
		Fragment(nodes...),
	)
}
