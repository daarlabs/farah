package menu_ui

import (
	"github.com/daarlabs/hirokit/alpine"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/hirokit/gox"
	
	"github.com/daarlabs/farah/ui"
)

func Menu(props Props, handler gox.Node, nodes ...gox.Node) gox.Node {
	if len(props.PositionX) == 0 {
		props.PositionX = ui.Left
	}
	if len(props.PositionY) == 0 {
		props.PositionY = ui.Bottom
	}
	return gox.Div(
		gox.If(
			props.Id != "",
			gox.Id(props.Id),
		),
		tempest.Class().Relative().Flex(),
		alpine.Data(hiro.Map{"open": props.Open}),
		gox.If(
			props.Clickable,
			alpine.Click("open = false", alpine.Outside),
		),
		gox.If(
			!props.Clickable,
			alpine.MouseEnter("open = true"),
			alpine.MouseLeave("open = false"),
		),
		handler,
		gox.Div(
			gox.If(
				props.OptionsId != "",
				gox.Id(props.OptionsId),
			),
			alpine.Class(
				map[string]string{
					tempest.Class().Invisible().Opacity(0).Scale(95).H(0).String(): "!open",
					tempest.Class().Visible().Opacity(1).Scale(100).String():       "open",
				},
			),
			tempest.Class().Transition().Absolute().Z(40).OverflowY("auto").Rounded().ShadowLg().
				BgWhite().BgSlate(800, tempest.Dark()).
				Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
				If(!props.Autoheight, tempest.Class().MinH("200px")).
				If(props.Scrollable, tempest.Class().MaxH("200px")).
				If(!props.Open, tempest.Class().Invisible().Opacity(0).Scale(95)).H(0).
				If(props.Open, tempest.Class().Visible().Opacity(1).Scale(100)).
				If(props.PositionY == ui.Bottom, tempest.Class().Top("full")).
				If(props.PositionX == ui.Right, tempest.Class().Right(0)).
				If(props.PositionY == ui.Top, tempest.Class().Bottom("full")).
				If(props.PositionX == ui.Left, tempest.Class().Left(0)).
				If(props.PositionY == ui.Bottom && props.PositionX == ui.Left, tempest.Class().Origin("top-left")).
				If(props.PositionY == ui.Bottom && props.PositionX == ui.Right, tempest.Class().Origin("top-right")).
				If(props.PositionY == ui.Top && props.PositionX == ui.Left, tempest.Class().Origin("bottom-left")).
				If(props.PositionY == ui.Top && props.PositionX == ui.Right, tempest.Class().Origin("bottom-right")).
				If(props.Fullwidth, tempest.Class().W("full")).
				If(!props.Fullwidth, tempest.Class().W("200px")),
			gox.Fragment(nodes...),
		),
	)
}
