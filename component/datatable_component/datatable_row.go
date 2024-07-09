package datatable_component

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
	
	"github.com/daarlabs/farah/ui"
)

type RowBuilder[T any] interface {
	Data() T
	Row(nodes ...Node) Node
	Field(node Node, config ...FieldConfig) Node
	Link(link string, nodes ...Node) Node
}

type FieldConfig struct {
	Truncate bool
}

type rowBuilder[T any] struct {
	data       T
	fields     []Field
	sizeStyle  Node
	loadMore   Node
	rowIndex   int
	fieldIndex int
}

func (b *rowBuilder[T]) Data() T {
	return b.data
}

func (b *rowBuilder[T]) Link(link string, nodes ...Node) Node {
	return A(
		tempest.Class().Transition().Block().
			BgSlate(100, tempest.Hover()).BgSlate(700, tempest.Dark(), tempest.Hover()),
		Href(link),
		Fragment(nodes...),
	)
}

func (b *rowBuilder[T]) Row(nodes ...Node) Node {
	if len(b.fields) != len(nodes) {
		panic(ErrorFieldsLenMismtach)
	}
	return Div(
		tempest.Class().Grid(),
		b.sizeStyle,
		If(b.loadMore != nil, b.loadMore),
		Fragment(nodes...),
	)
}

func (b *rowBuilder[T]) Field(node Node, config ...FieldConfig) Node {
	var cfg FieldConfig
	if len(config) == 0 {
		cfg.Truncate = true
	}
	field := b.findField()
	if len(field.Name) == 0 {
		return Fragment()
	}
	if len(config) > 0 {
		cfg = config[0]
	}
	return Div(
		tempest.Class().Transition().Px(4).H(10).TextSize("10px").
			Flex().ItemsCenter().
			TextSlate(900).TextWhite(tempest.Dark()).
			BorderB(1).BorderSlate(300).BorderSlate(600, tempest.Dark()).
			Overflow("hidden").
			If(field.AlignX == ui.Left, tempest.Class().JustifyStart()).
			If(field.AlignX == ui.Center, tempest.Class().JustifyCenter()).
			If(field.AlignX == ui.Right, tempest.Class().JustifyEnd()),
		If(
			cfg.Truncate,
			Span(tempest.Class().Truncate(), node),
		),
		If(
			!cfg.Truncate,
			node,
		),
	)
}

func (b *rowBuilder[T]) findField() Field {
	if len(b.fields)-1 < b.fieldIndex {
		return Field{}
	}
	f := b.fields[b.fieldIndex]
	b.fieldIndex += 1
	return f
}
