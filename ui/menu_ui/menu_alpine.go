package menu_ui

import (
	"github.com/daarlabs/arcanum/alpine"
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func Open() gox.Node {
	return alpine.Click("open = true")
}

func Close() gox.Node {
	return alpine.Click("open = false")
}

func Chevron() gox.Node {
	return alpine.Bind("class", "open && '"+tempest.Class().Rotate(-180).String()+"'")
}
