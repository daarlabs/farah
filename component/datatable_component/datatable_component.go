package datatable_component

import (
	"fmt"
	"slices"
	"strings"
	
	"golang.org/x/exp/maps"
	
	"github.com/daarlabs/farah/ui/button_ui"
	"github.com/daarlabs/farah/ui/drawer_ui"
	"github.com/daarlabs/farah/ui/search_ui"
	"github.com/daarlabs/farah/ui/tag_ui"
	"github.com/daarlabs/hirokit/dyna"
	"github.com/daarlabs/hirokit/esquel"
	
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/ui/box_ui"
	"github.com/daarlabs/farah/ui/form_ui"
	"github.com/daarlabs/farah/ui/spinner_ui"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/hirokit/hx"
	
	. "github.com/daarlabs/hirokit/gox"
)

type Datatable[T any] struct {
	hiro.Component
	Props         Props                               `json:"-"`
	Param         dyna.Param                          `json:"-"`
	Query         dyna.Query                          `json:"-"`
	FindDataFunc  func(param dyna.Param, t any) error `json:"-"`
	FieldsFunc    func() [][]Field                    `json:"-"`
	RowFunc       func(builder RowBuilder[T]) Node    `json:"-"`
	ActiveFilters map[string]string                   `json:"-"`
	FiltersFunc   func() Node                         `json:"-"`
	Data          []T                                 `json:"-"`
	
	fields     []Field
	headerRows [][]Field
}

var (
	zeroFiltersFunc = func() Node { return Fragment() }
)

func (c *Datatable[T]) Name() string {
	return "datatable"
}

func (c *Datatable[T]) Mount() {
	if len(c.Props.MinWidth) == 0 {
		c.Props.MinWidth = "1200px"
	}
	c.Param.Fields.Map = c.Query.Fields
	if c.FieldsFunc != nil {
		c.headerRows = c.FieldsFunc()
		c.fields = c.headerRows[len(c.headerRows)-1]
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

func (c *Datatable[T]) HandleFilter() error {
	c.Param = c.Param.Parse(c)
	c.Param.Offset = 0
	c.Data = c.getData()
	return c.Response().Render(c.createDatatable())
}

func (c *Datatable[T]) HandleSearch() error {
	c.Param = c.Param.Parse(c)
	c.Param.Offset = 0
	c.Data = c.getData()
	return c.Response().Render(c.createDatatable())
}

func (c *Datatable[T]) HandleAutocomplete() error {
	c.Param = c.Param.Parse(c)
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
	if len(param.Fields.Fulltext) == 0 {
		param.Fields.Fulltext = []string{c.Query.Alias + "." + esquel.Vectors}
	}
	if len(param.Fields.Map) == 0 {
		param.Fields.Map = c.Query.Fields
	}
	q := dyna.New()
	if c.FindDataFunc != nil {
		q = q.FindFunc(c.FindDataFunc)
	}
	if c.Query.CanUse() {
		q = q.DB(c.DB(), c.Query)
	}
	if len(c.Data) > 0 {
		q = q.Data(convertDataToMapSlice(c.Data))
	}
	q.MustFind(param, &result)
	return result
}

func (c *Datatable[T]) createDatatable() Node {
	titleExists := len(c.Props.Title) > 0
	searchExists := len(c.Param.Fields.Fulltext) > 0
	showHeadBar := searchExists || c.FiltersFunc != nil || c.FiltersFunc != nil && len(c.ActiveFilters) > 0
	rows := make([]string, 0)
	if titleExists {
		rows = append(rows, "1rem")
	}
	if showHeadBar {
		rows = append(rows, "2.5rem")
	}
	rows = append(rows, "1fr")
	return Div(
		tempest.Class().H("full").Grid().Gap(4).GridRows(strings.Join(rows, " ")),
		Id(hx.Id(c.Props.Name)),
		If(
			titleExists,
			Div(
				tempest.Class().TextSlate(900).TextWhite(tempest.Dark()).TextXs().Mr(4).FontSemibold(),
				Text(c.Props.Title),
			),
		),
		If(
			showHeadBar,
			Div(
				tempest.Class().Flex().ItemsCenter().Gap(2),
				c.createSearch(),
				c.createFilters(),
				c.createActiveFilters(),
			),
		),
		box_ui.Box(
			box_ui.Props{
				Class: tempest.Class().W("full").H("full"),
			},
			Div(
				tempest.Class().Grid().H("full").W("full").OverflowY("hidden").OverflowX("auto"),
				Div(
					tempest.Class().Grid().H("full").MinW(c.Props.MinWidth).Overflow("hidden"),
					Style(Attribute(), Raw(fmt.Sprintf("grid-template-rows: %drem 1fr", len(c.headerRows)*2))),
					c.createHead(),
					c.createBody(),
				),
			),
		),
	)
}

func (c *Datatable[T]) createSearchLabel() string {
	if !c.Config().Localization.Enabled {
		return "Search"
	}
	return c.Translate("component.datatable.search")
}

func (c *Datatable[T]) createHead() Node {
	return Div(
		Range(
			c.headerRows, func(row []Field, _ int) Node {
				return Div(
					tempest.Class().Transition().Grid().H(8).W("full").BorderB(1).
						BorderColor(palette.Primary, 400).BorderColor(palette.Primary, 200, tempest.Dark()).Pr(4),
					c.createSizeStyle(row),
					Range(
						row, func(field Field, _ int) Node {
							el := "a"
							if !field.Sortable {
								el = "div"
							}
							return CreateElement(el)(
								tempest.Class().Relative().Transition().Flex().ItemsCenter().Gap(1).H("full").Px(2).
									TextSize("10px").FontSemibold().TextSlate(900).TextWhite(tempest.Dark()).
									If(
										field.Border,
										tempest.Class().BorderR(1).BorderSlate(300).
											BorderSlate(600, tempest.Dark()),
									).
									If(
										field.Sortable,
										tempest.Class().BgSlate(100, tempest.Hover()).
											BgSlate(700, tempest.Dark(), tempest.Hover()).CursorPointer(),
									).
									If(field.AlignX == ui.Left, tempest.Class().JustifyStart()).
									If(field.AlignX == ui.Center, tempest.Class().JustifyCenter()).
									If(field.AlignX == ui.Right, tempest.Class().JustifyEnd()),
								Text(field.Title),
								If(
									field.Sortable,
									spinner_ui.Spinner(
										spinner_ui.Props{Overlay: true, Class: tempest.Class(spinner_ui.HxIndicator)},
									),
									c.createOrder(
										hiro.Map{
											dyna.Fulltext: c.Param.Fulltext, dyna.Order: c.createNextOrder(field.Name),
										},
									),
									If(
										slices.Contains(c.Param.Order, field.Name+":"+dyna.Asc),
										icon_ui.Icon(
											icon_ui.Props{
												Icon:  icon_ui.ChevronUp,
												Class: tempest.Class().FlexNone().TextSlate(900).TextWhite(tempest.Dark()),
												Size:  ui.Sm,
											},
										),
									),
									If(
										slices.Contains(c.Param.Order, field.Name+":"+dyna.Desc),
										icon_ui.Icon(
											icon_ui.Props{
												Icon:  icon_ui.ChevronDown,
												Class: tempest.Class().FlexNone().TextSlate(900).TextWhite(tempest.Dark()),
												Size:  ui.Sm,
											},
										),
									),
								),
							)
						},
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

func (c *Datatable[T]) createActiveFilters() Node {
	var activeFilters = make(map[string]string)
	var fields []Field
	if c.FieldsFunc != nil {
		c.headerRows = c.FieldsFunc()
		c.fields = c.headerRows[len(c.headerRows)-1]
		fields = c.fields
	}
	for fieldName, fieldText := range c.ActiveFilters {
		var value, fieldTitle string
		for _, item := range fields {
			if item.Name == fieldText {
				fieldTitle = item.Title
				break
			}
		}
		for _, item := range c.Param.Filter {
			if strings.HasPrefix(item, fieldName+":") {
				value = strings.TrimPrefix(item, fieldName+":")
			}
		}
		if len(value) == 0 || value == "0" {
			continue
		}
		t := make(hiro.Map)
		if c.FindDataFunc == nil {
			q := dyna.New()
			if c.Query.CanUse() {
				q = q.DB(c.DB(), c.Query)
			}
			q.MustFindOne(fieldName, value, &t)
		}
		if c.FindDataFunc != nil {
			if err := c.FindDataFunc(
				dyna.Param{
					Fields: dyna.Fields{Map: c.Param.Fields.Map},
					Limit:  1,
					Filter: []string{fieldName + ":" + value},
				}, &t,
			); err != nil {
				panic(err)
			}
		}
		text, ok := t[fieldText]
		if ok {
			var activeFilterText string
			switch v := text.(type) {
			case []byte:
				activeFilterText = string(v)
			default:
				activeFilterText = fmt.Sprint(text)
			}
			if activeFilterText == "on" || activeFilterText == "true" {
				activeFilterText = c.createBoolTitle()
			}
			activeFilters[fieldName] = fieldTitle + ": " + strings.TrimSpace(activeFilterText)
		}
	}
	if len(c.Param.Filter) == 0 {
		return Fragment()
	}
	return Div(
		tempest.Class().Grid().GridRows(2).H("full"),
		Div(
			tempest.Class().TextXs().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()),
			Text(c.createActiveFiltersTitle()+":"),
		),
		Div(
			tempest.Class().Flex().FlexWrap().Gap(1),
			MapRange(
				activeFilters, func(key string, text string) Node {
					param := hiro.Map{dyna.Order: c.Param.Order, dyna.Fulltext: c.Param.Fulltext}
					for _, item := range c.Param.Filter {
						if !strings.Contains(item, ":") {
							continue
						}
						parts := strings.Split(item, ":")
						if key == parts[0] {
							continue
						}
						param[parts[0]] = parts[1]
					}
					return tag_ui.Tag(
						tag_ui.Props{},
						Text(text),
						Button(
							Type("button"),
							c.On("click").
								Action("HandleFilter", param).
								Replace(c.Props.Name),
							icon_ui.Icon(icon_ui.Props{Icon: icon_ui.Close, Size: ui.Sm, Class: tempest.Class().TextWhite()}),
						),
					)
				},
			),
		),
	)
}

func (c *Datatable[T]) createSearch() Node {
	if len(c.Param.Fields.Fulltext) == 0 {
		return Fragment()
	}
	return Div(
		tempest.Class().MaxW("300px"),
		search_ui.Search(
			search_ui.Props{
				Placeholder: c.createSearchLabel(), Value: c.Param.Fulltext, Name: dyna.Fulltext,
				Id: hx.Id(c.Props.Name + "-" + dyna.Fulltext),
			},
			c.createFulltext(hiro.Map{dyna.Order: c.Param.Order}),
		),
	)
}

func (c *Datatable[T]) createFilters() Node {
	if c.FiltersFunc == nil {
		return Fragment()
	}
	filterFields := maps.Keys(c.Query.Fields)
	include := make([]string, len(filterFields))
	for i, field := range filterFields {
		include[i] = hx.HashId(c.Props.Name) + " [name=" + field + "]"
	}
	param := hiro.Map{
		dyna.Fulltext: c.Param.Fulltext,
		dyna.Order:    c.Param.Order,
	}
	createFilters := zeroFiltersFunc
	if c.FiltersFunc != nil {
		createFilters = c.FiltersFunc
	}
	return drawer_ui.Container(
		button_ui.IconButton(
			button_ui.Props{
				Icon: icon_ui.Filter,
			},
			drawer_ui.OpenEvent(),
		),
		drawer_ui.Drawer(
			Text(c.createFiltersTitle()),
			Div(
				tempest.Class().Grid().GridCols(2).Gap(2).ItemsCenter().H("full"),
				button_ui.RedButton(
					button_ui.Props{Size: ui.Sm, Icon: icon_ui.Refresh},
					hx.Get(c.Generate().Action("HandleFilter", param)),
					hx.Swap(hx.SwapOuterHtml),
					hx.Trigger("click"),
					hx.Target(hx.HashId(c.Props.Name)),
					Text(c.createFiltersResetLabel()),
				),
				button_ui.EmeraldButton(
					button_ui.Props{Size: ui.Sm, Icon: icon_ui.Check},
					hx.Get(c.Generate().Action("HandleFilter", param)),
					hx.Swap(hx.SwapOuterHtml),
					hx.Trigger("click"),
					hx.Target(hx.HashId(c.Props.Name)),
					hx.Include(strings.Join(include, ",")),
					Text(c.createFiltersApplyLabel()),
				),
			),
			createFilters(),
		),
	)
}

func (c *Datatable[T]) createRows() Node {
	fields := c.headerRows[len(c.headerRows)-1]
	sizeStyle := c.createSizeStyle(fields)
	return If(
		len(c.Data) > 0 && c.RowFunc != nil,
		Range(
			c.Data,
			func(item T, index int) Node {
				var loadMore Node
				if (index+1)%dyna.DefaultLimit == 0 {
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
		return append(c.Param.Order, name+":"+dyna.Asc)
	}
	for _, o := range c.Param.Order {
		isName := strings.HasPrefix(o, name+":")
		if !isName {
			result = append(result, o)
			continue
		}
		if strings.HasSuffix(o, dyna.Asc) {
			result = append(result, strings.Replace(o, dyna.Asc, dyna.Desc, 1))
		}
		continue
	}
	return result
}

func (c *Datatable[T]) createSizeStyle(fields []Field) Node {
	n := len(c.fields)
	if n == 0 {
		return Fragment()
	}
	size := make([]string, n)
	for i, f := range fields {
		size[i] = f.Size
	}
	return Style(Text("grid-template-columns: " + strings.Join(size, " ") + ";"))
}

func (c *Datatable[T]) createFulltext(param hiro.Map) Node {
	param = param.Merge(c.Param.FilterMap)
	return Fragment(
		form_ui.Autofocus(),
		hx.Get(c.Generate().Action("HandleSearch", param)),
		hx.Trigger("change"),
		hx.Swap(hx.SwapOuterHtml),
		hx.Target(hx.HashId(c.Props.Name)),
	)
}

func (c *Datatable[T]) createOrder(param hiro.Map) Node {
	param = param.Merge(c.Param.FilterMap)
	return Fragment(
		hx.Get(c.Generate().Action("HandleOrder", param)),
		hx.Trigger("click"),
		hx.Swap(hx.SwapOuterHtml),
		hx.Target(hx.HashId(c.Props.Name)),
	)
}

func (c *Datatable[T]) createLoadMore() Node {
	param := hiro.Map{
		dyna.Offset:   c.Param.Offset + dyna.DefaultLimit,
		dyna.Fulltext: c.Param.Fulltext,
		dyna.Order:    c.Param.Order,
	}.Merge(c.Param.FilterMap)
	return Fragment(
		hx.Get(c.Generate().Action("HandleLoadMore", param)),
		hx.Target(hx.HashId(c.Props.Name+"-rows")),
		hx.Swap(hx.SwapBeforeEnd),
		hx.Trigger("intersect once"),
	)
}

func (c *Datatable[T]) createActiveFiltersTitle() string {
	if !c.Config().Localization.Enabled {
		return "Active filters"
	}
	return c.Translate("component.datatable.filters.active.title")
}

func (c *Datatable[T]) createFiltersTitle() string {
	if !c.Config().Localization.Enabled {
		return "Filtration"
	}
	return c.Translate("component.datatable.filters.title")
}

func (c *Datatable[T]) createFiltersApplyLabel() string {
	if !c.Config().Localization.Enabled {
		return "Apply"
	}
	return c.Translate("component.datatable.filters.apply")
}

func (c *Datatable[T]) createFiltersResetLabel() string {
	if !c.Config().Localization.Enabled {
		return "Reset"
	}
	return c.Translate("component.datatable.filters.reset")
}

func (c *Datatable[T]) createBoolTitle() string {
	if !c.Config().Localization.Enabled {
		return "Yes"
	}
	return c.Translate("data.yes")
}
