package form_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Field(nodes ...gox.Node) gox.Node {
	return gox.Div(
		tempest.Class().Mb(4),
		gox.Fragment(nodes...),
	)
}
