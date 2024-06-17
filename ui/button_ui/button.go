package button_ui

import (
	"github.com/daarlabs/arcanum/alpine"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/ui/spinner_ui"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func CreateButton(props Props, variantClass tempest.Tempest, nodes ...Node) Node {
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
		tempest.Class().Transition().Relative().InlineFlex().Justify("between").Items("center").Gap(3).Rounded().
			TextLeft().FontSemibold().Shadow("focus", tempest.Focus()).Cursor("pointer").
			If(variantClass != nil, variantClass).
			If(props.Class != nil, props.Class).
			If(props.Size == tempest.SizeMain, tempest.Class().H(10).Px(3).TextXs()).
			If(props.Size == tempest.SizeSm, tempest.Class().H(8).Px(3).TextSize("10px")),
		If(!isLink, Type(props.Type)),
		If(
			props.Type == TypeSubmit,
		),
		spinner_ui.Spinner(
			spinner_ui.Props{},
			alpine.Cloak(),
		),
		Fragment(nodes...),
		If(len(props.Icon) > 0, icon_ui.Icon(icon_ui.Props{Icon: props.Icon, Size: ui.Sm})),
	)
}
