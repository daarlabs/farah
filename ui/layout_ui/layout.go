package layout_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	
	"component/feature/dark_mode_switcher_feature"
)

func Layout(c mirage.Ctx, nodes ...Node) Node {
	title := c.Page().Get().Title()
	return Html(
		Head(
			If(len(title) > 0, Title(Text(title))),
			If(len(title) == 0, Title(Text("SWTP"))),
			Range(
				c.Page().Get().Metas(), func(item [2]string, _ int) Node {
					return Meta(Name(item[0]), Content(item[1]))
				},
			),
			Link(Rel("preconnect"), Href("https://fonts.googleapis.com")),
			Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), CrossOrigin()),
			Link(Rel("stylesheet"), Href("https://fonts.googleapis.com/css2?family=Sora:wght@100..800&display=swap")),
			c.Generate().Assets(),
		),
		Body(
			Clsx{
				"font-sora": true,
				"dark":      c.Cookie().Get(dark_mode_switcher_feature.DarkModeCookieKey) == "true",
			},
			Div(
				Class("transition bg-slate-100 dark:bg-slate-900 overflow-hidden w-screen h-screen"),
				Fragment(nodes...),
			),
		),
	)
}
