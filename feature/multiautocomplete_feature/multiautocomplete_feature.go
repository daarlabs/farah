package multiautocomplete_feature

import (
	"slices"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/mystiq"
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/stimulus"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"component/ui/form_ui/field_label_ui"
	"github.com/daarlabs/arcanum/hx"
	
	"component/model/select_model"
	"component/ui"
	"component/ui/form_ui/error_message_ui"
	"component/ui/form_ui/hidden_field_ui"
	"component/ui/icon_ui"
	"component/ui/menu_ui"
	"component/ui/menu_ui/menu_item_ui"
)

type MultiAutocomplete[T comparable] struct {
	mirage.Component
	Props    Props[T]                 `json:"-"`
	Query    mystiq.Query             `json:"-"`
	Options  []select_model.Option[T] `json:"-"`
	Selected []select_model.Option[T] `json:"-"`
	Offset   int                      `json:"-"`
}

func (c *MultiAutocomplete[T]) Name() string {
	return "multi-autocomplete-" + c.Props.Name
}

func (c *MultiAutocomplete[T]) Mount() {
	if !c.Request().Is().Action() {
		c.Options = c.getAll(mystiq.Param{}.Parse(c))
		c.Selected = c.getSelected()
	}
}

func (c *MultiAutocomplete[T]) Node() Node {
	return c.createMultiAutocomplete(false)
}

func (c *MultiAutocomplete[T]) HandleChooseOption() error {
	c.Parse().MustQuery(mystiq.Fulltext, &c.Props.Text)
	c.Parse().Many().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	c.Options = c.getAll(mystiq.Param{}.Parse(c))
	return c.Response().Render(c.createMultiAutocomplete(true))
}

func (c *MultiAutocomplete[T]) HandleSearch() error {
	c.Parse().MustQuery(mystiq.Fulltext, &c.Props.Text)
	c.Parse().Many().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	c.Options = c.getAll(mystiq.Param{Fulltext: c.Props.Text})
	c.Offset = 0
	return c.Response().Render(
		c.createMultiAutocomplete(true),
	)
}

func (c *MultiAutocomplete[T]) HandleLoadMore() error {
	c.Parse().MustQuery(mystiq.Offset, &c.Offset)
	c.Options = c.getAll(mystiq.Param{}.Parse(c))
	c.Parse().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	return c.Response().Render(c.createOptions())
}

func (c *MultiAutocomplete[T]) getAll(param mystiq.Param) []select_model.Option[T] {
	result := make([]select_model.Option[T], 0)
	param.Columns.Fulltext = []string{quirk.Vectors}
	param.Columns.Order = map[string]string{c.Query.Alias: c.Query.Name}
	q := mystiq.New()
	if c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Options) > 0 {
		q = q.Data(select_model.ConvertToMapSlice(c.Options))
	}
	q.MustGetAll(param, &result)
	return result
}

func (c *MultiAutocomplete[T]) getSelected() []select_model.Option[T] {
	result := make([]select_model.Option[T], 0)
	if len(c.Props.Value) == 0 {
		return result
	}
	q := mystiq.New()
	if c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Options) > 0 {
		q = q.Data(select_model.ConvertToMapSlice(c.Options))
	}
	
	q.MustGetMany(c.Query.Name, c.Props.Value, &result)
	return result
}

func (c *MultiAutocomplete[T]) createMultiAutocomplete(open bool) Node {
	return menu_ui.Menu(
		c.createMenuProps(open),
		c.createHandler(),
		Div(
			Class("grid grid-rows-[3.5rem_1fr] h-full overflow-hidden"),
			Div(
				Class("flex items-center px-3 h-full border-b border-slate-300 dark:border-slate-600"),
				Input(
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					Clsx{
						"transition w-full border h-8 pl-3 pr-7 rounded text-[10px] focus:shadow-focus":                                             true,
						"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600 focus:border-primary-400": true,
						"text-slate-800 dark:text-white": true,
					},
					Type("text"),
					Name(mystiq.Fulltext),
					Value(c.Props.Text),
					If(c.Props.Text != "", CustomData("autofocus")),
					stimulus.Controller("search"),
					hx.Get(c.Generate().Action("HandleSearch", mirage.Map{"value": c.Props.Value})),
					hx.Trigger("input delay:500ms"),
					hx.Swap(hx.SwapOuterHtml),
					hx.Target(hx.HashId(c.Props.Name)),
				),
			),
			Div(
				Class("h-full overflow-auto"),
				c.createOptions(),
			),
		),
	)
}

func (c *MultiAutocomplete[T]) createHandler() Node {
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
			Class("relative min-h-[2.5rem]"),
			stimulus.Action("click", "menu", "handleOpen"),
			Button(
				Type("button"),
				If(len(c.Props.Id) > 0, Id(c.Props.Id)),
				Clsx{
					"transition w-full border py-2 pl-3 pr-7 rounded focus:shadow-focus text-left min-h-[2.5rem] text-xs":                       true,
					"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600 focus:border-primary-400": true,
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
	)
}

func (c *MultiAutocomplete[T]) createSelectedTag(title string, value T) Node {
	return Div(
		Class("transition flex items-center bg-primary-400 dark:bg-primary-200 rounded py-0.5 px-1 text-xs text-white shadow"),
		Text(title),
		A(
			Class("inline-flex ml-1 cursor-pointer"),
			hx.Get(
				c.Generate().Action(
					"HandleChooseOption", mirage.Map{"value": c.removeValue(value), "fulltext": c.Props.Text},
				),
			),
			hx.Target(hx.HashId(c.Props.Id)),
			hx.Swap(hx.SwapOuterHtml),
			hx.Trigger("click"),
			icon_ui.Icon(icon_ui.Props{Icon: icon_ui.Close, Size: ui.Sm}),
		),
	)
}

func (c *MultiAutocomplete[T]) createHiddens() Node {
	return If(
		len(c.Props.Value) > 0,
		Range(
			c.Props.Value, func(item T, _ int) Node {
				return hidden_field_ui.HiddenField(c.Props.Name, item)
			},
		),
	)
}

func (c *MultiAutocomplete[T]) createMenuProps(open bool) menu_ui.Props {
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

func (c *MultiAutocomplete[T]) createLoadMore(offset int) Node {
	return Fragment(
		hx.Get(
			c.Generate().Action(
				"HandleLoadMore", mirage.Map{"offset": offset + mystiq.DefaultLimit, "value": c.Props.Value},
			),
		),
		hx.Target(hx.HashId(c.Props.Name+"-options")),
		hx.Swap(hx.SwapBeforeEnd),
		hx.Trigger("intersect once"),
	)
}

func (c *MultiAutocomplete[T]) createOptions() Node {
	return Range(
		c.Options,
		func(option select_model.Option[T], i int) Node {
			exist := slices.Contains(c.Props.Value, option.Value)
			return A(
				Class("cursor-pointer"),
				If((i+1)%mystiq.DefaultLimit == 0, c.createLoadMore(c.Offset)),
				If(
					!exist,
					hx.Get(
						c.Generate().Action(
							"HandleChooseOption",
							mirage.Map{"value": append(c.Props.Value, option.Value), "fulltext": c.Props.Text},
						),
					),
				),
				If(
					exist,
					hx.Get(
						c.Generate().Action(
							"HandleChooseOption",
							mirage.Map{"value": c.removeValue(option.Value), "fulltext": c.Props.Text},
						),
					),
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

func (c *MultiAutocomplete[T]) removeValue(value T) []T {
	result := make([]T, 0)
	for _, item := range c.Props.Value {
		if item != value {
			result = append(result, item)
		}
	}
	return result
}

func (c *MultiAutocomplete[T]) getOptionTitleWithValue(value T) string {
	for _, option := range c.Selected {
		if option.Value == value {
			return option.Text
		}
	}
	return ""
}
