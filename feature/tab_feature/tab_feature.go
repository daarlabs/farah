package tab_feature

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
)

type TabFeature struct {
	mirage.Component
	Props Props `json:"-"`
}

func (c *TabFeature) Name() string {
	return c.Props.Name + "-tabs"
}

func (c *TabFeature) Mount() {
	c.Parse().MustQuery("tab", &c.Props.Active)
}

func (c *TabFeature) Node() Node {
	return Div(
		Div(
			Class("flex items-center gap-2"),
			Range(
				c.Props.Tabs, func(tab Tab, _ int) Node {
					return Div(
						A(
							Clsx{
								"block transition text-xs px-4 py-2 rounded":                                                  true,
								"bg-transparent text-slate-900 dark:text-white border border-slate-300 dark:border-slate-600": tab.Name != c.Props.Active,
								"text-white bg-primary-400 dark:bg-primary-200":                                               c.Props.Active != "" && tab.Name == c.Props.Active,
							},
							Href(c.Request().Path()+c.Generate().Query(mirage.Map{"tab": tab.Name})),
							Text(tab.Title),
						),
					)
				},
			),
		),
		Div(
			Range(
				c.Props.Tabs, func(tab Tab, _ int) Node {
					if tab.Name != c.Props.Active || tab.NodeFunc == nil {
						return Fragment()
					}
					return tab.NodeFunc()
				},
			),
		),
	)
}
