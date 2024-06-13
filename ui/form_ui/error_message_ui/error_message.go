package error_message_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func ErrorMessage(msg string, nodes ...Node) Node {
	return Div(
		tempest.Class().
			W("full").
			Px(2).
			Py(1).
			Rounded().
			Border(1).
			BorderColor(tempest.Red, 400).
			BorderColor(tempest.Red, 600, tempest.Dark()).
			Bg(tempest.Red, 50).
			Bg(tempest.Red, 500, tempest.Dark()).
			Text(tempest.Red, 600).
			TextWhite(tempest.Dark()).
			TextSize("10px"),
		Text(msg),
		Fragment(nodes...),
	)
}
