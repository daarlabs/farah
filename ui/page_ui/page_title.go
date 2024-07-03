package page_ui

import (
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Title(title string) gox.Node {
	return gox.H1(
		tempest.Class().TextLg().FontBold().TextSlate(900).TextWhite(tempest.Dark()).Mb(4),
		gox.Text(title),
	)
}
