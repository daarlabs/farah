package checkbox_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Checkbox(props Props, nodes ...Node) Node {
	return Div(
		Class("flex flex-col gap-1"),
		Label(
			If(props.Id != "", For(props.Id)),
			Class("inline-flex items-start gap-2 cursor-pointer relative"),
			Div(
				Class("relative w-5 h-5"),
				Input(
					If(props.Id != "", Id(props.Id)),
					If(props.Name != "", Name(props.Name)),
					Type("checkbox"),
					Class("peer w-0 h-0 absolute"),
					If(props.Value != nil, Value(props.Value)),
					If(props.Checked, Checked()),
					Fragment(nodes...),
				),
				Div(
					Clsx{
						"transition grid place-items-center w-5 h-5 rounded border":                                        true,
						"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600": true,
						"peer-checked:bg-primary-400 peer-checked:border-primary-400 peer-focus:shadow-focus":              true,
						"dark:peer-checked:bg-primary-200 dark:peer-checked:border-primary-200":                            true,
					},
				),
				icon_ui.Icon(
					icon_ui.Props{
						Icon:  icon_ui.Check,
						Size:  ui.Sm,
						Class: "text-transparent peer-checked:text-white absolute inset-0 m-auto w-4 h-4 z-10",
					},
				),
			),
			If(
				len(props.Label) > 0,
				Div(
					Class("flex min-h-[20px] pt-0.5"),
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
