package tab_component

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/hirokit/hiro"
)

type TabFeature struct {
	hiro.Component
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
			tempest.Class().Flex().ItemsCenter().Gap(2),
			Range(
				c.Props.Tabs, func(tab Tab, _ int) Node {
					return Div(
						A(
							tempest.Class().Block().Transition().TextXs().Px(4).Py(2).Rounded().
								If(
									tab.Name != c.Props.Active,
									tempest.Class().BgTransparent().TextSlate(900).TextWhite(tempest.Dark()).
										Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
								).
								If(
									c.Props.Active != "" && tab.Name == c.Props.Active,
									tempest.Class().Bg(palette.Primary, 400).Bg(palette.Primary, 200, tempest.Dark()),
								),
							Href(c.Request().Path()+c.Generate().Query(hiro.Map{"tab": tab.Name})),
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
