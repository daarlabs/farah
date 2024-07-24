package alert_ui

import "github.com/daarlabs/hirokit/tempest"

type Props struct {
	Type  string
	Class tempest.Tempest
}

var (
	AlertSuccess = "success"
	AlertWarning = "warning"
	AlertError   = "error"
	AlertInfo    = "info"
)
