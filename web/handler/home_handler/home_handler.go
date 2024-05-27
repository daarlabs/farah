package home_handler

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
)

func Get() mirage.Handler {
	return func(c mirage.Ctx) error {
		return c.Response().Render(
			Div(Text("Home")),
		)
	}
}
