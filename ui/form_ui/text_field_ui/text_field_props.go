package text_field_ui

import (
	"github.com/daarlabs/farah/ui/form_ui"
	"github.com/daarlabs/hirokit/form"
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
	Boxed       bool
	Status      string
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

func (p Props) Box() Props {
	p.Boxed = true
	return p
}

func (p Props) Success(use bool) Props {
	if use {
		p.Status = form_ui.StatusSuccess
	}
	return p
}

func (p Props) Error(use bool) Props {
	if use {
		p.Status = form_ui.StatusError
	}
	return p
}
