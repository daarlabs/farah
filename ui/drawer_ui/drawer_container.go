package drawer_ui

import (
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
)

func Container(nodes ...Node) Node {
	return Div(
		alpine.Data(hiro.Map{"open": false}),
		Fragment(nodes...),
	)
}
