package radio_ui

import (
	"github.com/daarlabs/arcanum/gox"
	
	"component/ui/form_ui/error_message_ui"
	"component/ui/form_ui/field_label_ui"
)

func Radio(props Props) gox.Node {
	return gox.Div(
		gox.Class("flex flex-col gap-1"),
		gox.If(
			len(props.Label) > 0,
			field_label_ui.FieldLabel(
				field_label_ui.Props{
					For:      props.Id,
					Text:     props.Label,
					Required: props.Required,
				},
			),
		),
		gox.Div(
			gox.Class("grid gap-2"),
			gox.Range(
				props.Options,
				func(item Option, index int) gox.Node {
					return gox.Div(
						gox.Class("flex items-start gap-2"),
						gox.Input(
							gox.If(len(props.Id) > 0, gox.Id(props.Id+"-"+item.Value)),
							gox.Class("hidden-input peer"),
							gox.Type("radio"),
							gox.Name(props.Name),
							gox.Value(item.Value),
							gox.If(item.Checked, gox.Checked()),
							gox.If(props.Disabled, gox.Disabled()),
						),
						gox.Label(
							gox.For(props.Id+"-"+item.Value),
							gox.Clsx{
								"group relative flex-none transition border size-4 rounded-full block": true,
								"bg-transparent border-slate-300 dark:border-slate-600":                true,
								"peer-checked:border-primary-400 dark:peer-checked:border-primary-200": true,
							},
							gox.Div(
								gox.Class("transition absolute inset-0 m-auto size-1 rounded-full bg-transparent dark:bg-transparent peer-checked:group-[]:bg-primary-400 dark:peer-checked:group-[]:bg-primary-200"),
							),
						),
						gox.If(
							len(item.Title) > 0,
							field_label_ui.FieldLabel(
								field_label_ui.Props{
									For:  props.Id + "-" + item.Value,
									Text: item.Title,
								},
							),
						),
					)
				},
			),
		),
		gox.Range(
			props.Messages, func(msg string, _ int) gox.Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
