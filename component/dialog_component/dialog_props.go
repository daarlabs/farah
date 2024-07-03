package dialog_component

import . "github.com/daarlabs/hirokit/gox"

type Props struct {
	Open  bool
	Title string
	Name  string
	Nodes []Node
}
