package tab_component

import . "github.com/daarlabs/arcanum/gox"

type Props struct {
	Name   string
	Tabs   []Tab
	Active string
}

type Tab struct {
	Title    string
	Name     string
	NodeFunc func() Node
}
