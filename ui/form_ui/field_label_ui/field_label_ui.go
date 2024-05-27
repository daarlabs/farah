package field_label_ui

import . "github.com/daarlabs/arcanum/gox"

func FieldLabel(props Props) Node {
	return Label(
		If(len(props.For) > 0, For(props.For)),
		Class("transition text-slate-600 dark:text-slate-300 text-[10px] cursor-pointer"),
		Text(props.Text),
		If(props.Required, Span(Class("text-red-500 text-xs ml-0.5"), Text("*"))),
	)
}
