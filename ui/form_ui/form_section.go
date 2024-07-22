package form_ui

import (
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Section(title string, nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).Rounded(),
		gox.If(
			len(title) > 0,
			gox.Div(
				tempest.Class().TextSlate(900).TextWhite(tempest.Dark()).FontBold().Px(4).Pt(4),
				gox.Text(title),
			),
		),
		gox.Div(
			tempest.Class().P(4),
			gox.Fragment(nodes...),
		),
	)
}
