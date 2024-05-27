package text_field_ui

import (
	"github.com/daarlabs/arcanum/form"
)

type Props struct {
	Id          string
	Name        string
	Label       string
	Value       string
	Placeholder string
	Messages    []string
	Autofocus   bool
	Disabled    bool
	Required    bool
}

const (
	TypeText     = "text"
	TypeEmail    = "email"
	TypePassword = "password"
)

func CreateProps(field form.Field[string]) Props {
	return Props{
		Id:        field.Id,
		Name:      field.Name,
		Label:     field.Label,
		Value:     field.Value,
		Messages:  field.Messages,
		Disabled:  field.Disabled,
		Autofocus: field.Autofocus,
		Required:  field.Required,
	}
}
