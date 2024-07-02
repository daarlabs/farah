package anchor_ui

import (
	"github.com/daarlabs/arcanum/alpine"
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/ui/spinner_ui"
)

func Anchor(nodes ...Node) Node {
	return A(
		tempest.Class().Block().Relative(),
		alpine.Data(map[string]any{"pending": false}),
		alpine.Class(map[string]string{"link-request": "pending"}),
		alpine.Click("pending = true"),
		alpine.KeyUp("pending = false", alpine.Escape),
		Fragment(nodes...),
		spinner_ui.Spinner(
			spinner_ui.Props{
				Overlay: true,
				Class:   tempest.Class(spinner_ui.Indicator),
			},
		),
	)
}
