package select_field_component

import (
	"reflect"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/form_tempest/form_input_tempest"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/arcanum/hx"
	
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/form_ui/hidden_field_ui"
	
	"github.com/daarlabs/farah/model/select_model"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/menu_ui"
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
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
	if !reflect.ValueOf(c.Props.Value).IsZero() && c.Props.Text == "" {
		for _, option := range c.Options {
			if c.Props.Value == option.Value {
				c.Props.Text = option.Text
				break
			}
		}
	}
}

func (c *SelectField[T]) Node() Node {
	return c.createSelectField(false)
}

func (c *SelectField[T]) HandleChooseOption() error {
	c.Parse().MustQuery("text", &c.Props.Text)
	c.Parse().MustQuery("value", &c.Props.Value)
	if c.Props.OnChange != nil {
		c.Props.OnChange(select_model.Option[T]{Value: c.Props.Value, Text: c.Props.Text})
	}
	if c.Props.Refresh {
		return c.Response().Refresh()
	}
	return c.Response().Render(c.createSelectField(false))
}

func (c *SelectField[T]) createSelectField(open bool) Node {
	handlerHeight := c.createHandlerHeight()
	return menu_ui.Menu(
		c.createMenuProps(open),
		Div(
			tempest.Class().Flex().FlexCol().Gap(1).W("full"),
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
				tempest.Class().Relative().H(handlerHeight),
				menu_ui.Open(),
				Button(
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					Type("button"),
					tempest.Class().
						H(handlerHeight).
						TextLeft().
						Extend(
							form_input_tempest.InputField(
								form_input_tempest.Props{
									Text:        c.Props.Text,
									Placeholder: c.Props.Placeholder,
								},
							),
						).
						Extend(form_tempest.FocusShadow()),
					If(c.Props.Text == "" && c.Props.Placeholder != "", Text(c.Props.Placeholder)),
					If(c.Props.Text != "", Text(c.Props.Text)),
				),
				Label(
					tempest.Class().Transition().Absolute().Right(2).InsetY(0).My("auto").Size(4),
					If(len(c.Props.Id) > 0, For(c.Props.Id)),
					menu_ui.Chevron(),
					icon_ui.Icon(
						icon_ui.Props{
							Icon: icon_ui.ChevronDown, Size: ui.Sm, Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
						},
					),
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
			action := c.Generate().Action("HandleChooseOption", mirage.Map{"value": option.Value, "text": option.Text})
			return A(
				tempest.Class().CursorPointer(),
				menu_ui.Close(),
				If(
					c.Props.Refresh,
					Href(action),
				),
				If(
					!c.Props.Refresh,
					hx.Get(action),
					hx.Target(hx.HashId(c.Props.Id)),
					hx.Swap(hx.SwapOuterHtml),
					hx.Trigger("click"),
				),
				menu_item_ui.MenuItem(
					menu_item_ui.Props{Selected: c.Props.Value == option.Value}, Text(option.Text),
				),
			)
		},
	)
}

func (c *SelectField[T]) createHandlerHeight() int {
	if c.Props.Size == ui.Sm {
		return 8
	}
	return 10
}
