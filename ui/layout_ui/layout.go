package layout_ui

import (
	"github.com/daarlabs/farah/component/flash_message_component"
	"github.com/daarlabs/hirokit/devtool"
	"github.com/daarlabs/hirokit/env"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
)

func Layout(c hiro.Ctx, nodes ...Node) Node {
	var currentLang string
	if c.Config().Localization.Enabled {
		currentLang = c.Lang().Current()
	}
	title := c.Page().Get().Title()
	c.Page().Set().Meta(
		"viewport", "width=device-width, initial-scale=1",
	)
	config := c.Config()
	return Html(
		If(len(currentLang) > 0, Lang(currentLang)),
		Head(
			If(len(title) > 0, Title(Text(title))),
			If(len(title) == 0, Title(Text(c.Config().App.Name))),
			Meta(CharSet("utf-8")),
			If(!config.App.Robots, Meta(Name("robots"), Content("noindex, nofollow"))),
			Raw(
				`
				<link rel="apple-touch-icon" sizes="180x180" href="/public/favicon/apple-touch-icon.png">
				<link rel="icon" type="image/png" sizes="32x32" href="/public/favicon/favicon-32x32.png">
				<link rel="icon" type="image/png" sizes="16x16" href="/public/favicon/favicon-16x16.png">
				<link rel="manifest" href="/public/favicon/site.webmanifest">
				<link rel="mask-icon" href="/public/favicon/safari-pinned-tab.svg" color="#5bbad5">
				<link rel="shortcut icon" href="/public/favicon/favicon.ico">
				<meta name="msapplication-TileColor" content="#00aba9">
				<meta name="msapplication-config" content="/public/favicon/browserconfig.xml">
				<meta name="theme-color" content="#ffffff">
			`,
			),
			Range(
				c.Page().Get().Metas(), func(item [2]string, _ int) Node {
					return Meta(Name(item[0]), Content(item[1]))
				},
			),
			c.Generate().Assets(c.Request().Name()),
		),
		Body(
			Clsx{
				tempest.Class().Dark(): Dark(c),
			},
			c.Create().Component(&flash_message_component.FlashMessage{}),
			Div(
				tempest.Class().Transition().Bg(tempest.Slate, 100).Bg(tempest.Slate, 900, tempest.Dark()).
					Overflow("hidden").W("screen").H("screen"),
				Fragment(nodes...),
			),
			If(
				env.Development() && c.Config().Dev.Tool,
				devtool.CreateNode(),
			),
		),
	)
}
