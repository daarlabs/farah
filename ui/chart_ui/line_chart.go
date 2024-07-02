package chart_ui

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
)

func LineChart(props Props) Node {
	props.Config = mergeConfig(
		mirage.Map{
			"chart": mirage.Map{
				"type":   "line",
				"height": props.Height,
			},
		},
		createConfig(props),
	)
	return Chart(props)
}
