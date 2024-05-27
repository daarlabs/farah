package datatable_feature

import (
	"reflect"
	
	"github.com/iancoleman/strcase"
)

func convertDataToMapSlice[T any](data []T) []map[string]any {
	result := make([]map[string]any, len(data))
	for i, item := range data {
		itemMap := make(map[string]any)
		v := reflect.ValueOf(item)
		t := reflect.TypeOf(item)
		for j := 0; j < v.NumField(); j++ {
			field := t.Field(j)
			itemMap[strcase.ToLowerCamel(field.Name)] = v.Field(j).Interface()
		}
		result[i] = itemMap
	}
	return result
}
