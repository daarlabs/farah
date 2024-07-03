package filepicker_ui

import "github.com/daarlabs/hirokit/form"

type Props struct {
	Id       string
	Name     string
	Label    string
	Value    form.Multipart
	Messages []string
	Disabled bool
	Required bool
}

func CreateProps(field form.Field[form.Multipart]) Props {
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
