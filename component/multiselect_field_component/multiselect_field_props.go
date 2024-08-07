package multiselect_field_component

import "github.com/daarlabs/hirokit/form"

type Props[T comparable] struct {
	Id       string
	Label    string
	Name     string
	Value    []T
	Messages []string
	Required bool
}

func CreateProps[T comparable](field form.Field[[]T]) Props[T] {
	return Props[T]{
		Id:       field.Id,
		Label:    field.Label,
		Name:     field.Name,
		Value:    field.Value,
		Required: field.Required,
		Messages: field.Messages,
	}
}
