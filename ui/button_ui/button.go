package button_ui

import (
	"github.com/daarlabs/arcanum/stimulus"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/spinner_ui"
)

func button(props Props, variantClass string, nodes ...Node) Node {
	isLink := props.Link != ""
	el := "button"
	if isLink {
		el = "a"
	}
	if len(props.Size) == 0 {
		props.Size = ui.Main
	}
	if !isLink && len(props.Type) == 0 {
		props.Type = TypeButton
	}
	return CreateElement(el)(
		If(isLink, Href(props.Link)),
		Clsx{
			"transition relative inline-flex justify-between gap-3 items-center rounded text-left focus:shadow-focus font-semibold": true,
			"h-10 px-3 text-xs":    props.Size == ui.Main,
			"h-8 px-3 text-[10px]": props.Size == ui.Sm,
			variantClass:           true,
			props.Class:            len(props.Class) > 0,
		},
		If(!isLink, Type(props.Type)),
		If(
			props.Type == TypeSubmit,
			stimulus.Controller("form"),
			stimulus.Action("click", "form", "handleSubmit"),
		),
		spinner_ui.Spinner(spinner_ui.Props{}),
		Fragment(nodes...),
		If(len(props.Icon) > 0, icon_ui.Icon(icon_ui.Props{Icon: props.Icon, Size: ui.Sm})),
	)
}
