package flash_message_component

import (
	"github.com/daarlabs/farah/ui/page_ui/flash_message_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/hx"
	"github.com/daarlabs/hirokit/tempest"
)

type FlashMessage struct {
	hiro.Component
}

const (
	TargetId = "flash-messages"
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
		Id(hx.Id(TargetId)),
		tempest.Class().Fixed().Grid().Gap(4).Top(8).Right(8).W("200px").Z(9999),
		Range(
			c.Flash().MustGet(), func(item hiro.Message, _ int) Node {
				return flash_message_ui.Message(item)
			},
		),
	)
}
