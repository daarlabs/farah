package phone_field_ui

import (
	"github.com/daarlabs/farah/model/select_model"
	"github.com/daarlabs/hirokit/form"
	"github.com/daarlabs/hirokit/gox"
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
	Options     []select_model.Option[string]
	PrefixField gox.Node
}

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
