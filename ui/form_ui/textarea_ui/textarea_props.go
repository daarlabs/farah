package textarea_ui

import "github.com/daarlabs/hirokit/form"

type Props struct {
	Id       string
	Name     string
	Label    string
	Value    string
	Type     string
	Messages []string
	Disabled bool
	Required bool
	Boxed    bool
}

func CreateProps(field form.Field[string]) Props {
	return Props{
		Id:       field.Id,
		Name:     field.Name,
		Label:    field.Label,
		Value:    field.Value,
		Messages: field.Messages,
		Disabled: field.Disabled,
		Required: field.Required,
	}
}

func (p Props) Box() Props {
	p.Boxed = true
	return p
}
