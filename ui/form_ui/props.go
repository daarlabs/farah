package form_ui

type Props struct {
	Id          string
	Name        string
	Label       string
	Value       string
	Placeholder string
	Status      string
	Messages    []string
	Autofocus   bool
	Disabled    bool
	Required    bool
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)
