package search_ui

import (
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/util_tempest"
	"github.com/daarlabs/farah/ui/spinner_ui"
	"github.com/daarlabs/hirokit/alpine"
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Search(props Props, nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Flex().FlexCol().Gap(1),
		gox.If(
			len(props.Label) > 0,
			field_label_ui.FieldLabel(
				field_label_ui.Props{
					For:      props.Id,
					Text:     props.Label,
					Required: props.Required,
				},
			),
		),
		gox.Div(
			tempest.Class().Relative(),
			alpine.Data(map[string]any{"submit": false}),
			alpine.Init("submit = false"),
			gox.Div(
				tempest.Class().Absolute().Size(4).InsetY(0).My("auto").Left(3),
				icon_ui.Icon(
					icon_ui.Props{
						Icon: icon_ui.SearchFilter, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()), Size: ui.Sm,
					},
				),
			),
			gox.Input(
				gox.If(len(props.Id) > 0, gox.Id(props.Id)),
				tempest.Class().
					H(10).
					Pr(10).
					Transition().W("full").Px(10).Rounded().
					// Font
					TextSize(tempest.SizeXs).Text(tempest.Slate, 900).TextWhite(tempest.Dark()).
					TextSlate(400, tempest.Placeholder()).TextSlate(500, tempest.Placeholder(), tempest.Dark()).
					// Border
					Border(1).
					BorderColor(tempest.Slate, 300).
					BorderColor(tempest.Slate, 600, tempest.Dark()).
					BorderColor(palette.Primary, 400, tempest.Focus()).
					BorderColor(
						palette.Primary, 200, tempest.Focus(), tempest.Dark(),
					).
					Bg(tempest.White, 0).Bg(tempest.Slate, 800, tempest.Dark()).
					Extend(form_tempest.FocusShadow()).
					If(props.Disabled, util_tempest.Disabled()),
				gox.Type("text"),
				gox.Name(props.Name),
				gox.Value(props.Value),
				gox.If(props.Placeholder != "", gox.Placeholder(props.Placeholder)),
				gox.If(props.Disabled, gox.Disabled()),
				gox.If(props.Autofocus, gox.AutoFocus()),
				gox.AutoComplete("off"),
				gox.Fragment(nodes...),
				alpine.KeyDown("submit = true", alpine.Enter),
				alpine.KeyDown("$el.value = '';$dispatch('change')", alpine.Escape),
			),
			spinner_ui.Spinner(
				spinner_ui.Props{},
				tempest.Class().Absolute().H(5).Right(3).InsetY(0).My("auto"),
				alpine.Show("submit"),
				alpine.Cloak(),
			),
		),
		gox.Range(
			props.Messages, func(msg string, _ int) gox.Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
