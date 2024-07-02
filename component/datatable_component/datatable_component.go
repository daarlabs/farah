package datatable_component

import (
	"slices"
	"strings"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/mystiq"
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/ui/box_ui"
	"github.com/daarlabs/farah/ui/form_ui"
	"github.com/daarlabs/farah/ui/spinner_ui"
	
	"github.com/daarlabs/arcanum/hx"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/search_ui"
	
	. "github.com/daarlabs/arcanum/gox"
)

type Datatable[T any] struct {
	mirage.Component
	Props       Props                                 `json:"-"`
	Param       mystiq.Param                          `json:"-"`
	Query       mystiq.Query                          `json:"-"`
	GetDataFunc func(param mystiq.Param, t any) error `json:"-"`
	FieldsFunc  func() []Field                        `json:"-"`
	RowFunc     func(builder RowBuilder[T]) Node      `json:"-"`
	Data        []T                                   `json:"-"`
	
	fields []Field
}

func (c *Datatable[T]) Name() string {
	return "datatable"
}

func (c *Datatable[T]) Mount() {
	if c.FieldsFunc != nil {
		c.fields = c.FieldsFunc()
	}
	if !c.Request().Is().Action() {
		c.Data = c.getData()
	}
}

func (c *Datatable[T]) Node() Node {
	return c.createDatatable()
}

func (c *Datatable[T]) HandleOrder() error {
	c.Param = c.Param.Parse(c)
	c.Data = c.getData()
	return c.Response().Render(c.createDatatable())
}

func (c *Datatable[T]) HandleSearch() error {
	c.Parse().MustQuery("fulltext", &c.Param.Fulltext)
	c.Param.Offset = 0
	c.Data = c.getData()
	return c.Response().Render(c.createDatatable())
}

func (c *Datatable[T]) HandleLoadMore() error {
	c.Param = c.Param.Parse(c)
	c.Data = c.getData()
	return c.Response().Render(c.createRows())
}

func (c *Datatable[T]) getData() []T {
	param := c.Param
	result := make([]T, 0)
	param.Fields.Fulltext = []string{quirk.Vectors}
	param.Fields.Order = c.Query.Fields
	q := mystiq.New()
	if c.GetDataFunc != nil {
		q = q.GetAllFunc(c.GetDataFunc)
	}
	if c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Data) > 0 {
		q = q.Data(convertDataToMapSlice(c.Data))
	}
	q.MustGetAll(param, &result)
	return result
}

func (c *Datatable[T]) createDatatable() Node {
	return Div(
		tempest.Class().H("full").Grid().GridRows("2.5rem 1fr").Gap(4),
		Id(hx.Id(c.Props.Name)),
		Div(
			tempest.Class().MaxW("300px"),
			search_ui.Search(
				search_ui.Props{
					Placeholder: "Search", Value: c.Param.Fulltext, Name: mystiq.Fulltext,
					Id: hx.Id(c.Props.Name + "-" + mystiq.Fulltext),
				},
				c.createFulltext(),
			),
		),
		box_ui.Box(
			box_ui.Props{
				Class: tempest.Class().W("full").H("full"),
			},
			Div(
				tempest.Class().Grid().GridRows("2.5rem 1fr").H("full").
					OverflowY("hidden").OverflowX("auto"),
				c.createHead(),
				c.createBody(),
			),
		),
	)
}

func (c *Datatable[T]) createHead() Node {
	return Div(
		tempest.Class().Transition().H(10).W("full").BorderB(1).
			BorderColor(palette.Primary, 400).BorderColor(palette.Primary, 200, tempest.Dark()),
		c.createSizeStyle(),
		Range(
			c.fields, func(field Field, _ int) Node {
				el := "a"
				if !field.Sortable {
					el = "div"
				}
				return CreateElement(el)(
					tempest.Class().Relative().Transition().Flex().ItemsCenter().Gap(1).H("full").Px(4).CursorPointer().
						TextXs().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()).
						BgSlate(100, tempest.Hover()).BgSlate(700, tempest.Dark(), tempest.Hover()).
						If(field.AlignX == ui.Left, tempest.Class().JustifyStart()).
						If(field.AlignX == ui.Center, tempest.Class().JustifyCenter()).
						If(field.AlignX == ui.Right, tempest.Class().JustifyEnd()),
					Text(field.Title),
					If(
						field.Sortable,
						spinner_ui.Spinner(
							spinner_ui.Props{Overlay: true, Class: tempest.Class(spinner_ui.Indicator)},
						),
						c.createOrder(mirage.Map{mystiq.Fulltext: c.Param.Fulltext, mystiq.Order: c.createNextOrder(field.Name)}),
						If(
							slices.Contains(c.Param.Order, field.Name+":"+mystiq.Asc),
							icon_ui.Icon(
								icon_ui.Props{
									Icon:  icon_ui.ChevronUp,
									Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
									Size:  ui.Sm,
								},
							),
						),
						If(
							slices.Contains(c.Param.Order, field.Name+":"+mystiq.Desc),
							icon_ui.Icon(
								icon_ui.Props{
									Icon:  icon_ui.ChevronDown,
									Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
									Size:  ui.Sm,
								},
							),
						),
					),
				)
			},
		),
	)
}

func (c *Datatable[T]) createBody() Node {
	return Div(
		Id(hx.Id(c.Props.Name+"-rows")),
		tempest.Class().OverflowY("scroll").H("full"),
		c.createRows(),
	)
}

func (c *Datatable[T]) createRows() Node {
	sizeStyle := c.createSizeStyle()
	return If(
		len(c.Data) > 0 && c.RowFunc != nil,
		Range(
			c.Data,
			func(item T, index int) Node {
				var loadMore Node
				if (index+1)%mystiq.DefaultLimit == 0 {
					loadMore = c.createLoadMore()
				}
				return c.RowFunc(
					&rowBuilder[T]{
						data: item, fields: c.fields, sizeStyle: sizeStyle, rowIndex: index, loadMore: loadMore,
					},
				)
			},
		),
	)
}

func (c *Datatable[T]) createNextOrder(name string) []string {
	result := make([]string, 0)
	var exist bool
	for _, o := range c.Param.Order {
		if strings.HasPrefix(o, name+":") {
			exist = true
			break
		}
	}
	if !exist {
		return append(c.Param.Order, name+":"+mystiq.Asc)
	}
	if exist {
		for _, o := range c.Param.Order {
			isName := strings.HasPrefix(o, name+":")
			if !isName {
				result = append(result, o)
				continue
			}
			if isName {
				if strings.HasSuffix(o, mystiq.Asc) {
					result = append(result, strings.Replace(o, mystiq.Asc, mystiq.Desc, 1))
				}
				continue
			}
		}
	}
	return result
}

func (c *Datatable[T]) createSizeStyle() Node {
	n := len(c.fields)
	if n == 0 {
		return Fragment()
	}
	size := make([]string, n)
	for i, f := range c.fields {
		size[i] = f.Size
	}
	return Style(Text("grid-template-columns: " + strings.Join(size, " ") + ";"))
}

func (c *Datatable[T]) createFulltext() Node {
	return Fragment(
		form_ui.Autofocus(),
		hx.Get(c.Generate().Action("HandleSearch")),
		hx.Trigger("input delay:300ms"),
		hx.Swap(hx.SwapOuterHtml),
		hx.Target(hx.HashId(c.Props.Name)),
	)
}

func (c *Datatable[T]) createOrder(param mirage.Map) Node {
	return Fragment(
		hx.Get(c.Generate().Action("HandleOrder", param)),
		hx.Trigger("click"),
		hx.Swap(hx.SwapOuterHtml),
		hx.Target(hx.HashId(c.Props.Name)),
	)
}

func (c *Datatable[T]) createLoadMore() Node {
	param := mirage.Map{
		mystiq.Offset:   c.Param.Offset + mystiq.DefaultLimit,
		mystiq.Fulltext: c.Param.Fulltext,
		mystiq.Order:    c.Param.Order,
	}
	return Fragment(
		hx.Get(c.Generate().Action("HandleLoadMore", param)),
		hx.Target(hx.HashId(c.Props.Name+"-rows")),
		hx.Swap(hx.SwapBeforeEnd),
		hx.Trigger("intersect once"),
	)
}
