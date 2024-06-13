package textarea_ui

import (
	"fmt"
	
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/form_tempest/form_input_tempest"
	"github.com/daarlabs/farah/tempest/util_tempest"
	
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
)

func TextArea(props Props) Node {
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
		Textarea(
			If(len(props.Id) > 0, Id(props.Id)),
			tempest.Class().MinH("200px").
				Extend(form_input_tempest.InputField(form_input_tempest.Props{Boxed: props.Boxed})).
				Extend(form_tempest.FocusShadow()).
				If(props.Disabled, util_tempest.Disabled()),
			Name(props.Name),
			If(props.Disabled, Disabled()),
			If(len(fmt.Sprintf("%v", props.Value)) > 0, Text(props.Value)),
		),
		Range(
			props.Messages, func(msg string, _ int) Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
