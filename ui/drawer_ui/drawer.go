package drawer_ui

import (
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Drawer(header Node, footer Node, nodes ...Node) Node {
	return Div(
		tempest.Class().Transition().Fixed().M("auto").Inset(0).H("screen").W("screen").
			BgSlate(900, tempest.Opacity(80)).Z(50),
		alpine.Cloak(),
		alpine.Show("open"),
		Div(
			alpine.Click("open = false", alpine.Outside),
			tempest.Class().Absolute().InsetY(0).My("auto").Right(0).W("300px").H("screen").
				BgWhite().BgSlate(900, tempest.Dark()).BorderL(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
				Grid().GridRows("48px_1fr_48px"),
			Div(
				tempest.Class().W("full").H("full").BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
					Px(4).Flex().ItemsCenter().Justify("between"),
				Div(
					tempest.Class().TextXs().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()),
					header,
				),
				Button(
					alpine.Click("open = false"),
					Type("button"),
					icon_ui.Icon(
						icon_ui.Props{
							Icon: icon_ui.Close, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
						},
					),
				),
			),
			Div(
				tempest.Class().OverflowX("hidden").OverflowY("auto").H("full").P(4),
				Fragment(nodes...),
			),
			Div(
				tempest.Class().W("full").H("full").BorderT(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).Px(2),
				footer,
			),
		),
	)
}
