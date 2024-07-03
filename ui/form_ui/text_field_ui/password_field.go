package text_field_ui

import (
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/form_tempest/form_input_tempest"
	"github.com/daarlabs/farah/tempest/util_tempest"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func PasswordField(props Props) Node {
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
			alpine.Data(map[string]any{"type": "password"}),
			tempest.Class().Relative(),
			Input(
				If(len(props.Id) > 0, Id(props.Id)),
				tempest.Class().
					H(10).
					Extend(form_input_tempest.InputField(form_input_tempest.Props{Boxed: props.Boxed})).
					Extend(form_tempest.FocusShadow()).
					If(props.Disabled, util_tempest.Disabled()),
				alpine.Bind("type"),
				Type(TypePassword),
				Name(props.Name),
				Value(props.Value),
				If(props.Disabled, Disabled()),
				If(props.Autofocus, AutoFocus()),
			),
			Button(
				tempest.Class().Absolute().Right(2).InsetY(0).My("auto").Size(4),
				alpine.Click("type = type === 'password' ? 'text' : 'password'"),
				Type("button"),
				Div(
					alpine.Show("type === 'password'"),
					icon_ui.Icon(
						icon_ui.Props{
							Icon: icon_ui.EyeOff, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
						},
					),
				),
				Div(
					alpine.Show("type === 'text'"),
					alpine.Cloak(),
					icon_ui.Icon(
						icon_ui.Props{
							Icon: icon_ui.Eye, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
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
