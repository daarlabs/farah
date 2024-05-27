package radio_ui

type Props struct {
	Id       string
	Name     string
	Label    string
	Messages []string
	Disabled bool
	Required bool
	Options  []Option
}

type Option struct {
	Value   string
	Title   string
	Checked bool
}
