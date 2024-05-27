package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
)

func NavBar(nodes ...Node) Node {
	return Div(
		Class("transition flex items-center px-6 gap-6 w-full h-16 bg-white dark:bg-slate-800 border-b border-slate-300 dark:border-slate-600"),
		Fragment(nodes...),
	)
}
