package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
)

func NavLink(link string, nodes ...Node) Node {
	return A(
		Class("transition text-xs font-semibold text-slate-900 dark:text-white hover:text-primary-400 dark:hover:text-primary-100"),
		Href(link),
		Fragment(nodes...),
	)
}
