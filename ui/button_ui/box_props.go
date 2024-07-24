package button_ui

import "github.com/daarlabs/hirokit/tempest"

type Props struct {
	Class   tempest.Tempest
	Icon    string
	Size    string
	Type    string
	Link    string
	Pending bool
}

const (
	TypeButton = "button"
	TypeSubmit = "submit"
)
