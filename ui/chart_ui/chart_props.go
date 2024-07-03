package chart_ui

import "github.com/daarlabs/hirokit/hiro"

type Props struct {
	Id     string
	Defer  bool
	Dark   bool
	DataX  any
	DataY  any
	Config hiro.Map
	Height int
}
