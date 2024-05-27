package text_field_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"component/ui/form_ui/error_message_ui"
	"component/ui/form_ui/field_label_ui"
)

func TextField(props Props) Node {
	return Div(
		Class("flex flex-col gap-1"),
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
			Clsx{
				"transition w-full border h-10 px-3 rounded text-xs focus:shadow-focus":                            true,
				"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600": true,
				"focus:border-primary-400 dark:focus:border-primary-200":                                           true,
				"is-disabled": props.Disabled,
			},
			Type(TypeText),
			Name(props.Name),
			Value(props.Value),
			If(props.Disabled, Disabled()),
			If(props.Autofocus, AutoFocus()),
		),
		Range(
			props.Messages, func(msg string, _ int) Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
