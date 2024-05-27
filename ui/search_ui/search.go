package search_ui

import (
	"github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Search(props Props, nodes ...gox.Node) gox.Node {
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
			gox.Class("relative"),
			gox.Div(
				gox.Class("absolute size-4 inset-y-0 my-auto left-3"),
				icon_ui.Icon(icon_ui.Props{Icon: icon_ui.SearchFilter, Class: "text-slate-900 dark:text-white", Size: ui.Sm}),
			),
			gox.Input(
				gox.If(len(props.Id) > 0, gox.Id(props.Id)),
				gox.Clsx{
					"transition w-full border h-10 pl-10 pr-3 rounded text-xs focus:shadow-focus":                      true,
					"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600": true,
					"placeholder:text-slate-600 dark:placeholder:text-slate-300":                                       true,
					"focus:border-primary-400 dark:focus:border-primary-200":                                           true,
					"is-disabled": props.Disabled,
				},
				gox.Type("text"),
				gox.Name(props.Name),
				gox.Value(props.Value),
				gox.If(props.Placeholder != "", gox.Placeholder(props.Placeholder)),
				gox.If(props.Disabled, gox.Disabled()),
				gox.If(props.Autofocus, gox.AutoFocus()),
				gox.Fragment(nodes...),
			),
		),
		gox.Range(
			props.Messages, func(msg string, _ int) gox.Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
