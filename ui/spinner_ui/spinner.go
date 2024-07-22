package spinner_ui

import (
	"github.com/daarlabs/farah/palette"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

const (
	FormIndicator = "form-indicator"
	LinkIndicator = "link-indicator"
	HxIndicator   = "htmx-indicator"
)

func Spinner(props Props, nodes ...Node) Node {
	return Div(
		If(
			props.Overlay,
			tempest.Class().Transition().Absolute().Inset(0).M("auto").Grid().PlaceItemsCenter().
				BgWhite(tempest.Opacity(0.7)).BgSlate(800, tempest.Dark(), tempest.Opacity(0.7)).
				If(props.Class != nil, props.Class),
		),
		spinner(),
		// Div(
		// 	tempest.Class().Spin().Size(5).RoundedFull().Border(3).BorderWhite().BorderSlate(600, tempest.Dark()).
		// 		BorderTColor(palette.Primary, 400).BorderTColor(palette.Primary, 100, tempest.Dark()),
		// ),
		Fragment(nodes...),
	)
}

func spinner() Node {
	return Div(
		tempest.Class().Spin().Size(5).Grid().PlaceItemsCenter(),
		Svg(
			tempest.Class().FillCurrent().Text(palette.Primary, 400).TextWhite(tempest.Dark()).W(5).Mt(1),
			Xmlns("http://www.w3.org/2000/svg"),
			ViewBox("0 0 320 320"),
			Path(
				FillRule("evenodd"),
				ClipRule("evenodd"),
				D("M133.043 284C145.359 305.333 176.151 305.333 188.468 284L313.176 68C325.493 46.6667 310.097 20 285.463 20L36.0477 20C11.4141 20 -3.98188 46.6667 8.33493 68L133.043 284ZM78.483 76.5L160.755 219L243.028 76.5L78.483 76.5Z"),
			),
		),
	)
}
