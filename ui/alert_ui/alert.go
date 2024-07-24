package alert_ui

import (
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
)

func Alert(props Props, nodes ...Node) Node {
	return Div(
		alpine.Data(hiro.Map{}),
		tempest.Class().Relative().W("full").P(4).TextCenter().TextXs().Border(1).Rounded().TextWhite(tempest.Dark()).
			If(
				props.Type == AlertInfo,
				tempest.Class().
					BgSky(100).BorderSky(400).TextSky(700).
					BgSky(700, tempest.Dark()).BorderSky(400, tempest.Dark()),
			).
			If(
				props.Type == AlertSuccess,
				tempest.Class().
					BgEmerald(100).BorderEmerald(400).TextEmerald(700).
					BgEmerald(700, tempest.Dark()).BorderEmerald(400, tempest.Dark()),
			).
			If(
				props.Type == AlertWarning,
				tempest.Class().
					BgAmber(100).BorderAmber(400).TextAmber(700).
					BgAmber(700, tempest.Dark()).BorderAmber(400, tempest.Dark()),
			).
			If(
				props.Type == AlertError,
				tempest.Class().
					BgRed(100).BorderRed(400).TextRed(700).
					BgRed(700, tempest.Dark()).BorderRed(400, tempest.Dark()),
			).
			If(props.Class != nil, props.Class),
		Button(
			tempest.Class().Absolute().Top(1).Right(1),
			Type("button"),
			alpine.Click("$el.parentNode.parentNode.removeChild($el.parentNode)"),
			icon_ui.Icon(
				icon_ui.Props{
					Icon:  icon_ui.Close,
					Size:  ui.Sm,
					Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
				},
			),
		),
		Fragment(nodes...),
	)
}
