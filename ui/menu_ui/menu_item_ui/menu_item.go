package menu_item_ui

import (
	. "github.com/daarlabs/arcanum/gox"
)

func MenuItem(props Props, nodes ...Node) Node {
	return Div(
		Clsx{
			"transition border-b border-slate-200 dark:border-slate-700 w-full px-4 py-2 text-xs":                  true,
			"text-slate-800 dark:text-white bg-white dark:bg-slate-800 hover:bg-slate-100 dark:hover:bg-slate-700": !props.Selected,
			"text-white bg-primary-400 dark:bg-primary-200":                                                        props.Selected,
		},
		Fragment(nodes...),
	)
}
