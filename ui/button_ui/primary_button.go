package button_ui

import . "github.com/daarlabs/arcanum/gox"

func PrimaryButton(props Props, nodes ...Node) Node {
	return button(
		props,
		"bg-primary-400 dark:bg-primary-200 hover:bg-primary-500 dark:hover:bg-primary-300 text-white",
		nodes...,
	)
}
