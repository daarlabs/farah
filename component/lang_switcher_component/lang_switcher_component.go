package lang_switcher_component

import (
	"embed"
	
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/menu_ui"
	"github.com/daarlabs/farah/ui/menu_ui/menu_item_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
)

type LangSwitcher struct {
	hiro.Component
}

//go:embed flag/*.svg
var flags embed.FS

func (c *LangSwitcher) Name() string {
	return "lang-switcher"
}

func (c *LangSwitcher) Mount() {
}

func (c *LangSwitcher) Node() Node {
	currentLang := c.Lang().Current()
	langs := c.Config().Localization.Languages
	langsNodes := make([]Node, len(langs))
	for i, lang := range langs {
		langsNodes[i] = A(
			Href(c.Generate().Current(hiro.Map{"lang": lang.Code})),
			menu_item_ui.MenuItem(
				menu_item_ui.Props{
					Selected: lang.Code == currentLang,
				},
				Text(c.Translate("lang."+lang.Code)),
			),
		)
	}
	return menu_ui.Menu(
		menu_ui.Props{
			Clickable:  true,
			Autoheight: true,
			PositionX:  ui.Right,
		},
		Button(
			Type("button"),
			menu_ui.Open(),
			tempest.Class().Block().Transition().W("16px").H("11px").My("auto").
				Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()),
			c.getFlag(currentLang),
		),
		langsNodes...,
	)
}

func (c *LangSwitcher) getFlag(lang string) Node {
	flagBytes, err := flags.ReadFile("flag/" + lang + ".svg")
	if err != nil {
		return Text(lang)
	}
	return Raw(string(flagBytes))
}
