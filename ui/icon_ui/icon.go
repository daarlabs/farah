package icon_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
)

func Icon(props Props) Node {
	if len(props.Size) == 0 {
		props.Size = ui.Main
	}
	return Svg(
		Clsx{
			"transition fill-current": true,
			"w-6 h-6":                 props.Size == "main",
			"w-4 h-4":                 props.Size == "sm",
			"w-3 h-3":                 props.Size == "xs",
			props.Class:               len(props.Class) > 0,
		},
		ViewBox("0 0 24 24"),
		icons[props.Icon],
	)
}
