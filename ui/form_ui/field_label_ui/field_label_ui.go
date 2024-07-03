package field_label_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func FieldLabel(props Props) Node {
	return Label(
		If(len(props.For) > 0, For(props.For)),
		tempest.Class().Transition().Text(tempest.Slate, 600).Text(tempest.Slate, 300, tempest.Dark()).
			TextSize("10px").Cursor("pointer"),
		Text(props.Text),
		If(
			props.Required,
			Span(
				tempest.Class().Text(tempest.Red, 500).TextSize(tempest.SizeXs).Ml(0.5),
				Text("*"),
			),
		),
	)
}
