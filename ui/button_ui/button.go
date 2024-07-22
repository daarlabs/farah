package button_ui

import (
	"github.com/daarlabs/farah/ui/spinner_ui"
	"github.com/daarlabs/hirokit/alpine"
	"github.com/daarlabs/hirokit/tempest"
	
	. "github.com/daarlabs/hirokit/gox"
	
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
			spinner_ui.Spinner(
				spinner_ui.Props{
					Overlay: true,
					Class:   tempest.Class(spinner_ui.FormIndicator),
				},
			),
		),
		If(
			isLink,
			alpine.Data(map[string]any{"pending": false}),
			alpine.Class(map[string]string{"link-request": "pending"}),
			alpine.Click("pending = true"),
			alpine.KeyUp("pending = false", alpine.Escape),
			spinner_ui.Spinner(
				spinner_ui.Props{
					Overlay: true,
					Class:   tempest.Class(spinner_ui.LinkIndicator),
				},
			),
		),
		Fragment(nodes...),
		If(len(props.Icon) > 0, icon_ui.Icon(icon_ui.Props{Icon: props.Icon, Size: ui.Sm})),
	)
}
