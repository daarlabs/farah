package nav_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func NavImgLink(link, imgSrc string, nodes ...Node) Node {
	return A(
		Href(link),
		Img(tempest.Class().Block().H(6), Src(imgSrc)),
		Fragment(nodes...),
	)
}
