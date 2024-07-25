package farah

import (
	"github.com/daarlabs/farah/palette"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

var (
	defaultSpinner = Svg(
		tempest.Class().FillCurrent().Text(palette.Primary, 400).TextWhite(tempest.Dark()).W(5).Mt(1),
		Xmlns("http://www.w3.org/2000/svg"),
		ViewBox("0 0 24 24"),
		Path(
			D(`M12,4a8,8,0,0,1,7.89,6.7A1.53,1.53,0,0,0,21.38,12h0a1.5,1.5,0,0,0,1.48-1.75,11,11,0,0,0-21.72,0A1.5,1.5,0,0,0,2.62,12h0a1.53,1.53,0,0,0,1.49-1.3A8,8,0,0,1,12,4Z`),
		),
	)
)

var (
	Config = struct {
		Spinner Node
	}{
		Spinner: defaultSpinner,
	}
)
