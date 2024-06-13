package form_ui

import (
	"github.com/daarlabs/arcanum/alpine"
	"github.com/daarlabs/arcanum/gox"
)

func Autofocus() gox.Node {
	return alpine.Init("$el.focus();$el.setSelectionRange($el.value.length, $el.value.length);")
}
