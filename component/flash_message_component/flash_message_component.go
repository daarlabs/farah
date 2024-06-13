package flash_message_component

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/hx"
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/ui/page_ui/flash_message_ui"
)

type FlashMessage struct {
	mirage.Component
}

const (
	id = "flash-messages"
)

func (c *FlashMessage) Name() string {
	return "flash-message"
}

func (c *FlashMessage) Mount() {
}

func (c *FlashMessage) Node() Node {
	return c.createFlashMessages()
}

func (c *FlashMessage) HandleRefresh() error {
	return c.Response().Render(c.createFlashMessages())
}

func (c *FlashMessage) createFlashMessages() Node {
	return Div(
		Id(hx.Id(id)),
		tempest.Class().Fixed().Grid().Gap(4).Top(8).Right(8).W("200px").Z(9999),
		Range(
			c.Flash().MustGet(), func(item mirage.Message, _ int) Node {
				return flash_message_ui.Message(item)
			},
		),
	)
}
