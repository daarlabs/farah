package box_ui

import . "github.com/daarlabs/arcanum/gox"

func Box(props Props, nodes ...Node) Node {
	return Div(
		Clsx{
			"transition w-full h-full border rounded shadow overflow-hidden": true,
			"border-slate-300 bg-white":                                      true,
			props.Class:                                                      len(props.Class) > 0,
		},
		Fragment(nodes...),
	)
}
