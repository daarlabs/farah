package anchor_ui

import (
	"github.com/daarlabs/farah/ui/spinner_ui"
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
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
				Class:   tempest.Class(spinner_ui.LinkIndicator),
			},
		),
	)
}
