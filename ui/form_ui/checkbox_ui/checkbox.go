package checkbox_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Checkbox(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Flex().FlexCol().Gap(1),
		Label(
			If(props.Id != "", For(props.Id)),
			tempest.Class().Relative().InlineFlex().ItemsStart().Gap(2).Cursor("pointer"),
			Div(
				tempest.Class().Relative().Size(5),
				Input(
					tempest.Class().Peer().Size(0).Absolute().Opacity(0).Invisible(),
					If(props.Id != "", Id(props.Id)),
					If(props.Name != "", Name(props.Name)),
					Type("checkbox"),
					If(props.Value != nil, Value(props.Value)),
					If(props.Checked, Checked()),
					Fragment(nodes...),
				),
				Div(
					tempest.Class().Transition().Grid().PlaceItems("center").Size(5).Rounded().
						TextSlate(900).TextWhite(tempest.Dark()).
						Shadow("focus", tempest.Checked(tempest.Peer)).
						// Bg
						BgWhite().BgSlate(800, tempest.Dark()).
						Bg(palette.Primary, 400, tempest.Checked(tempest.Peer)).
						Bg(palette.Primary, 200, tempest.Dark(), tempest.Checked(tempest.Peer)).
						// Border
						Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
						BorderColor(palette.Primary, 400, tempest.Checked(tempest.Peer)).
						BorderColor(palette.Primary, 200, tempest.Dark(), tempest.Checked(tempest.Peer)),
				),
				icon_ui.Icon(
					icon_ui.Props{
						Icon: icon_ui.Check,
						Size: ui.Sm,
						Class: tempest.Class().TextTransparent().Size(4).Absolute().Inset(0).Z(10).M("auto").
							TextWhite(tempest.Checked(tempest.Peer)),
					},
				),
			),
			If(
				len(props.Label) > 0,
				Div(
					tempest.Class().Flex().Mt(1),
					field_label_ui.FieldLabel(
						field_label_ui.Props{
							For:      props.Id,
							Text:     props.Label,
							Required: props.Required,
						},
					),
				),
			),
		),
		Range(
			props.Messages, func(msg string, _ int) Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
