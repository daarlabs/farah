package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Subtitle(title string) gox.Node {
	return gox.H1(
		tempest.Class().TextLg().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()).Mb(4),
		gox.Text(title),
	)
}
