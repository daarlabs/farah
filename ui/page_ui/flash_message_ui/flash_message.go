package flash_message_ui

import (
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/hirokit/alpine"
	"github.com/daarlabs/hirokit/tempest"
	"github.com/daarlabs/hirokit/hiro"
	
	. "github.com/daarlabs/hirokit/gox"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

const (
	autoremoveMessageDelay = "5000"
)

func Message(fm hiro.Message) Node {
	id := uniuri.New()
	return Div(
		alpine.Data(hiro.Map{}),
		alpine.Init(`setTimeout(() => !!$refs['`+id+`'] && $refs['`+id+`'].remove(), `+autoremoveMessageDelay+`)`),
		alpine.Ref(id),
		tempest.Class().Transition().Relative().Flex().Gap(4).W("full").
			BgWhite().BgSlate(800, tempest.Dark()).
			Border(1).Rounded().Px(4).Py(2).ShadowXl().
			If(
				fm.Type == hiro.FlashSuccess,
				tempest.Class().BorderEmerald(400).BorderEmerald(300, tempest.Dark()).
					TextEmerald(600).TextEmerald(500, tempest.Dark()),
			).
			If(
				fm.Type == hiro.FlashWarning,
				tempest.Class().BorderAmber(400).BorderAmber(300, tempest.Dark()).
					TextAmber(600).TextAmber(500, tempest.Dark()),
			).
			If(
				fm.Type == hiro.FlashError,
				tempest.Class().BorderRed(400).BorderRed(400, tempest.Dark()).
					TextRed(600).TextRed(500, tempest.Dark()),
			),
		Div(
			tempest.Class().Flex().FlexCol().Gap(1),
			Div(
				tempest.Class().FontBold().TextXs(),
				Text(fm.Title),
			),
			If(
				len(fm.Value) > 0,
				Div(
					tempest.Class().Transition().TextSlate(600).TextSlate(300, tempest.Dark()).
						TextSize("10px").LhRelax(),
					Text(fm.Value),
				),
			),
		),
		Button(
			alpine.Click(`$refs['`+id+`'].remove()`),
			tempest.Class().Absolute().Top(2).Right(2),
			Type("button"),
			icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.Close, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
				},
			),
		),
	)
}
