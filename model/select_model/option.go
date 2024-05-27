package select_model

type Option[T any] struct {
	Text  string `json:"text"`
	Link  string `json:"link"`
	Value T      `json:"value"`
}

func ConvertToMapSlice[T comparable](options []Option[T]) []map[string]any {
	result := make([]map[string]any, len(options))
	for i, option := range options {
		result[i] = map[string]any{
			"value": option.Value,
			"text":  option.Text,
		}
	}
	return result
}
