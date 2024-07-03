package datepicker_component

import (
	"fmt"
	"time"
	
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/tempest/form_tempest"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	. "github.com/daarlabs/hirokit/gox"
	
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
	"github.com/daarlabs/farah/ui/form_ui/hidden_field_ui"
	"github.com/daarlabs/hirokit/hx"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/menu_ui"
)

type Datepicker struct {
	hiro.Component
	Props Props `json:"-"`
	
	currentDay       int
	lastDay          int
	firstDayPosition int
}

const (
	dayKey   = "day"
	monthKey = "month"
	yearKey  = "year"
)

var (
	months = createMonths()
	years  = createYears()
)

func (c *Datepicker) Name() string {
	return c.Props.Name + "-datepicker"
}

func (c *Datepicker) Mount() {
	isAction := c.Request().Is().Action()
	if !isAction && c.Props.Value.IsZero() {
		c.Props.Value = time.Now()
	}
	if isAction {
		day, month, year := c.getDateFromAction()
		lastDay := c.getLastDayOfMonth(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, c.Props.Value.Location()))
		if day > lastDay {
			day = lastDay
		}
		c.Props.Value = time.Date(year, time.Month(month), day, 0, 0, 0, 0, c.Props.Value.Location())
	}
	c.currentDay = c.Props.Value.Day()
	c.firstDayPosition = c.getFirstDayOfMonthPosition(c.Props.Value)
	c.lastDay = c.getLastDayOfMonth(c.Props.Value)
}

func (c *Datepicker) Node() Node {
	return c.createDatepicker(false)
}

func (c *Datepicker) HandleSelectDay() error {
	return c.Response().Render(c.createDatepicker(false))
}

func (c *Datepicker) HandleSelectParameter() error {
	return c.Response().Render(c.createDatepicker(true))
}

func (c *Datepicker) createDatepicker(open bool) Node {
	return menu_ui.Menu(
		menu_ui.Props{
			Id:         hx.Id(c.Props.Id),
			Clickable:  true,
			PositionX:  ui.Left,
			Open:       open,
			Fullwidth:  true,
			Autoheight: true,
		},
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
				tempest.Class().Relative(),
				Button(
					tempest.Class().Absolute().Left(3).InsetY(0).My("auto").Size(4),
					Type("button"),
					icon_ui.Icon(
						icon_ui.Props{
							Icon:  icon_ui.Calendar,
							Size:  ui.Sm,
							Class: tempest.Class().TextSlate(900).TextWhite(tempest.Dark()),
						},
					),
					menu_ui.Open(),
				),
				Button(
					If(len(c.Props.Id) > 0, Id(c.Props.Id)),
					Type("button"),
					tempest.Class().
						H(10).
						Transition().W("full").Pr(3).Pl(10).Rounded().
						BgWhite().BgSlate(800, tempest.Dark()).
						// Font
						TextSize(tempest.SizeXs).TextSlate(900).TextWhite(tempest.Dark()).TextLeft().
						// Border
						Border(1).
						BorderColor(tempest.Slate, 300).
						BorderColor(tempest.Slate, 600, tempest.Dark()).
						BorderColor(palette.Primary, 400, tempest.Focus()).
						BorderColor(
							palette.Primary, 200, tempest.Focus(), tempest.Dark(),
						).
						Extend(form_tempest.FocusShadow()),
					menu_ui.Open(),
					Text(c.createButtonText()),
				),
			),
			hidden_field_ui.HiddenField(c.Props.Name, c.Props.Value.UTC().String()),
		),
		c.createHead(),
		c.createBody(),
	)
}

func (c *Datepicker) createHead() Node {
	return Fragment(
		Div(
			tempest.Class().Transition().Grid().GridCols(2).
				BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
			c.createMonthPicker(),
			c.createYearPicker(),
		),
		Div(
			tempest.Class().Transition().Grid().GridCols(7).H(8).
				BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
			c.createHeadDay(c.getMondayShortLabelText()),
			c.createHeadDay(c.getTuesdayShortLabelText()),
			c.createHeadDay(c.getWednesdayShortLabelText()),
			c.createHeadDay(c.getThursdayShortLabelText()),
			c.createHeadDay(c.getFridayShortLabelText()),
			c.createHeadDay(c.getSaturdayShortLabelText()),
			c.createHeadDay(c.getSundayShortLabelText()),
		),
	)
}

func (c *Datepicker) createHeadDay(label string) Node {
	return Div(
		tempest.Class().Transition().TextXs().Grid().PlaceItemsCenter().H(8).
			BorderR(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
			FontSemibold().TextSlate(900).TextWhite(tempest.Dark()),
		Text(label),
	)
}

func (c *Datepicker) createMonthPicker() Node {
	year, month, day := c.Props.Value.Date()
	selectedMonth := int(month)
	return Div(
		tempest.Class().Transition().P(2).BorderR(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
		field_label_ui.FieldLabel(
			field_label_ui.Props{
				For:  c.Props.Id + "-month",
				Text: c.getMonthLabelText(),
			},
		),
		Select(
			If(len(c.Props.Id) > 0, Id(c.Props.Id+"-month")),
			tempest.Class().Transition().W("full").Flex().Rounded().
				TextXs().FontSemibold().Text(palette.Primary, 400).Text(palette.Primary, 100, tempest.Dark()).
				BgWhite().BgSlate(800, tempest.Dark()).
				Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
			Name(monthKey),
			hx.Get(c.Generate().Action("HandleSelectParameter", hiro.Map{dayKey: day, yearKey: year})),
			hx.Trigger("change"),
			hx.Swap(hx.SwapOuterHtml),
			hx.Target(hx.HashId(c.Props.Id)),
			Range(
				months, func(i int, _ int) Node {
					return Option(Value(i), Text(i), If(selectedMonth == i, Selected()))
				},
			),
		),
	)
}

func (c *Datepicker) createYearPicker() Node {
	selectedYear, month, day := c.Props.Value.Date()
	return Div(
		tempest.Class().Transition().P(2).BorderR(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
		field_label_ui.FieldLabel(
			field_label_ui.Props{
				For:  c.Props.Id + "-year",
				Text: c.getYearLabelText(),
			},
		),
		Select(
			If(len(c.Props.Id) > 0, Id(c.Props.Id+"-year")),
			tempest.Class().Transition().W("full").Flex().Rounded().
				TextXs().FontSemibold().Text(palette.Primary, 400).Text(palette.Primary, 100, tempest.Dark()).
				BgWhite().BgSlate(800, tempest.Dark()).
				Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
			Name(yearKey),
			hx.Get(c.Generate().Action("HandleSelectParameter", hiro.Map{dayKey: day, monthKey: int(month)})),
			hx.Trigger("change"),
			hx.Swap(hx.SwapOuterHtml),
			hx.Target(hx.HashId(c.Props.Id)),
			Range(
				years, func(i int, _ int) Node {
					return Option(Value(i), Text(i), If(selectedYear == i, Selected()))
				},
			),
		),
	)
}

func (c *Datepicker) createBody() Node {
	return Div(
		tempest.Class().Grid().GridCols(7).GridRows(5),
		c.createBodyDays(),
	)
}

func (c *Datepicker) createBodyDays() Node {
	year, month, _ := c.Props.Value.Date()
	currentDay := c.currentDay - 1
	firstDayPosition := c.firstDayPosition - 1
	days := make([]Node, c.lastDay+firstDayPosition)
	if firstDayPosition > 0 {
		for i := 0; i < firstDayPosition; i++ {
			days[i] = Div(
				tempest.Class().Transition().BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).H(8).
					If(firstDayPosition-1 == i, tempest.Class().BorderR(1)),
			)
		}
	}
	for i := firstDayPosition; i < c.lastDay+firstDayPosition; i++ {
		isActive := i-firstDayPosition == currentDay
		el := "a"
		if isActive {
			el = "div"
		}
		days[i] = CreateElement(el)(
			tempest.Class().Transition().Grid().PlaceItemsCenter().TextXs().H(8).
				BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
				If((i+1)%7 != 0, tempest.Class().BorderR(1)).
				If(
					!isActive,
					tempest.Class().CursorPointer().
						BgSlate(100, tempest.Hover()).BgSlate(700, tempest.Dark(), tempest.Hover()).
						TextSlate(900).TextWhite(tempest.Dark()),
				).
				If(
					isActive,
					tempest.Class().Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()).
						TextWhite().FontSemibold(),
				),
			Text(i+1-firstDayPosition),
			If(
				!isActive,
				hx.Get(
					c.Generate().Action(
						"HandleSelectDay", hiro.Map{dayKey: i + 1 - firstDayPosition, monthKey: int(month), yearKey: year},
					),
				),
				menu_ui.Close(),
			),
			hx.Trigger("click"),
			hx.Swap(hx.SwapOuterHtml),
			hx.Target(hx.HashId(c.Props.Id)),
		)
	}
	return Fragment(days...)
}

func (c *Datepicker) createButtonText() string {
	year, month, day := c.Props.Value.Date()
	return fmt.Sprintf(`%d. %d. %d`, day, int(month), year)
}

func (c *Datepicker) getLastDayOfMonth(t time.Time) int {
	return time.Date(
		t.Year(),
		t.Month(), 1, 0, 0, 0, 0,
		t.Location(),
	).AddDate(0, 1, -1).Day()
}

func (c *Datepicker) getFirstDayOfMonthPosition(t time.Time) int {
	pos := int(time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Weekday())
	if pos == 0 {
		return 7
	}
	return pos
}

func (c *Datepicker) getDateFromAction() (int, int, int) {
	var day, month, year int
	c.Parse().MustQuery(dayKey, &day)
	c.Parse().MustQuery(monthKey, &month)
	c.Parse().MustQuery(yearKey, &year)
	return day, month, year
}

func (c *Datepicker) getMonthLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Month"
	}
	return c.Translate("component.datepicker.month")
}

func (c *Datepicker) getYearLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Year"
	}
	return c.Translate("component.datepicker.year")
}

func (c *Datepicker) getMondayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Mo"
	}
	return c.Translate("component.datepicker.day.short.monday")
}

func (c *Datepicker) getTuesdayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Tu"
	}
	return c.Translate("component.datepicker.day.short.tuesday")
}

func (c *Datepicker) getWednesdayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "We"
	}
	return c.Translate("component.datepicker.day.short.wednesday")
}

func (c *Datepicker) getThursdayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Th"
	}
	return c.Translate("component.datepicker.day.short.thursday")
}

func (c *Datepicker) getFridayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Fr"
	}
	return c.Translate("component.datepicker.day.short.friday")
}

func (c *Datepicker) getSaturdayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Sa"
	}
	return c.Translate("component.datepicker.day.short.saturday")
}

func (c *Datepicker) getSundayShortLabelText() string {
	if !c.Config().Localization.Enabled {
		return "Su"
	}
	return c.Translate("component.datepicker.day.short.sunday")
}
