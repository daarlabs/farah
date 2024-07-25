package dialog_component

import . "github.com/daarlabs/hirokit/gox"

type Props struct {
	Open       bool
	Submitable bool
	Title      string
	Name       string
	Content    Node
}
