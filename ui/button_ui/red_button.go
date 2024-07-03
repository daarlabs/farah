package button_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func RedButton(props Props, nodes ...Node) Node {
	return CreateButton(
		props,
		tempest.Class().BgRed(400).BgRed(600, tempest.Hover()).TextWhite(),
		nodes...,
	)
}
