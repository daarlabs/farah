package button_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func EmeraldButton(props Props, nodes ...Node) Node {
	return button(
		props,
		tempest.Class().BgEmerald(500).BgEmerald(600, tempest.Hover()).TextWhite(),
		nodes...,
	)
}
