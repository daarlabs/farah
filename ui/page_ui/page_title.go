package page_ui

import "github.com/daarlabs/arcanum/gox"

func Title(title string) gox.Node {
	return gox.H1(
		gox.Class("text-2xl font-bold text-slate-900 dark:text-white mb-6"),
		gox.Text(title),
	)
}
