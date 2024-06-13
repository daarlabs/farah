package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func NavImgLink(link, imgSrc string) Node {
	return A(
		Href(link),
		Img(tempest.Class().Block().H(6), Src(imgSrc)),
	)
}
