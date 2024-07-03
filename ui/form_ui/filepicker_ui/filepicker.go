package filepicker_ui

import (
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Filepicker(props Props) Node {
	return Div(
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
			tempest.Class().Transition().Rounded().P(4).
				TextXs().TextSlate(900).TextWhite(tempest.Dark()).
				Bg(tempest.White, 0).Bg(tempest.Slate, 800, tempest.Dark()).
				Border(1).
				BorderSlate(300).BorderSlate(600, tempest.Dark()).
				BorderSlate(400, tempest.Focus()).BorderSlate(100, tempest.Dark(), tempest.Focus()).
				Extend(form_tempest.FocusShadow()),
			TabIndex(0),
			Role("button"),
			Input(
				tempest.Class().Block(),
				Id(props.Id),
				Name(props.Name),
				Type("file"),
				Value(props.Value),
				If(props.Disabled, Disabled()),
			),
		),
		Range(
			props.Messages, func(msg string, _ int) Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
