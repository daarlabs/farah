package dialog_feature

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	
	"component/ui/button_ui"
	"github.com/daarlabs/arcanum/hx"
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
				Class("fixed inset-0 m-auto h-screen w-screen bg-slate-900/80 z-40"),
				c.createCloseAction(),
			),
			Div(
				Clsx{
					"fixed z-50 top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%]": true,
					"border border-slate-300 dark:border-slate-600":                         true,
					"bg-white dark:bg-slate-800":                                            true,
					"flex flex-col rounded min-w-[300px] min-h-[150px] max-w-[500px]":       true,
				},
				If(
					len(c.Props.Title) > 0,
					Div(
						Class("p-4 text-slate-900 dark:text-white font-bold"),
						Text(c.Props.Title),
					),
				),
				Fragment(c.Props.Nodes...),
				Div(
					Class("flex items-center justify-end p-4 gap-4 mt-auto"),
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
