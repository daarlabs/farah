package dialog_component

import . "github.com/daarlabs/arcanum/gox"

type Props struct {
	Open  bool
	Title string
	Name  string
	Nodes []Node
}
