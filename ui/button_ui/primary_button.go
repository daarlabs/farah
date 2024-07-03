package button_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	"github.com/daarlabs/farah/palette"
)

func PrimaryButton(props Props, nodes ...Node) Node {
	return CreateButton(
		props,
		tempest.Class().Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()).
			Bg(palette.Primary, 500, tempest.Hover()).Bg(palette.Primary, 300, tempest.Dark(), tempest.Hover()).
			TextWhite(),
		nodes...,
	)
}
