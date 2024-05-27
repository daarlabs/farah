package main

import (
	"github.com/daarlabs/arcanum/cache/memory"
	"github.com/daarlabs/arcanum/mirage"
	
	"github.com/daarlabs/arcanum/config"
	
	"component/web/handler/feature_handler"
	"component/web/handler/home_handler"
	"component/web/handler/ui_handler"
	"component/web/ui/web_ui/web_layout_ui"
)

func main() {
	cfg := config.Config{
		App: config.App{
			Name:   "farah-showcase",
			Public: "/public/",
			Assets: "web/public/dist/",
		},
		Cache: config.Cache{Memory: memory.New(".cache")},
		Router: config.Router{
			Recover: true,
		},
	}
	app := mirage.New(cfg)
	app.Layout().Add(mirage.Main, web_layout_ui.Layout)
	app.Static("/public/", "./web/public")
	app.Route("/", home_handler.Get(), mirage.Name("home"))
	app.Route("/ui", ui_handler.Get(), mirage.Name("ui"))
	app.Route("/feature", feature_handler.Get(), mirage.Name("feature"))
	app.Run(":8000")
}
