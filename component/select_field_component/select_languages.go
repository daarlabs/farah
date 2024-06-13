package select_field_component

import (
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/farah/model/select_model"
)

func CreateLanguagesOptions(c mirage.Ctx) []select_model.Option[string] {
	langs := c.Config().Localization.Languages
	n := len(c.Config().Localization.Languages)
	result := make([]select_model.Option[string], n)
	for i, lang := range langs {
		result[i] = select_model.Option[string]{
			Text:  c.Translate("lang." + lang.Code),
			Value: lang.Code,
		}
	}
	return result
}
