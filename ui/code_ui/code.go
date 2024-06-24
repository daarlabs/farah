package code_ui

import (
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/alpine"
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/farah/ui/icon_ui"
)

func Code(props Props) gox.Node {
	id := "copy-" + uniuri.New()
	return gox.Div(
		alpine.Data(mirage.Map{}),
		alpine.Init("new Clipboard('#"+id+"');"),
		tempest.Class().Relative().TextXs().TextSlate(900).TextWhite(tempest.Dark()).
			Py(4).Pl(4).Pr(8).Rounded().
			BgWhite().BgSlate(700, tempest.Dark()).
			Border(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
			BreakAll(),
		gox.Button(
			gox.Id(id),
			gox.Type("button"),
			tempest.Class().Absolute().Top(2).Right(2),
			gox.CustomData("clipboard-text", props.Value),
			icon_ui.Icon(
				icon_ui.Props{
					Icon: icon_ui.Copy,
					Size: ui.Sm,
					Class: tempest.Class().TextSlate(600).TextSlate(300, tempest.Dark()).
						Text(palette.Primary, 400, tempest.Hover()).
						Text(palette.Primary, 100, tempest.Dark(), tempest.Hover()),
				},
			),
		),
		gox.Text(props.Value),
	)
}
