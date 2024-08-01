package drawer_ui

import (
	"github.com/daarlabs/hirokit/alpine"
	"github.com/daarlabs/hirokit/gox"
)

func OpenEvent() gox.Node {
	return alpine.Click("open = !open")
}
