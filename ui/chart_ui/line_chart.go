package chart_ui

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
)

func LineChart(props Props) Node {
	props.Config = mergeConfig(
		hiro.Map{
			"chart": hiro.Map{
				"type":   "line",
				"height": props.Height,
			},
		},
		createConfig(props),
	)
	return Chart(props)
}
