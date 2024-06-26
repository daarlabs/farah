package form_input_tempest

import (
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
)

func InputField(props Props) tempest.Tempest {
	textExists := len(props.Text) > 0
	placeholderExists := len(props.Placeholder) > 0
	return tempest.Class().Transition().W("full").Px(3).Rounded().
		// Font
		TextSize(tempest.SizeXs).
		If(textExists || !placeholderExists, tempest.Class().TextSlate(900).TextWhite(tempest.Dark())).
		If(!textExists && !placeholderExists, tempest.Class().TextSlate(900).TextWhite(tempest.Dark())).
		If(
			!textExists && placeholderExists,
			tempest.Class().TextSlate(600).TextSlate(400, tempest.Dark()),
		).
		// Border
		Border(1).
		BorderColor(tempest.Slate, 300).
		BorderColor(tempest.Slate, 600, tempest.Dark()).
		BorderColor(palette.Primary, 400, tempest.Focus()).
		BorderColor(
			palette.Primary, 200, tempest.Focus(), tempest.Dark(),
		).
		If(!props.Boxed, tempest.Class().Bg(tempest.White, 0).Bg(tempest.Slate, 800, tempest.Dark())).
		If(
			props.Boxed,
			tempest.Class().Bg(tempest.Slate, 100).Bg(tempest.Slate, 700, tempest.Dark()),
		)
}
