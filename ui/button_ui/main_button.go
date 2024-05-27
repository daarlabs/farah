package button_ui

import . "github.com/daarlabs/arcanum/gox"

func MainButton(props Props, nodes ...Node) Node {
	return button(
		props,
		"bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 font-bold hover:border-primary-400 dark:hover:border-primary-200 text-slate-900 dark:text-white",
		nodes...,
	)
}
