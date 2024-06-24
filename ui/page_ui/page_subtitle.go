package page_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Subtitle(title string) gox.Node {
	return gox.H2(
		tempest.Class().TextSlate(700).TextSlate(200, tempest.Dark()).Mb(4).TextSm().FontNormal(),
		gox.Text(title),
	)
}
