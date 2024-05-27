package nav_ui

import (
	. "github.com/daarlabs/arcanum/gox"
)

func NavImgLink(link, imgSrc string) Node {
	return A(
		Href(link),
		Img(Class("h-6"), Src(imgSrc)),
	)
}
