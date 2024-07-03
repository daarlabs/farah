package nav_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	"github.com/daarlabs/farah/palette"
)

func NavLink(props Props, nodes ...Node) Node {
	return A(
		tempest.Class().
			Transition().
			TextXs().
			FontSemibold().TextSlate(900).TextWhite(tempest.Dark()).
			Text(palette.Primary, 400, tempest.Hover()).
			Text(palette.Primary, 100, tempest.Dark(), tempest.Hover()),
		Href(props.Link),
		Fragment(nodes...),
	)
}
