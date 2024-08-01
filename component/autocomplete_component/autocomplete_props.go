package autocomplete_component

import (
	"github.com/daarlabs/farah/ui"
	"github.com/daarlabs/hirokit/form"
)

type Props[T comparable] struct {
	Id        string
	Label     string
	Name      string
	Text      string
	Value     T
	Messages  []string
	Required  bool
	PositionX string
	PositionY string
	Sync      string
}

type Query struct {
	Table string
	Title string
	Value string
}

func CreateProps[T comparable](field form.Field[T]) Props[T] {
	return Props[T]{
		Id:        field.Id,
		Label:     field.Label,
		Name:      field.Name,
		Value:     field.Value,
		Required:  field.Required,
		Messages:  field.Messages,
		Text:      field.Text,
		PositionX: ui.Left,
		PositionY: ui.Bottom,
	}
}

func CreateQuery(table, title, value string) Query {
	return Query{
		Table: table,
		Title: title,
		Value: value,
	}
}

func (p Props[T]) Top() Props[T] {
	p.PositionY = ui.Top
	return p
}

func (p Props[T]) Right() Props[T] {
	p.PositionX = ui.Right
	return p
}
