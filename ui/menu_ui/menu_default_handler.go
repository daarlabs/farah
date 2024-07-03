package menu_ui

import (
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func DefaultHandler(props Props) gox.Node {
	size := 8
	return gox.Div(
		tempest.Class().Relative().W("full").H(size),
		Open(),
		gox.Button(
			gox.If(len(props.Id) > 0, gox.Id(props.Id)),
			gox.Type("button"),
			tempest.Class().Transition().W("full").H(size).Pl(3).Pr(7).Rounded().
				TextSize("10px").TextLeft().
				BgWhite().BgSlate(800, tempest.Dark()).
				Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
				BorderColor(palette.Primary, 400, tempest.Focus()).
				BorderColor(palette.Primary, 200, tempest.Focus(), tempest.Dark()).
				Extend(form_tempest.FocusShadow()).
				If(len(props.Text) > 0, tempest.Class().TextSlate(900).TextWhite(tempest.Dark())).
				If(
					len(props.Text) == 0 && len(props.Placeholder) > 0,
					tempest.Class().TextSlate(600).TextSlate(400, tempest.Dark()),
				),
			gox.If(len(props.Text) == 0 && len(props.Placeholder) > 0, gox.Text(props.Placeholder)),
			gox.If(len(props.Text) > 0, gox.Text(props.Text)),
		),
		gox.Label(
			tempest.Class().Transition().Absolute().Right(2).InsetY(0).My("auto").Size(4),
			gox.If(len(props.Id) > 0, gox.For(props.Id)),
			Chevron(),
			icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.ChevronDown, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
				},
			),
		),
	)
}
