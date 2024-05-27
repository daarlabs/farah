package button_ui

import . "github.com/daarlabs/arcanum/gox"

func EmeraldButton(props Props, nodes ...Node) Node {
	return button(
		props,
		"bg-emerald-500 hover:bg-emerald-600 text-white",
		nodes...,
	)
}
