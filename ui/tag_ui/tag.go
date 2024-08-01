package tag_ui

import (
	"github.com/daarlabs/farah/palette"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Tag(props Props, nodes ...Node) Node {
	return Div(
		tempest.Class().Px(2).Py(1).Rounded().
			Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()).
			TextWhite().TextSize("10px").Flex().ItemsCenter().Gap(1),
		Fragment(nodes...),
	)
}
