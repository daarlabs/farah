package select_field_component

import (
	"github.com/daarlabs/arcanum/form"
	"github.com/daarlabs/farah/model/select_model"
)

type Props[T comparable] struct {
	Id          string
	Label       string
	Name        string
	Text        string
	Placeholder string
	Value       T
	Messages    []string
	Required    bool
	Size        string
	Refresh     bool
	OnChange    func(option select_model.Option[T])
}

func CreateProps[T comparable](field form.Field[T]) Props[T] {
	return Props[T]{
		Id:       field.Id,
		Label:    field.Label,
		Name:     field.Name,
		Value:    field.Value,
		Required: field.Required,
		Messages: field.Messages,
		Text:     field.Text,
	}
}
