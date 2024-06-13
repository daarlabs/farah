package chart_ui

import "github.com/daarlabs/arcanum/mirage"

type Props struct {
	Id     string
	Defer  bool
	Dark   bool
	DataX  any
	DataY  any
	Config mirage.Map
}
