package chart_ui

import (
	"encoding/json"
	"strings"
	
	"github.com/daarlabs/farah/palette"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
)

func Chart(props Props) Node {
	cfg, err := json.Marshal(props.Config)
	if err != nil {
		return Div(
			Text(err),
		)
	}
	script := make([]string, 0)
	if !props.Defer {
		script = append(script, `document.addEventListener('DOMContentLoaded', function () {`)
	}
	script = append(script, `var chart = new ApexCharts(document.querySelector("#`+props.Id+`"), `+string(cfg)+`);`)
	script = append(script, `chart.render();`)
	if !props.Defer {
		script = append(script, `});`)
	}
	return Div(
		Div(
			Id(props.Id),
		),
		Script(
			Raw(strings.Join(script, "\n")),
		),
	)
}

func mergeConfig(c1, c2 hiro.Map) hiro.Map {
	for k1, v1 := range c1 {
		if v1, ok := v1.(hiro.Map); ok {
			if v2, ok := c2[k1]; ok {
				if v2, ok := v2.(hiro.Map); ok {
					c2[k1] = mergeConfig(v1, v2)
					continue
				}
			}
		}
		c2[k1] = v1
	}
	return c2
}

func createConfig(props Props) hiro.Map {
	gridColor := tempest.Pallete[tempest.Slate][200]
	legendColor := tempest.Pallete[tempest.Slate][900]
	primaryColor := palette.PrimaryPallete[400]
	if props.Dark {
		gridColor = tempest.Pallete[tempest.Slate][400]
		legendColor = "#FFFFFF"
		primaryColor = palette.PrimaryPallete[200]
	}
	return hiro.Map{
		"chart": hiro.Map{
			"toolbar": hiro.Map{
				"show": false,
			},
		},
		"grid": hiro.Map{
			"borderColor": gridColor,
		},
		"series": hiro.Slice{
			{
				"name":  props.Id,
				"color": primaryColor,
				"data":  props.DataY,
			},
		},
		"stroke": hiro.Map{
			"colors": []string{primaryColor},
		},
		"markers": hiro.Map{
			"colors": []string{primaryColor},
		},
		"xaxis": hiro.Map{
			"categories": props.DataX,
			"labels": hiro.Map{
				"style": hiro.Map{
					"colors": legendColor,
				},
			},
		},
		"yaxis": hiro.Map{
			"labels": hiro.Map{
				"style": hiro.Map{
					"colors": legendColor,
				},
			},
		},
	}
}
