package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"component/ui"
	"component/ui/icon_ui"
)

func NavMenuHandler(nodes ...Node) Node {
	return Div(
		Class("flex items-center gap-1"),
		Fragment(nodes...),
		Div(
			Class("transition group-hover:-rotate-180"),
			icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.ChevronDown, Size: ui.Sm,
					Class: "text-slate-900 dark:text-white group-hover:text-primary-400 dark:group-hover:text-primary-100",
				},
			),
		),
	)
}
