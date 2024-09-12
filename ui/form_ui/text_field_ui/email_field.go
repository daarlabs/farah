package text_field_ui

import (
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/form_tempest/form_input_tempest"
	"github.com/daarlabs/farah/tempest/util_tempest"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func EmailField(props Props, nodes ...Node) Node {
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
		Input(
			If(len(props.Id) > 0, Id(props.Id)),
			tempest.Class().
				H(10).
				Extend(
					form_input_tempest.InputField(
						form_input_tempest.Props{
							Text:        props.Value,
							Placeholder: props.Placeholder,
							Boxed:       props.Boxed,
						},
					),
				).
				Extend(form_tempest.FocusShadow()).
				If(props.Disabled, util_tempest.Disabled()),
			Type(TypeEmail),
			Name(props.Name),
			Value(props.Value),
			If(props.Disabled, Disabled()),
			If(props.Autofocus, AutoFocus()),
			Fragment(nodes...),
		),
		Range(
			props.Messages, func(msg string, _ int) Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
