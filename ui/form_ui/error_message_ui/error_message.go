package error_message_ui

import . "github.com/daarlabs/arcanum/gox"

func ErrorMessage(msg string) Node {
	return Div(
		Class("w-full px-2 py-1 border border-red-400 dark:border-red-600 bg-red-50 dark:bg-red-500 text-red-600 dark:text-white text-[10px] rounded"),
		Text(msg),
	)
}
