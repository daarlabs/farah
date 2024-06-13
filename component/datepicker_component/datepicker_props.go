package datepicker_component

import (
	"time"
	
	"github.com/daarlabs/arcanum/form"
)

type Props struct {
	Id       string
	Label    string
	Name     string
	Value    time.Time
	Required bool
}

func CreateProps(field form.Field[time.Time]) Props {
	return Props{
		Id:       field.Id,
		Label:    field.Label,
		Name:     field.Name,
		Value:    field.Value,
		Required: field.Required,
	}
}
