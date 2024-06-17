package button_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func RedButton(props Props, nodes ...Node) Node {
	return CreateButton(
		props,
		tempest.Class().BgRed(400).BgRed(600, tempest.Hover()).TextWhite(),
		nodes...,
	)
}
