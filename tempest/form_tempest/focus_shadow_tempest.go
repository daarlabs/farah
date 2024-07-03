package form_tempest

import (
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/hirokit/tempest"
)

func FocusShadow() tempest.Tempest {
	return tempest.Class().Shadow("focus", tempest.Focus()).
		ShadowColor(palette.Primary, 200, tempest.Dark(), tempest.Opacity(0.2))
}
