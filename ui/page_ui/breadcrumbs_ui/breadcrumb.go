package breadcrumbs_ui

import (
	. "github.com/daarlabs/arcanum/gox"
)

func Breadcrumb(link, label string, last ...bool) Node {
	isLast := false
	if len(last) > 0 {
		isLast = last[0]
	}
	return Fragment(
		Div(Class("text-xs text-slate-600"), Text("/")),
		A(
			Href(link),
			Text(label),
			Clsx{
				"transition text-[10px]": true,
				"underline hover:no-underline text-slate-900 dark:text-white": !isLast,
				"text-primary-400 dark:text-primary-100":                      isLast,
			},
		),
	)
}
