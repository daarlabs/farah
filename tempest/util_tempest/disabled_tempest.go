package util_tempest

import "github.com/daarlabs/hirokit/tempest"

func Disabled() tempest.Tempest {
	return tempest.Class().UserSelect("none").PointerEvents("none").Opacity(30)
}
