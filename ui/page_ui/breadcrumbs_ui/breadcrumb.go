package breadcrumbs_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	"github.com/daarlabs/farah/palette"
)

func Breadcrumb(link, label string, last ...bool) Node {
	isLast := false
	if len(last) > 0 {
		isLast = last[0]
	}
	return Fragment(
		Div(
			tempest.Class().TextXs().TextSlate(600),
			Text("/"),
		),
		A(
			tempest.Class().Transition().TextSize("10px").
				If(!isLast, tempest.Class().Underline().NoUnderline(tempest.Hover()).TextSlate(900).TextWhite(tempest.Dark())).
				If(isLast, tempest.Class().Text(palette.Primary, 400).Text(palette.Primary, 100, tempest.Dark())),
			Href(link),
			Text(label),
		),
	)
}
