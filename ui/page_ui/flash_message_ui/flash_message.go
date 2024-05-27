package flash_message_ui

import (
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/stimulus"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"component/ui"
	"component/ui/icon_ui"
)

func Message(fm mirage.Message) Node {
	return Div(
		Clsx{
			"transition relative flex flex-start gap-4 w-full border rounded pl-4 pr-8 py-2 shadow-xl": true,
			"bg-white dark:bg-slate-800": true,
			"border-emerald-400 dark:border-emerald-300 text-emerald-600 dark:text-emerald-500": fm.Type == mirage.FlashSuccess,
			"border-amber-400 text-amber-600 dark:border-amber-300 dark:text-amber-500":         fm.Type == mirage.FlashWarning,
			"border-red-400 text-red-600 dark:border-red-400 dark:text-red-500":                 fm.Type == mirage.FlashError,
		},
		stimulus.Controller("flash-message"),
		Div(
			Class("flex flex-col gap-1"),
			Div(
				Class("font-bold text-xs"),
				Text(fm.Title),
			),
			If(
				len(fm.Value) > 0,
				Div(
					Class("transition text-slate-600 dark:text-slate-300 text-[10px]"),
					Text(fm.Value),
				),
			),
		),
		Button(
			Type("button"),
			Class("absolute top-2 right-2"),
			stimulus.Action("click", "flash-message", "handleClose"),
			icon_ui.Icon(icon_ui.Props{Icon: icon_ui.Close, Size: ui.Sm, Class: "text-slate-900 dark:text-white"}),
		),
	)
}
