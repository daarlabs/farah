package autocomplete_component

import (
	"fmt"
	"reflect"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/mystiq"
	"github.com/daarlabs/arcanum/quirk"
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
		c.Options = c.getAll(
			mystiq.Param{}.Parse(c),
		)
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
			tempest.Class().Relative().H(10),
			menu_ui.Open(),
			Input(
				If(len(c.Props.Id) > 0, Id(c.Props.Id)),
				tempest.Class().
					H(10).
					Extend(form_input_tempest.InputField(form_input_tempest.Props{})).
					Extend(form_tempest.FocusShadow()),
				Type("text"),
				Value(c.Props.Text),
				Name("fulltext"),
				hx.Get(c.Generate().Action("HandleSearch", c.Param)),
				hx.Trigger("input delay:500ms"),
				hx.Swap(hx.SwapOuterHtml),
				hx.Target(hx.HashId(c.Props.Name)),
				hx.Vals(fmt.Sprintf(`{"value":"%v"}`, c.Props.Value)),
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
				tempest.Class().Cursor("pointer"),
				If((i+1)%mystiq.DefaultLimit == 0, c.createLoadMore(c.Offset)),
				menu_ui.Close(),
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
	param.Order = []string{"text:" + mystiq.Asc}
	param.Fields.Fulltext = []string{quirk.Vectors}
	textField, ok := c.Query.Fields["text"]
	if ok {
		param.Fields.Order = map[string]string{"text": textField}
	}
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
	q = q.DB(c.DB(), c.Query)
	q.MustGetOne(c.Query.Value, c.Props.Value, &result)
	return result
}
