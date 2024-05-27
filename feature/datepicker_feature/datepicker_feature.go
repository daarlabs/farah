package datepicker_feature

import (
	"fmt"
	"time"
	
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/stimulus"
	
	. "github.com/daarlabs/arcanum/gox"
	
	"component/ui/form_ui/field_label_ui"
	"component/ui/form_ui/hidden_field_ui"
	"github.com/daarlabs/arcanum/hx"
	
	"component/ui"
	"component/ui/menu_ui"
)

type Datepicker struct {
	mirage.Component
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
			Button(
				If(len(c.Props.Id) > 0, Id(c.Props.Id)),
				Type("button"),
				Clsx{
					"transition w-full border pl-3 pr-7 rounded focus:shadow-focus text-left h-10 text-xs":                                      true,
					"bg-white dark:bg-slate-800 text-slate-900 dark:text-white border-slate-300 dark:border-slate-600 focus:border-primary-400": true,
					"text-slate-800 dark:text-white": true,
				},
				stimulus.Action("click", "menu", "handleOpen"),
				Text(c.createButtonText()),
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
			Class("grid grid-cols-2"),
			c.createMonthPicker(),
			c.createYearPicker(),
		),
		Div(
			Class("transition grid grid-cols-7 border-b border-t border-slate-300 dark:border-slate-600 h-8"),
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
		Class("transition text-xs grid place-items-center h-8 border-r border-slate-300 dark:border-slate-600 font-semibold text-slate-900 dark:text-white"),
		Text(label),
	)
}

func (c *Datepicker) createMonthPicker() Node {
	year, month, day := c.Props.Value.Date()
	selectedMonth := int(month)
	return Div(
		Class("transition p-2 border-r border-slate-300 dark:border-slate-600"),
		field_label_ui.FieldLabel(
			field_label_ui.Props{
				For:  c.Props.Id + "-month",
				Text: c.getMonthLabelText(),
			},
		),
		Select(
			If(len(c.Props.Id) > 0, Id(c.Props.Id+"-month")),
			Class("transition flex w-full text-xs font-semibold text-primary-400 dark:text-primary-100 bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 dark:border-slate-600 rounded"),
			Name(monthKey),
			hx.Get(c.Generate().Action("HandleSelectParameter", mirage.Map{dayKey: day, yearKey: year})),
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
		Class("transition p-2 border-r border-slate-300 dark:border-slate-600"),
		field_label_ui.FieldLabel(
			field_label_ui.Props{
				For:  c.Props.Id + "-year",
				Text: c.getYearLabelText(),
			},
		),
		Select(
			If(len(c.Props.Id) > 0, Id(c.Props.Id+"-year")),
			Class("transition flex w-full text-xs font-semibold text-primary-400 dark:text-primary-100 bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-600 rounded"),
			Name(yearKey),
			hx.Get(c.Generate().Action("HandleSelectParameter", mirage.Map{dayKey: day, monthKey: int(month)})),
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
		Class("grid grid-cols-7 grid-rows-5"),
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
				Clsx{
					"transition border-b border-slate-300 dark:border-slate-600 h-8": true,
					"border-r": firstDayPosition-1 == i,
				},
				Class("transition border-b border-slate-300 dark:border-slate-600 h-8"),
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
			Clsx{
				"transition grid place-items-center text-xs border-b border-slate-300 dark:border-slate-600 h-8": true,
				// "border-l": firstDayPosition != 0 && i-firstDayPosition == 0,
				"border-r": (i+1)%7 != 0,
				"transition hover:bg-slate-100 dark:hover:bg-slate-700 cursor-pointer text-slate-900 dark:text-white": !isActive,
				"bg-primary-400 dark:bg-primary-200 text-white font-semibold":                                         isActive,
			},
			Text(i+1-firstDayPosition),
			If(
				!isActive,
				hx.Get(
					c.Generate().Action(
						"HandleSelectDay", mirage.Map{dayKey: i + 1 - firstDayPosition, monthKey: int(month), yearKey: year},
					),
				),
				stimulus.Action("click", "menu", "handleClose"),
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
