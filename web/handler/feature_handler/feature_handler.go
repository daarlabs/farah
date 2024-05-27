package feature_handler

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/mystiq"
	
	"component/feature/autocomplete_feature"
	"component/feature/datatable_feature"
	"component/feature/datepicker_feature"
	"component/feature/dialog_feature"
	"component/feature/multiautocomplete_feature"
	"component/feature/multiselect_field_feature"
	"component/feature/select_field_feature"
	"component/feature/tab_feature"
	"component/model/fake_model"
	"component/model/select_model"
	"component/ui/button_ui"
	"component/ui/page_ui"
	"component/ui/page_ui/breadcrumbs_ui"
)

func Get() mirage.Handler {
	return func(c mirage.Ctx) error {
		return c.Response().Render(
			page_ui.Page(
				page_ui.Props{},
				page_ui.Header(
					breadcrumbs_ui.Breadcrumbs(
						c.Generate().Link("home"),
						breadcrumbs_ui.Breadcrumb(c.Generate().Link("feature"), "Feature", true),
					),
				),
				page_ui.Content(
					page_ui.Title("Feature"),
					c.Create().Component(
						&tab_feature.TabFeature{
							Props: tab_feature.Props{
								Name: "test",
								Tabs: []tab_feature.Tab{
									{Title: "Tab 1", Name: "tab1", NodeFunc: func() Node { return Text("Tab 1") }},
									{Title: "Tab 2", Name: "tab2", NodeFunc: func() Node { return Text("Tab 2") }},
									{Title: "Tab 3", Name: "tab3", NodeFunc: func() Node { return Text("Tab 3") }},
								},
							},
						},
					),
					c.Create().Component(
						&dialog_feature.DialogFeature{
							Props:  dialog_feature.Props{Name: "test", Title: "Submit dialog"},
							Submit: dialog_feature.Config{Link: c.Generate().Link("feature")},
							HandlerFunc: func(action Node) Node {
								return button_ui.PrimaryButton(button_ui.Props{}, action, Text("Open dialog"))
							},
						},
					),
					Div(
						Class("grid grid-cols-6 gap-8"),
						c.Create().Component(
							&select_field_feature.SelectField[int]{
								Props: select_field_feature.Props[int]{
									Id:    "test-select",
									Name:  "test-select",
									Label: "Select field",
								},
								Options: []select_model.Option[int]{
									{Text: "One", Value: 1},
									{Text: "Two", Value: 2},
									{Text: "Three", Value: 3},
								},
							},
						),
						c.Create().Component(
							&multiselect_field_feature.MultiSelectField[int]{
								Props: multiselect_field_feature.Props[int]{
									Id:    "test-multiselect",
									Name:  "test-multiselect",
									Label: "Multiselect field",
								},
								Options: []select_model.Option[int]{
									{Text: "One", Value: 1},
									{Text: "Two", Value: 2},
									{Text: "Three", Value: 3},
								},
							},
						),
						c.Create().Component(
							&autocomplete_feature.Autocomplete[int]{
								Props: autocomplete_feature.Props[int]{
									Id:    "test-autocomplete",
									Name:  "test-autocomplete",
									Label: "Autocomplete field",
								},
								Options: fake_model.Categories,
							},
						),
						c.Create().Component(
							&multiautocomplete_feature.MultiAutocomplete[int]{
								Props: multiautocomplete_feature.Props[int]{
									Id:    "test-multiautocomplete",
									Name:  "test-multiautocomplete",
									Label: "Multiautocomplete field",
								},
								Options: fake_model.Categories,
								Query:   mystiq.Query{Name: "value"},
							},
						),
						c.Create().Component(
							&datepicker_feature.Datepicker{
								Props: datepicker_feature.Props{
									Id:    "test-datepicker",
									Name:  "test-datepicker",
									Label: "Datepicker field",
								},
							},
						),
					),
					Div(
						Class("w-full h-[500px] overflow-hidden mt-16"),
						c.Create().Component(
							&datatable_feature.Datatable[select_model.Option[int]]{
								Props: datatable_feature.Props{
									Name: "test-datatable",
								},
								FieldsFunc: func() []datatable_feature.Field {
									return []datatable_feature.Field{
										{Size: "1fr", Sortable: true, Name: "value"},
										{Size: "1fr", Sortable: true, Name: "text"},
									}
								},
								RowFunc: func(builder datatable_feature.RowBuilder[select_model.Option[int]]) Node {
									return builder.Row(
										builder.Field(Text(builder.Data().Value)),
										builder.Field(Text(builder.Data().Text)),
									)
								},
								Data: fake_model.Categories,
							},
						),
					),
				),
			),
		)
	}
}
