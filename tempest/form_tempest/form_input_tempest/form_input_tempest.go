package form_input_tempest

import (
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/ui/form_ui"
	"github.com/daarlabs/hirokit/tempest"
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
		If(
			len(props.Status) == 0,
			tempest.Class().BorderColor(tempest.Slate, 300).
				BorderColor(tempest.Slate, 600, tempest.Dark()),
		).
		If(
			props.Status == form_ui.StatusSuccess,
			tempest.Class().BorderColor(tempest.Emerald, 400).
				BorderColor(tempest.Emerald, 500, tempest.Dark()),
		).
		If(
			props.Status == form_ui.StatusError,
			tempest.Class().BorderColor(tempest.Red, 400).
				BorderColor(tempest.Red, 500, tempest.Dark()),
		).
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
