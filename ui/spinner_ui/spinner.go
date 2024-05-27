package spinner_ui

import . "github.com/daarlabs/arcanum/gox"

func Spinner(props Props) Node {
	return Div(
		Clsx{
			"transition absolute inset-0 m-auto grid place-items-center bg-white dark:bg-slate-800 bg-opacity-70 dark:bg-opacity-70": true,
			"indicator": !props.Visible,
		},
		Div(Class("animate-spin size-5 rounded-full border-[3px] border-white dark:border-slate-700 border-t-primary-400 dark:border-t-primary-100")),
	)
}
