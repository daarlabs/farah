package hidden_field_ui

import . "github.com/daarlabs/arcanum/gox"

func HiddenField(name string, value any) Node {
	return Input(
		Type("hidden"),
		Name(name),
		Value(value),
	)
}
