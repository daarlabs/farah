package radio_ui

import (
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	"github.com/daarlabs/farah/palette"
	
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/field_label_ui"
)

func Radio(props Props) gox.Node {
	return gox.Div(
		tempest.Class().Flex().FlexCol().Gap(1),
		gox.If(
			len(props.Label) > 0,
			field_label_ui.FieldLabel(
				field_label_ui.Props{
					For:      props.Id,
					Text:     props.Label,
					Required: props.Required,
				},
			),
		),
		gox.Div(
			tempest.Class().Grid().Gap(2),
			gox.Range(
				props.Options,
				func(item Option, index int) gox.Node {
					return gox.Div(
						tempest.Class().Flex().ItemsCenter().Gap(2),
						gox.Input(
							gox.If(len(props.Id) > 0, gox.Id(props.Id+"-"+item.Value)),
							tempest.Class().Invisible().Peer(),
							gox.Type("radio"),
							gox.Name(props.Name),
							gox.Value(item.Value),
							gox.If(item.Checked, gox.Checked()),
							gox.If(props.Disabled, gox.Disabled()),
						),
						gox.Label(
							gox.For(props.Id+"-"+item.Value),
							tempest.Class().Group().Relative().Block().FlexNone().Size(4).RoundedFull().
								BgTransparent().BorderSlate(300).BorderSlate(600, tempest.Dark()).
								BorderColor(palette.Primary, 400, tempest.Checked(tempest.Peer)).
								BorderColor(palette.Primary, 200, tempest.Dark(), tempest.Checked(tempest.Peer)),
							gox.Div(
								tempest.Class().Transition().Absolute().Inset(0).M("auto").Size(1).RoundedFull().
									BgTransparent().BgTransparent(tempest.Dark()),
								// TODO: Peer checked group
								// gox.Class("peer-checked:group-[]:bg-primary-400 dark:peer-checked:group-[]:bg-primary-200"),
							),
						),
						gox.If(
							len(item.Title) > 0,
							field_label_ui.FieldLabel(
								field_label_ui.Props{
									For:  props.Id + "-" + item.Value,
									Text: item.Title,
								},
							),
						),
					)
				},
			),
		),
		gox.Range(
			props.Messages, func(msg string, _ int) gox.Node {
				return error_message_ui.ErrorMessage(msg)
			},
		),
	)
}
