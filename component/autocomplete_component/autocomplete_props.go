package autocomplete_component

import (
	"github.com/daarlabs/arcanum/form"
)

type Props[T comparable] struct {
	Id       string
	Label    string
	Name     string
	Text     string
	Value    T
	Messages []string
	Required bool
}

type Query struct {
	Table string
	Title string
	Value string
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

func CreateQuery(table, title, value string) Query {
	return Query{
		Table: table,
		Title: title,
		Value: value,
	}
}
