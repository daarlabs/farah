package button_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func EmeraldButton(props Props, nodes ...Node) Node {
	return CreateButton(
		props,
		tempest.Class().BgEmerald(500).BgEmerald(600, tempest.Hover()).TextWhite(),
		nodes...,
	)
}
