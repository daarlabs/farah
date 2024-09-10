package number_field_ui

import (
	"golang.org/x/exp/constraints"
	
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/form_tempest/form_input_tempest"
	"github.com/daarlabs/farah/tempest/util_tempest"
	"github.com/daarlabs/farah/ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
)

func NumberField[T constraints.Integer | constraints.Float](props Props[T], nodes ...Node) Node {
	if len(props.Size) == 0 {
		props.Size = ui.Main
	}
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
				If(props.Size == ui.Main, tempest.Class().H(10)).
				If(props.Size == ui.Sm, tempest.Class().H(8)).
				Extend(form_input_tempest.InputField(form_input_tempest.Props{Boxed: props.Boxed, Size: props.Size})).
				Extend(form_tempest.FocusShadow()).
				If(props.Disabled, util_tempest.Disabled()),
			Type("number"),
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
