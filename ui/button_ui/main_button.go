package button_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func MainButton(props Props, nodes ...Node) Node {
	return button(
		props,
		tempest.Class().BgWhite().BgSlate(900, tempest.Dark()).
			BorderSlate(300).BorderSlate(600, tempest.Dark()).
			BorderSlate(400, tempest.Hover()).BorderSlate(200, tempest.Dark(), tempest.Hover()).
			TextSlate(900).TextWhite(tempest.Dark()).
			FontBold(),
		nodes...,
	)
}
