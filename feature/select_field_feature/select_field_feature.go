package select_field_feature

import (
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/stimulus"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/arcanum/hx"
	
	"component/ui/form_ui/field_label_ui"
	"component/ui/form_ui/hidden_field_ui"
	
	"component/model/select_model"
	"component/ui"
	"component/ui/form_ui/error_message_ui"
	"component/ui/icon_ui"
	"component/ui/menu_ui"
	"component/ui/menu_ui/menu_item_ui"
)

type SelectField[T comparable] struct {
	mirage.Component
	Props   Props[T]                 `json:"-"`
	Options []select_model.Option[T] `json:"-"`
}

func (c *SelectField[T]) Name() string {
	return "select-" + c.Props.Name
}

func (c *SelectField[T]) Mount() {
}

func (c *SelectField[T]) Node() Node {
	return c.createSelectField(false)
}

func (c *SelectField[T]) HandleChooseOption() error {
	c.Parse().MustQuery("text", &c.Props.Text)
	c.Parse().MustQuery("value", &c.Props.Value)
	return c.Response().Render(c.createSelectField(false))
}

func (c *SelectField[T]) createSelectField(open bool) Node {
	return menu_ui.Menu(
		c.createMenuProps(open),
		Div(
			Class("flex flex-col gap-1 w-full"),
			If(
				len(c.Props.Label) > 0,
				field_label_ui.FieldLabel(
					field_label_ui.Props{
						For:      c.Props.Id,
						Text:     c.Props.Label,
						Required: c.Props.Required,
					},
				),
			),
			Div(
				Class("relative h-10"),
				Class("relative h-10"),
				stimulus.Action("click", "menu", "handleOpen"),
				Button(
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					Type("button"),
					Clsx{
						"transition w-full border pl-3 pr-7 rounded focus:shadow-focus text-left h-10 text-xs":                                      true,
						"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600 focus:border-primary-400": true,
						"text-slate-800 dark:text-white":     c.Props.Text != "",
						"text-slate-400 dark:text-slate-200": c.Props.Text == "" && c.Props.Placeholder != "",
					},
					If(c.Props.Text == "" && c.Props.Placeholder != "", Text(c.Props.Placeholder)),
					If(c.Props.Text != "", Text(c.Props.Text)),
				),
				Label(
					Class("transition absolute right-2 inset-y-0 my-auto h-4 w-4"),
					If(len(c.Props.Id) > 0, For(c.Props.Id)),
					stimulus.Target("menu", "chevron"),
					icon_ui.Icon(icon_ui.Props{Icon: icon_ui.ChevronDown, Size: ui.Sm, Class: "text-slate-900 dark:text-white"}),
				),
				hidden_field_ui.HiddenField(c.Props.Name, c.Props.Value),
			),
			Range(
				c.Props.Messages, func(msg string, _ int) Node {
					return error_message_ui.ErrorMessage(msg)
				},
			),
		),
		c.createOptions(),
	)
}

func (c *SelectField[T]) createMenuProps(open bool) menu_ui.Props {
	return menu_ui.Props{
		Id:         hx.Id(c.Props.Id),
		OptionsId:  c.Props.Name + "-" + "options",
		Fullwidth:  true,
		Clickable:  true,
		Scrollable: true,
		Open:       open,
		PositionX:  ui.Left,
		PositionY:  ui.Bottom,
	}
}

func (c *SelectField[T]) createOptions() Node {
	return Range(
		c.Options,
		func(option select_model.Option[T], i int) Node {
			return A(
				Class("cursor-pointer"),
				stimulus.Action("click", "menu", "handleClose"),
				hx.Get(c.Generate().Action("HandleChooseOption", mirage.Map{"value": option.Value, "text": option.Text})),
				hx.Target(hx.HashId(c.Props.Id)),
				hx.Swap(hx.SwapOuterHtml),
				hx.Trigger("click"),
				menu_item_ui.MenuItem(
					menu_item_ui.Props{Selected: c.Props.Value == option.Value}, Text(option.Text),
				),
			)
		},
	)
}
