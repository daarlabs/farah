package checkbox_ui

import "github.com/daarlabs/arcanum/form"

type Props struct {
	Id       string
	Name     string
	Label    string
	Messages []string
	Value    any
	Checked  bool
	Required bool
}

func CreateProps(field form.Field[bool]) Props {
	return Props{
		Id:       field.Id,
		Name:     field.Name,
		Label:    field.Label,
		Checked:  field.Value,
		Required: field.Required,
	}
}
