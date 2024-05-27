package multiselect_field_feature

import (
	"slices"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/stimulus"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/arcanum/hx"
	
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	
	"github.com/daarlabs/farah/model/select_model"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/hidden_field_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/menu_ui"
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
)

type MultiSelectField[T comparable] struct {
	mirage.Component
	Props   Props[T]                 `json:"-"`
	Options []select_model.Option[T] `json:"-"`
}

func (c *MultiSelectField[T]) Name() string {
	return "multi-select-" + c.Props.Name
}

func (c *MultiSelectField[T]) Mount() {
}

func (c *MultiSelectField[T]) Node() Node {
	return c.createMultiSelectField(false)
}

func (c *MultiSelectField[T]) HandleChooseOption() error {
	c.Parse().Many().MustQuery("value", &c.Props.Value)
	return c.Response().Render(c.createMultiSelectField(true))
}

func (c *MultiSelectField[T]) createMultiSelectField(open bool) Node {
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
				Class("relative min-h-[2.5rem]"),
				stimulus.Action("click", "menu", "handleOpen"),
				Button(
					Type("button"),
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					Clsx{
						"transition w-full border border-slate-300 focus:border-primary-400 min-h-[2.5rem] py-3 pl-3 pr-7 rounded text-sm focus:shadow-focus text-left cursor-pointer": true,
						"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600 focus:border-primary-400":                                    true,
						"text-slate-800 dark:text-white": true,
					},
					Div(
						Class("flex flex-wrap gap-1"),
						Range(
							c.Props.Value, func(value T, _ int) Node {
								return c.createSelectedTag(c.getOptionTitleWithValue(value), value)
							},
						),
					),
				),
				Label(
					Class("transition absolute right-2 inset-y-0 my-auto h-4 w-4"),
					If(len(c.Props.Id) > 0, For(c.Props.Id)),
					stimulus.Target("menu", "chevron"),
					icon_ui.Icon(icon_ui.Props{Icon: icon_ui.ChevronDown, Size: ui.Sm, Class: "text-slate-900 dark:text-white"}),
				),
				c.createHiddens(),
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

func (c *MultiSelectField[T]) getOptionTitleWithValue(value T) string {
	for _, option := range c.Options {
		if option.Value == value {
			return option.Text
		}
	}
	return ""
}

func (c *MultiSelectField[T]) createHiddens() Node {
	return If(
		len(c.Props.Value) > 0,
		Range(
			c.Props.Value, func(item T, _ int) Node {
				return hidden_field_ui.HiddenField(c.Props.Name, item)
			},
		),
	)
}

func (c *MultiSelectField[T]) removeValue(value T) []T {
	result := make([]T, 0)
	for _, item := range c.Props.Value {
		if item != value {
			result = append(result, item)
		}
	}
	return result
}

func (c *MultiSelectField[T]) createSelectedTag(title string, value T) Node {
	return Div(
		Class("transition flex items-center bg-primary-400 dark:bg-primary-200 rounded px-1.5 text-[10px] text-white shadow"),
		Text(title),
		A(
			Class("inline-flex ml-1"),
			hx.Get(
				c.Generate().Action("HandleChooseOption", mirage.Map{"value": c.removeValue(value)}),
			),
			hx.Target(hx.HashId(c.Props.Id)),
			hx.Swap(hx.SwapOuterHtml),
			hx.Trigger("click"),
			icon_ui.Icon(icon_ui.Props{Icon: icon_ui.Close, Size: ui.Sm, Class: "text-white"}),
		),
	)
}

func (c *MultiSelectField[T]) createMenuProps(open bool) menu_ui.Props {
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

func (c *MultiSelectField[T]) createOptions() Node {
	return Range(
		c.Options,
		func(option select_model.Option[T], i int) Node {
			exist := slices.Contains(c.Props.Value, option.Value)
			return A(
				Class("cursor-pointer"),
				If(
					!exist,
					hx.Get(c.Generate().Action("HandleChooseOption", mirage.Map{"value": append(c.Props.Value, option.Value)})),
				),
				If(
					exist,
					hx.Get(c.Generate().Action("HandleChooseOption", mirage.Map{"value": c.removeValue(option.Value)})),
				),
				hx.Target(hx.HashId(c.Props.Id)),
				hx.Swap(hx.SwapOuterHtml),
				hx.Trigger("click"),
				menu_item_ui.MenuItem(
					menu_item_ui.Props{Selected: slices.Contains(c.Props.Value, option.Value)}, Text(option.Text),
				),
			)
		},
	)
}
