package filepicker_ui

import (
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
)

func Filepicker(c hiro.Ctx, props Props) Node {
	return Div(
		alpine.Data(hiro.Map{"files": []string{}, "chosen": false, "dragover": false}),
		tempest.Class().Flex().FlexCol().Gap(1),
		If(
			len(props.Label) > 0,
			field_label_ui.FieldLabel(
				field_label_ui.Props{
					For:      props.Id,
					Text:     props.Label,
					Required: props.Required,
				},
			),
		),
		Div(
			alpine.Click("$refs.input.click()"),
			alpine.DragOver("dragover = true", alpine.Prevent),
			alpine.DragLeave("dragover = false", alpine.Prevent),
			alpine.Drop(
				"dragover = false; $refs.input.files = $event.dataTransfer.files;files = Array.from($refs.input.files).map(f => f.name); chosen = $refs.input.files.length > 0",
				alpine.Prevent,
			),
			alpine.Class(
				map[string]string{
					tempest.Class().BorderSlate(300).
						BorderSlate(600, tempest.Dark()).
						String(): "!dragover",
					tempest.Class().BorderColor(palette.Primary, 400).
						BorderColor(palette.Primary, 200, tempest.Dark()).
						String(): "dragover",
				},
			),
			tempest.Class().Transition().Rounded().P(4).CursorPointer().
				TextXs().TextSlate(900).TextWhite(tempest.Dark()).
				Bg(tempest.White, 0).Bg(tempest.Slate, 800, tempest.Dark()).
				Border(1).
				BorderSlate(300).BorderSlate(600, tempest.Dark()).
				BorderColor(palette.Primary, 400, tempest.Focus()).BorderColor(
				palette.Primary, 200, tempest.Dark(), tempest.Focus(),
			).
				Extend(form_tempest.FocusShadow()),
			TabIndex(0),
			Role("button"),
			Input(
				alpine.Ref("input"),
				alpine.Change("files = Array.from($el.files).map(f => f.name); chosen = $el.files.length > 0"),
				tempest.Class().Hidden(),
				Id(props.Id),
				Name(props.Name),
				Type("file"),
				If(props.Disabled, Disabled()),
				If(props.Multiple, Multiple()),
			),
			Div(
				alpine.Show("!chosen"),
				tempest.Class().TextSlate(500).TextSlate(400, tempest.Dark()).Flex().ItemsStart().Gap(2),
				icon_ui.Icon(
					icon_ui.Props{
						Icon: icon_ui.File, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
					},
				),
				Text(c.Translate("component.filepicker.choose.file")),
			),
			Div(
				tempest.Class().Grid().Gap(2),
				Template(
					alpine.For("files", "file"),
					Div(
						alpine.Text("file"),
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
