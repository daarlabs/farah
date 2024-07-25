package dialog_component

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui/button_ui"
	"github.com/daarlabs/hirokit/hx"
)

type DialogComponent struct {
	hiro.Component
	Props       Props                  `json:"-"`
	Cancel      Config                 `json:"-"`
	Submit      Config                 `json:"-"`
	HandlerFunc func(action Node) Node `json:"-"`
}

func (c *DialogComponent) Name() string {
	return c.Props.Name + "-dialog"
}

func (c *DialogComponent) Mount() {
	if len(c.Cancel.Title) == 0 {
		c.Cancel.Title = c.getCancelTitle()
	}
	if len(c.Submit.Title) == 0 {
		c.Submit.Title = c.getSubmitTitle()
	}
}

func (c *DialogComponent) Node() Node {
	return c.createDialog()
}

func (c *DialogComponent) HandleOpen() error {
	c.Props.Open = true
	return c.Response().Render(c.createDialog())
}

func (c *DialogComponent) HandleClose() error {
	c.Props.Open = false
	return c.Response().Render(c.createDialog())
}

func (c *DialogComponent) createDialog() Node {
	var handler Node
	if c.HandlerFunc != nil {
		handler = c.HandlerFunc(c.createOpenAction())
	}
	return Div(
		Id(hx.Id(c.Name())),
		If(handler != nil, handler),
		If(
			c.Props.Open,
			Div(
				tempest.Class().Name(c.Request().Action()).
					Fixed().Inset(0).M("auto").H("screen").W("screen").
					BgSlate(900, tempest.Opacity(80)).Z(40),
				c.createCloseAction(),
			),
			Div(
				tempest.Class().Name(c.Request().Action()).
					Fixed().Z(50).Top("50%").Left("50%").TranslateX("-50%").TranslateY("-50%").
					Flex().FlexCol().Rounded().MinW("300px").MaxW("500px").
					Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
					BgWhite().BgSlate(800, tempest.Dark()),
				If(
					len(c.Props.Title) > 0,
					Div(
						tempest.Class().Name(c.Request().Action()).
							P(4).TextSlate(900).TextWhite(tempest.Dark()).FontBold(),
						Text(c.Props.Title),
					),
				),
				c.Props.Content,
				If(
					c.Props.Submitable,
					Div(
						tempest.Class().Name(c.Request().Action()).
							Flex().ItemsCenter().JustifyEnd().P(4).Gap(4).Mt("auto"),
						c.createCancelButton(),
						c.createSubmitButton(),
					),
				),
			),
		),
	)
}

func (c *DialogComponent) createOpenAction() Node {
	return Fragment(
		hx.Get(c.Generate().Action("HandleOpen")),
		hx.Swap(hx.SwapOuterHtml),
		hx.Trigger("click"),
		hx.Target(hx.HashId(c.Name())),
	)
}

func (c *DialogComponent) createCloseAction() Node {
	return Fragment(
		hx.Get(c.Generate().Action("HandleClose")),
		hx.Swap(hx.SwapOuterHtml),
		hx.Trigger("click"),
		hx.Target(hx.HashId(c.Name())),
	)
}

func (c *DialogComponent) createCancelButton() Node {
	if !c.Props.Submitable {
		return Fragment()
	}
	return button_ui.MainButton(
		button_ui.Props{},
		c.createCloseAction(),
		Text(c.Cancel.Title),
	)
}

func (c *DialogComponent) createSubmitButton() Node {
	if !c.Props.Submitable {
		return Fragment()
	}
	return button_ui.PrimaryButton(
		button_ui.Props{Link: c.Submit.Link, Type: button_ui.TypeSubmit},
		Text(c.Submit.Title),
	)
}

func (c *DialogComponent) getCancelTitle() string {
	if !c.Config().Localization.Enabled {
		return "Cancel"
	}
	return c.Translate("button.cancel")
}

func (c *DialogComponent) getSubmitTitle() string {
	if !c.Config().Localization.Enabled {
		return "Submit"
	}
	return c.Translate("button.submit")
}
