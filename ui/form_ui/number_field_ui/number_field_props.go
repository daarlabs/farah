package number_field_ui

import (
	"golang.org/x/exp/constraints"
	
	"github.com/daarlabs/hirokit/form"
)

type Props[T constraints.Integer | constraints.Float] struct {
	Id        string
	Name      string
	Label     string
	Value     T
	Messages  []string
	Autofocus bool
	Disabled  bool
	Required  bool
	Boxed     bool
}

func CreateProps[T constraints.Integer | constraints.Float](field form.Field[T]) Props[T] {
	return Props[T]{
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
