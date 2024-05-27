package autocomplete_feature

import (
	"fmt"
	"reflect"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/mystiq"
	"github.com/daarlabs/arcanum/quirk"
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

type Autocomplete[T comparable] struct {
	mirage.Component
	Props   Props[T]                 `json:"-"`
	Param   mirage.Map               `json:"-"`
	Query   mystiq.Query             `json:"-"`
	Options []select_model.Option[T] `json:"-"`
	Offset  int                      `json:"-"`
}

func (c *Autocomplete[T]) Name() string {
	return "autocomplete-" + c.Props.Name
}

func (c *Autocomplete[T]) Mount() {
	if !c.Request().Is().Action() {
		c.Options = c.getAll(mystiq.Param{}.Parse(c))
	}
	if !reflect.ValueOf(c.Props.Value).IsZero() && c.Props.Text == "" {
		c.Props.Text = c.getOne().Text
	}
}

func (c *Autocomplete[T]) Node() Node {
	return c.createAutocomplete(false)
}

func (c *Autocomplete[T]) HandleSearch() error {
	c.Parse().MustQuery("value", &c.Props.Value)
	c.Parse().MustQuery("fulltext", &c.Props.Text)
	c.Options = c.getAll(mystiq.Param{Fulltext: c.Props.Text})
	c.Offset = 0
	return c.Response().Render(
		c.createAutocomplete(true),
	)
}

func (c *Autocomplete[T]) HandleLoadMore() error {
	param := mystiq.Param{}.Parse(c)
	c.Offset = param.Offset
	c.Options = c.getAll(param)
	c.Parse().MustQuery("value", &c.Props.Value)
	return c.Response().Render(c.createOptions())
}

func (c *Autocomplete[T]) HandleChooseOption() error {
	c.Options = c.getAll(mystiq.Param{}.Parse(c))
	c.Parse().MustQuery("text", &c.Props.Text)
	c.Parse().MustQuery("value", &c.Props.Value)
	return c.Response().Render(c.createAutocomplete(false))
}

func (c *Autocomplete[T]) createAutocomplete(open bool) Node {
	return menu_ui.Menu(
		c.createMenuProps(open),
		c.createHandler(),
		c.createOptions(),
	)
}

func (c *Autocomplete[T]) createHandler() Node {
	return Div(
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
			stimulus.Action("click", "menu", "handleOpen"),
			Input(
				If(len(c.Props.Id) > 0, Id(c.Props.Id)),
				Clsx{
					"transition w-full border pl-3 pr-7 rounded focus:shadow-focus text-left h-10 text-xs":                                      true,
					"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600 focus:border-primary-400": true,
					"text-slate-800 dark:text-white": true,
				},
				Type("text"),
				Value(c.Props.Text),
				Name("fulltext"),
				stimulus.Controller("search"),
				hx.Get(c.Generate().Action("HandleSearch", c.Param)),
				hx.Trigger("input delay:500ms"),
				hx.Swap(hx.SwapOuterHtml),
				hx.Target(hx.HashId(c.Props.Name)),
				hx.Vals(fmt.Sprintf(`{"value":"%v"}`, c.Props.Value)),
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
	)
}

func (c *Autocomplete[T]) createMenuProps(open bool) menu_ui.Props {
	return menu_ui.Props{
		Id:         hx.Id(c.Props.Id),
		OptionsId:  hx.Id(c.Props.Name + "-" + "options"),
		Fullwidth:  true,
		Clickable:  true,
		Scrollable: true,
		Open:       open,
		PositionX:  ui.Left,
		PositionY:  ui.Bottom,
	}
}

func (c *Autocomplete[T]) createLoadMore(offset int) Node {
	return Fragment(
		hx.Get(c.Generate().Action("HandleLoadMore", mirage.Map{"offset": offset + mystiq.DefaultLimit}.Merge(c.Param))),
		hx.Target("#hx-"+c.Props.Name+"-options"),
		hx.Swap(hx.SwapBeforeEnd),
		hx.Trigger("intersect once"),
		hx.Vals(fmt.Sprintf(`{"value":"%v"}`, c.Props.Value)),
	)
}

func (c *Autocomplete[T]) createOptions() Node {
	return Range(
		c.Options,
		func(option select_model.Option[T], i int) Node {
			return A(
				Class("cursor-pointer"),
				If((i+1)%mystiq.DefaultLimit == 0, c.createLoadMore(c.Offset)),
				stimulus.Action("click", "menu", "handleClose"),
				hx.Get(
					c.Generate().Action(
						"HandleChooseOption", mirage.Map{"value": option.Value, "text": option.Text}.Merge(c.Param),
					),
				),
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

func (c *Autocomplete[T]) getAll(param mystiq.Param) []select_model.Option[T] {
	result := make([]select_model.Option[T], 0)
	param.Columns.Fulltext = []string{quirk.Vectors}
	param.Columns.Order = map[string]string{c.Query.Alias: c.Query.Name}
	q := mystiq.New()
	if len(c.Options) == 0 && c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Options) > 0 {
		q = q.Data(select_model.ConvertToMapSlice(c.Options))
	}
	q.MustGetAll(param, &result)
	return result
}

func (c *Autocomplete[T]) getOne() select_model.Option[T] {
	var result select_model.Option[T]
	q := mystiq.New()
	if len(c.Options) == 0 && c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Options) > 0 {
		q = q.Data(select_model.ConvertToMapSlice(c.Options))
	}
	q.MustGetOne(c.Query.Name, c.Props.Value, &result)
	return result
}
