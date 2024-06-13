package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Title(title string) gox.Node {
	return gox.H1(
		tempest.Class().TextXxl().FontBold().TextSlate(900).TextWhite(tempest.Dark()).Mb(6),
		gox.Text(title),
	)
}
