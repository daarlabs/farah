package menu_ui

import (
	"github.com/daarlabs/arcanum/stimulus"
	
	"github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
)

func Menu(props Props, handler gox.Node, nodes ...gox.Node) gox.Node {
	if len(props.PositionX) == 0 {
		props.PositionX = ui.Left
	}
	if len(props.PositionY) == 0 {
		props.PositionY = ui.Bottom
	}
	return gox.Div(
		gox.If(
			props.Id != "",
			gox.Id(props.Id),
		),
		gox.Clsx{
			"relative flex": true,
			"group":         !props.Clickable,
		},
		gox.If(
			props.Clickable,
			stimulus.Controller("menu"),
		),
		handler,
		gox.Div(
			gox.If(
				props.OptionsId != "",
				gox.Id(props.OptionsId),
			),
			gox.Clsx{
				"transition absolute z-40 overflow-y-auto rounded bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 shadow-lg": true,
				"min-h-[200px]":          !props.Autoheight,
				"max-h-[200px]":          props.Scrollable,
				"is-invisible":           !props.Open,
				"is-visible":             props.Open,
				"top-full":               props.PositionY == ui.Bottom,
				"right-0":                props.PositionX == ui.Right,
				"bottom-full":            props.PositionY == ui.Top,
				"left-0":                 props.PositionX == ui.Left,
				"origin-top-left":        props.PositionY == ui.Bottom && props.PositionX == ui.Left,
				"origin-top-right":       props.PositionY == ui.Bottom && props.PositionX == ui.Right,
				"origin-bottom-left":     props.PositionY == ui.Top && props.PositionX == ui.Left,
				"origin-bottom-right":    props.PositionY == ui.Top && props.PositionX == ui.Right,
				"group-hover:is-visible": !props.Clickable,
				"is-clickable":           props.Clickable,
				"w-full":                 props.Fullwidth,
				"w-[200px]":              !props.Fullwidth,
			},
			gox.If(
				props.Clickable,
				stimulus.Target("menu", "menu"),
			),
			gox.Fragment(nodes...),
		),
	)
}
