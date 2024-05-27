package datatable_feature

import (
	. "github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/farah/ui"
)

type RowBuilder[T any] interface {
	Data() T
	Row(nodes ...Node) Node
	Field(node Node, truncate ...bool) Node
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

func (b *rowBuilder[T]) Row(nodes ...Node) Node {
	if len(b.fields) != len(nodes) {
		panic(ErrorFieldsLenMismtach)
	}
	return Div(
		Class("grid"),
		b.sizeStyle,
		If(b.loadMore != nil, b.loadMore),
		Fragment(nodes...),
	)
}

func (b *rowBuilder[T]) Field(node Node, truncate ...bool) Node {
	field := b.findField()
	if len(field.Name) == 0 {
		return Fragment()
	}
	shouldTruncate := true
	if len(truncate) > 0 {
		shouldTruncate = truncate[0]
	}
	return Div(
		Clsx{
			"transition text-[10px] text-slate-900 dark:text-white px-4 h-10 flex items-center": true,
			"border-b border-slate-300 dark:border-slate-600":                                   true,
			"truncate":       shouldTruncate,
			"justify-start":  field.AlignX == ui.Left,
			"justify-center": field.AlignX == ui.Center,
			"justify-end":    field.AlignX == ui.Right,
		},
		node,
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
