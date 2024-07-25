package spinner_ui

import (
	"github.com/daarlabs/farah/farah"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

const (
	FormIndicator = "form-indicator"
	LinkIndicator = "link-indicator"
	HxIndicator   = "htmx-indicator"
)

func Spinner(props Props, nodes ...Node) Node {
	return Div(
		If(
			props.Overlay,
			tempest.Class().Transition().Absolute().Inset(0).M("auto").Grid().PlaceItemsCenter().
				BgWhite(tempest.Opacity(0.7)).BgSlate(800, tempest.Dark(), tempest.Opacity(0.7)).
				If(props.Class != nil, props.Class),
		),
		createSpinner(farah.Config.Spinner),
		Fragment(nodes...),
	)
}

func createSpinner(spinner Node) Node {
	return Div(
		tempest.Class().Spin().Size(5).Grid().PlaceItemsCenter(),
		spinner,
	)
}
