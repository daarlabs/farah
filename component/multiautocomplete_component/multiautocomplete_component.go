package multiautocomplete_component

import (
	"slices"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/mystiq"
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui/form_ui"
	
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
	c.Parse().Multiple().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	c.Options = c.getAll(mystiq.Param{}.Parse(c))
	return c.Response().Render(c.createMultiAutocomplete(true))
}

func (c *MultiAutocomplete[T]) HandleSearch() error {
	c.Parse().MustQuery(mystiq.Fulltext, &c.Props.Text)
	c.Parse().Multiple().MustQuery("value", &c.Props.Value)
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
	param.Fields.Fulltext = []string{quirk.Vectors}
	param.Fields.Order = map[string]string{c.Query.Alias: c.Query.Value}
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
	
	q.MustGetMany(c.Query.Value, c.Props.Value, &result)
	return result
}

func (c *MultiAutocomplete[T]) createMultiAutocomplete(open bool) Node {
	return menu_ui.Menu(
		c.createMenuProps(open),
		c.createHandler(),
		Div(
			tempest.Class().Grid().GridRows("3.5rem 1fr").H("full").Overflow("hidden"),
			Div(
				tempest.Class().Flex().ItemsCenter().Px(3).H("full").
					BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
				Input(
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					tempest.Class().Transition().W("full").H(8).Pl(3).Pr(7).Rounded().TextSize("10px").
						BgWhite().BgSlate(800, tempest.Dark()).
						TextSlate(900).TextWhite(tempest.Dark()).
						Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
						BorderColor(palette.Primary, 400, tempest.Focus()).
						Extend(form_tempest.FocusShadow()),
					Type("text"),
					Name(mystiq.Fulltext),
					Value(c.Props.Text),
					If(c.Props.Text != "", CustomData("autofocus")),
					form_ui.Autofocus(),
					hx.Get(c.Generate().Action("HandleSearch", mirage.Map{"value": c.Props.Value})),
					hx.Trigger("input delay:500ms"),
					hx.Swap(hx.SwapOuterHtml),
					hx.Target(hx.HashId(c.Props.Name)),
				),
			),
			Div(
				tempest.Class().H("full").Overflow("auto"),
				c.createOptions(),
			),
		),
	)
}

func (c *MultiAutocomplete[T]) createHandler() Node {
	return Div(
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
			tempest.Class().Relative().MinH(2.5),
			menu_ui.Open(),
			Button(
				Type("button"),
				If(len(c.Props.Id) > 0, Id(c.Props.Id)),
				tempest.Class().Transition().W("full").Py(2).Pl(3).Pr(7).Rounded().MinH(2.5).
					TextLeft().TextXs().TextSlate(900).TextWhite(tempest.Dark()).
					BgWhite().BgSlate(800, tempest.Dark()).
					Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
					BorderColor(palette.Primary, 400, tempest.Focus()).
					Extend(form_tempest.FocusShadow()),
				Div(
					tempest.Class().Flex().FlexWrap().Gap(1),
					Range(
						c.Props.Value, func(value T, _ int) Node {
							return c.createSelectedTag(c.getOptionTitleWithValue(value), value)
						},
					),
				),
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
		tempest.Class().Transition().Flex().ItemsCenter().Rounded().Px(1).Py(0.5).ShadowMain().
			TextXs().TextWhite().
			Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()),
		Text(title),
		A(
			tempest.Class().InlineFlex().Ml(1).CursorPointer(),
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
				tempest.Class().CursorPointer(),
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
					menu_item_ui.Props{Selected: slices.Contains(c.Props.Value, option.Value)},
					Text(option.Text),
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
