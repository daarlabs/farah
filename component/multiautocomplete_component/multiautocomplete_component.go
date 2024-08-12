package multiautocomplete_component

import (
	"slices"
	
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/tempest/form_tempest/form_input_tempest"
	"github.com/daarlabs/farah/ui/form_ui"
	"github.com/daarlabs/hirokit/dyna"
	"github.com/daarlabs/hirokit/esquel"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	. "github.com/daarlabs/hirokit/gox"
	
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/hirokit/hx"
	
	"github.com/daarlabs/farah/model/select_model"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/hidden_field_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/menu_ui"
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
)

type MultiAutocomplete[T comparable] struct {
	hiro.Component
	Props    Props[T]                 `json:"-"`
	Query    dyna.Query               `json:"-"`
	Options  []select_model.Option[T] `json:"-"`
	Selected []select_model.Option[T] `json:"-"`
	Offset   int                      `json:"-"`
}

func (c *MultiAutocomplete[T]) Name() string {
	return "multi-autocomplete-" + c.Props.Name
}

func (c *MultiAutocomplete[T]) Mount() {
	if !c.Request().Is().Action() {
		c.Options = c.find(dyna.Param{}.Parse(c))
		c.Selected = c.getSelected()
	}
}

func (c *MultiAutocomplete[T]) Node() Node {
	return c.createMultiAutocomplete(false)
}

func (c *MultiAutocomplete[T]) HandleChooseOption() error {
	c.Props.Value = make([]T, 0)
	c.Parse().MustQuery(dyna.Fulltext, &c.Props.Text)
	c.Parse().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	c.Options = c.find(dyna.Param{}.Parse(c))
	return c.Response().Render(c.createMultiAutocomplete(true))
}

func (c *MultiAutocomplete[T]) HandleSearch() error {
	c.Parse().MustQuery(dyna.Fulltext, &c.Props.Text)
	c.Parse().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	c.Options = c.find(dyna.Param{Fulltext: c.Props.Text})
	c.Offset = 0
	return c.Response().Render(
		c.createMultiAutocomplete(true),
	)
}

func (c *MultiAutocomplete[T]) HandleLoadMore() error {
	c.Parse().MustQuery(dyna.Offset, &c.Offset)
	c.Options = c.find(dyna.Param{}.Parse(c))
	c.Parse().MustQuery("value", &c.Props.Value)
	c.Selected = c.getSelected()
	return c.Response().Render(c.createOptions())
}

func (c *MultiAutocomplete[T]) find(param dyna.Param) []select_model.Option[T] {
	result := make([]select_model.Option[T], 0)
	param.Fields.Fulltext = []string{esquel.Vectors}
	param.Fields.Map = map[string]string{c.Query.Alias: c.Query.Value}
	q := dyna.New()
	if c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Options) > 0 {
		q = q.Data(select_model.ConvertToMapSlice(c.Options))
	}
	q.MustFind(param, &result)
	return result
}

func (c *MultiAutocomplete[T]) getSelected() []select_model.Option[T] {
	result := make([]select_model.Option[T], 0)
	if len(c.Props.Value) == 0 {
		return result
	}
	q := dyna.New()
	if c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Options) > 0 {
		q = q.Data(select_model.ConvertToMapSlice(c.Options))
	}
	
	q.MustFindMany(c.Query.Value, c.Props.Value, &result)
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
					tempest.Class().
						H(8).
						Extend(form_input_tempest.InputField(form_input_tempest.Props{})).
						Extend(form_tempest.FocusShadow()),
					Type("text"),
					Name(dyna.Fulltext),
					Value(c.Props.Text),
					If(c.Props.Text != "", CustomData("autofocus")),
					form_ui.Autofocus(),
					hx.Get(c.Generate().Action("HandleSearch", hiro.Map{"value": c.Props.Value})),
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
				tempest.Class().
					MinH(10).
					Py(2).
					Pr(8).
					Extend(form_input_tempest.InputField(form_input_tempest.Props{})).
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
		tempest.Class().Name(c.Request().Action()).Transition().Flex().ItemsCenter().Rounded().Px(1).Py(0.5).ShadowMain().
			TextXs().TextWhite().TextLeft().
			Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()),
		Text(title),
		A(
			tempest.Class().Name(c.Request().Action()).InlineFlex().Ml(1).CursorPointer(),
			hx.Get(
				c.Generate().Action(
					"HandleChooseOption", hiro.Map{"value": c.removeValue(value), "fulltext": c.Props.Text},
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
	return Fragment(
		If(
			len(c.Props.Value) > 0,
			Range(
				c.Props.Value, func(item T, _ int) Node {
					return hidden_field_ui.HiddenField(c.Props.Name, item)
				},
			),
		),
		If(
			len(c.Props.Value) == 0,
			hidden_field_ui.HiddenField(c.Props.Name, ""),
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
				"HandleLoadMore", hiro.Map{"offset": offset + dyna.DefaultLimit, "value": c.Props.Value},
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
				If((i+1)%dyna.DefaultLimit == 0, c.createLoadMore(c.Offset)),
				If(
					!exist,
					hx.Get(
						c.Generate().Action(
							"HandleChooseOption",
							hiro.Map{"value": append(c.Props.Value, option.Value), "fulltext": c.Props.Text},
						),
					),
				),
				If(
					exist,
					hx.Get(
						c.Generate().Action(
							"HandleChooseOption",
							hiro.Map{"value": c.removeValue(option.Value), "fulltext": c.Props.Text},
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
