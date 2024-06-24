package hidden_field_ui

import . "github.com/daarlabs/arcanum/gox"

func HiddenField(name string, value any) Node {
	var v any
	switch vv := value.(type) {
	case bool:
		if !vv {
			v = ""
		}
		if vv {
			v = "on"
		}
	default:
		v = vv
	}
	return Input(
		Type("hidden"),
		Name(name),
		Value(v),
	)
}
