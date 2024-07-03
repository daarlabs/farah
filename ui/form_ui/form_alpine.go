package form_ui

import (
	"github.com/daarlabs/hirokit/alpine"
	"github.com/daarlabs/hirokit/gox"
)

func Autofocus() gox.Node {
	return alpine.Init("$el.focus();$el.setSelectionRange($el.value.length, $el.value.length);")
}
