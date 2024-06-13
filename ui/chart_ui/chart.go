package chart_ui

import (
	"encoding/json"
	"strings"
	
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
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

func mergeConfig(c1, c2 mirage.Map) mirage.Map {
	for k1, v1 := range c1 {
		if v1, ok := v1.(mirage.Map); ok {
			if v2, ok := c2[k1]; ok {
				if v2, ok := v2.(mirage.Map); ok {
					c2[k1] = mergeConfig(v1, v2)
					continue
				}
			}
		}
		c2[k1] = v1
	}
	return c2
}

func createConfig(props Props) mirage.Map {
	gridColor := tempest.Pallete[tempest.Slate][200]
	legendColor := tempest.Pallete[tempest.Slate][900]
	primaryColor := palette.PrimaryPallete[400]
	if props.Dark {
		gridColor = tempest.Pallete[tempest.Slate][400]
		legendColor = "#FFFFFF"
		primaryColor = palette.PrimaryPallete[200]
	}
	return mirage.Map{
		"chart": mirage.Map{
			"toolbar": mirage.Map{
				"show": false,
			},
		},
		"grid": mirage.Map{
			"borderColor": gridColor,
		},
		"series": mirage.Slice{
			{
				"name":  props.Id,
				"color": primaryColor,
				"data":  props.DataY,
			},
		},
		"stroke": mirage.Map{
			"colors": []string{primaryColor},
		},
		"markers": mirage.Map{
			"colors": []string{primaryColor},
		},
		"xaxis": mirage.Map{
			"categories": props.DataX,
			"labels": mirage.Map{
				"style": mirage.Map{
					"colors": legendColor,
				},
			},
		},
		"yaxis": mirage.Map{
			"labels": mirage.Map{
				"style": mirage.Map{
					"colors": legendColor,
				},
			},
		},
	}
}
