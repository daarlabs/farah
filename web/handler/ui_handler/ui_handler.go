package ui_handler

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	
	"github.com/daarlabs/farah/ui/button_ui"
	"github.com/daarlabs/farah/ui/form_ui/checkbox_ui"
	"github.com/daarlabs/farah/ui/form_ui/number_field_ui"
	"github.com/daarlabs/farah/ui/form_ui/radio_ui"
	"github.com/daarlabs/farah/ui/form_ui/text_field_ui"
	"github.com/daarlabs/farah/ui/form_ui/textarea_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	"github.com/daarlabs/farah/ui/page_ui"
	"github.com/daarlabs/farah/ui/page_ui/breadcrumbs_ui"
	"github.com/daarlabs/farah/ui/page_ui/flash_message_ui"
	"github.com/daarlabs/farah/ui/spinner_ui"
)

func Get() mirage.Handler {
	return func(c mirage.Ctx) error {
		return c.Response().Render(
			page_ui.Page(
				page_ui.Props{},
				page_ui.Header(
					breadcrumbs_ui.Breadcrumbs(
						c.Generate().Link("home"),
						breadcrumbs_ui.Breadcrumb(c.Generate().Link("ui"), "UI", true),
					),
				),
				page_ui.Content(
					page_ui.Title("UI"),
					Div(
						Class("grid gap-16"),
						// Buttons
						Div(
							Div(
								Class("grid gap-16 grid-cols-6"),
								Div(
									Class("grid gap-4"),
									page_ui.Subtitle("Buttons"),
									button_ui.PrimaryButton(button_ui.Props{Icon: icon_ui.ChevronRight}, Text("Primary button")),
									button_ui.EmeraldButton(button_ui.Props{Icon: icon_ui.Check}, Text("Emerald button")),
									button_ui.MainButton(button_ui.Props{Icon: icon_ui.Add}, Text("Main button")),
								),
								Div(
									Class("grid gap-4"),
									page_ui.Subtitle("Spinner"),
									Div(
										Class("relative size-16"),
										spinner_ui.Spinner(spinner_ui.Props{Visible: true}),
									),
								),
							),
						),
						// Input fields
						Div(
							page_ui.Subtitle("Inputs"),
							Div(
								Class("grid gap-16 grid-cols-6"),
								Div(
									Class("grid gap-4"),
									text_field_ui.TextField(
										text_field_ui.Props{
											Label: "Text field", Id: "test-field", Messages: []string{"error message"},
										},
									),
								),
								Div(
									Class("grid gap-4"),
									number_field_ui.NumberField[int](
										number_field_ui.Props[int]{
											Label: "Number field", Id: "test-number-field", Messages: []string{"error message"},
										},
									),
								),
								Div(
									Class("grid gap-4"),
									checkbox_ui.Checkbox(
										checkbox_ui.Props{
											Label: "Checkbox", Id: "test-checkbox", Required: true, Messages: []string{"error message"},
										},
									),
								),
								Div(
									Class("grid gap-4"),
									textarea_ui.TextArea(
										textarea_ui.Props{
											Label: "Textarea", Id: "test-textarea", Required: true, Messages: []string{"error message"},
										},
									),
								),
								Div(
									Class("grid gap-4"),
									radio_ui.Radio(
										radio_ui.Props{
											Id:    "test-radio",
											Name:  "test-radio",
											Label: "Test radio",
											Options: []radio_ui.Option{
												{Value: "1", Title: "Option 1", Checked: true},
												{Value: "2", Title: "Option 2"},
												{Value: "3", Title: "Option 3"},
											},
										},
									),
								),
							),
						),
						// Flash messages
						Div(
							page_ui.Subtitle("Flash messages"),
							Div(
								Class("grid gap-16 grid-cols-6"),
								flash_message_ui.Message(
									mirage.Message{
										Type:  mirage.FlashSuccess,
										Title: "Success",
										Value: "Operation was successful",
									},
								),
								flash_message_ui.Message(
									mirage.Message{
										Type:  mirage.FlashWarning,
										Title: "Warning",
										Value: "Some issues",
									},
								),
								flash_message_ui.Message(
									mirage.Message{
										Type:  mirage.FlashError,
										Title: "Error",
										Value: "Something wrong",
									},
								),
							),
						),
					),
				),
			),
		)
	}
}
