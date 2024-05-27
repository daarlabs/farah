package multiautocomplete_feature

import "github.com/daarlabs/arcanum/form"

type Props[T comparable] struct {
	Id       string
	Label    string
	Name     string
	Text     string
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
