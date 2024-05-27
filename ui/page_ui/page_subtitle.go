package page_ui

import "github.com/daarlabs/arcanum/gox"

func Subtitle(title string) gox.Node {
	return gox.H1(
		gox.Class("text-lg font-semibold text-slate-900 dark:text-white mb-4"),
		gox.Text(title),
	)
}
