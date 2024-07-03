package multiselect_field_component

import (
	"slices"
	
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui/form_ui/hidden_field_ui"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	. "github.com/daarlabs/hirokit/gox"
	
	"github.com/daarlabs/hirokit/hx"
	
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	
	"github.com/daarlabs/farah/model/select_model"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/menu_ui"
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
)

type MultiSelectField[T comparable] struct {
	hiro.Component
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
	c.Parse().Multiple().MustQuery("value", &c.Props.Value)
	return c.Response().Render(c.createMultiSelectField(true))
}

func (c *MultiSelectField[T]) createMultiSelectField(open bool) Node {
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
				tempest.Class().Relative().MinH(10),
				menu_ui.Open(),
				Button(
					Type("button"),
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					tempest.Class().Transition().W("full").Py(3).Pl(3).Pr(7).Rounded().MinH(10).
						TextLeft().TextXs().TextSlate(900).TextWhite(tempest.Dark()).
						BgWhite().BgSlate(800, tempest.Dark()).
						Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
						BorderColor(palette.Primary, 400, tempest.Focus()).
						BorderColor(palette.Primary, 200, tempest.Focus(), tempest.Dark()).
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
		tempest.Class().Name(c.Request().Action()).Transition().Flex().ItemsCenter().Rounded().Px(1).Py(0.5).TextSize("10px").
			Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()).ShadowMain().TextWhite(),
		Text(title),
		A(
			tempest.Class().Name(c.Request().Action()).InlineFlex().Ml(1),
			hx.Get(
				c.Generate().Action("HandleChooseOption", hiro.Map{"value": c.removeValue(value)}),
			),
			hx.Target(hx.HashId(c.Props.Id)),
			hx.Swap(hx.SwapOuterHtml),
			hx.Trigger("click"),
			icon_ui.Icon(icon_ui.Props{Icon: icon_ui.Close, Size: ui.Sm, Class: tempest.Class().TextWhite()}),
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
				tempest.Class().CursorPointer(),
				If(
					!exist,
					hx.Get(c.Generate().Action("HandleChooseOption", hiro.Map{"value": append(c.Props.Value, option.Value)})),
				),
				If(
					exist,
					hx.Get(c.Generate().Action("HandleChooseOption", hiro.Map{"value": c.removeValue(option.Value)})),
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
