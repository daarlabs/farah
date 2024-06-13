package dialog_component

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/tempest"
	
	"github.com/daarlabs/arcanum/hx"
	"github.com/daarlabs/farah/ui/button_ui"
)

type DialogFeature struct {
	mirage.Component
	Props       Props                  `json:"-"`
	Cancel      Config                 `json:"-"`
	Submit      Config                 `json:"-"`
	HandlerFunc func(action Node) Node `json:"-"`
}

func (c *DialogFeature) Name() string {
	return c.Props.Name + "-dialog"
}

func (c *DialogFeature) Mount() {
	if len(c.Cancel.Title) == 0 {
		c.Cancel.Title = "Cancel"
	}
	if len(c.Submit.Title) == 0 {
		c.Submit.Title = "Submit"
	}
}

func (c *DialogFeature) Node() Node {
	return c.createDialog()
}

func (c *DialogFeature) HandleOpen() error {
	c.Props.Open = true
	return c.Response().Render(c.createDialog())
}

func (c *DialogFeature) HandleClose() error {
	c.Props.Open = false
	return c.Response().Render(c.createDialog())
}

func (c *DialogFeature) createDialog() Node {
	return Div(
		Id(hx.Id(c.Name())),
		If(c.HandlerFunc != nil, c.HandlerFunc(c.createOpenAction())),
		If(
			c.Props.Open,
			Div(
				tempest.Class().Fixed().Inset(0).M("auto").H("screen").W("screen").
					BgSlate(900, tempest.Opacity(80)).Z(40),
				c.createCloseAction(),
			),
			Div(
				tempest.Class().Fixed().Z(50).Top("50%").Left("50%").TranslateX("-50%").TranslateY("-50%").
					Flex().FlexCol().Rounded().MinW("300px").MinH("150px").MaxW("500px").
					Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
					BgWhite().BgSlate(800, tempest.Dark()),
				If(
					len(c.Props.Title) > 0,
					Div(
						tempest.Class().P(4).TextSlate(900).TextWhite(tempest.Dark()).FontBold(),
						Text(c.Props.Title),
					),
				),
				Fragment(c.Props.Nodes...),
				Div(
					tempest.Class().Flex().ItemsCenter().JustifyEnd().P(4).Gap(4).Mt("auto"),
					c.createCancelButton(),
					c.createSubmitButton(),
				),
			),
		),
	)
}

func (c *DialogFeature) createOpenAction() Node {
	return Fragment(
		hx.Get(c.Generate().Action("HandleOpen")),
		hx.Swap(hx.SwapOuterHtml),
		hx.Trigger("click"),
		hx.Target(hx.HashId(c.Name())),
	)
}

func (c *DialogFeature) createCloseAction() Node {
	return Fragment(
		hx.Get(c.Generate().Action("HandleClose")),
		hx.Swap(hx.SwapOuterHtml),
		hx.Trigger("click"),
		hx.Target(hx.HashId(c.Name())),
	)
}

func (c *DialogFeature) createCancelButton() Node {
	return button_ui.MainButton(
		button_ui.Props{},
		c.createCloseAction(),
		Text(c.Cancel.Title),
	)
}

func (c *DialogFeature) createSubmitButton() Node {
	return button_ui.PrimaryButton(
		button_ui.Props{Link: c.Submit.Link, Type: button_ui.TypeSubmit},
		Text(c.Submit.Title),
	)
}
