package breadcrumbs_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Breadcrumbs(mainLink string, nodes ...Node) Node {
	return Div(
		Class("flex items-center gap-2"),
		A(
			Href(mainLink), icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.Home, Size: ui.Sm,
					Class: "text-slate-900 dark:text-white hover:text-primary-400 dark:hover:text-primary-100",
				},
			),
		),
		Fragment(nodes...),
	)
}
