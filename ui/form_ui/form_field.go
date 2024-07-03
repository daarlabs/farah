package form_ui

import (
	"github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

func Field(nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Mb(6),
		gox.Fragment(nodes...),
	)
}
