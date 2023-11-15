/*
 * Copyright (c) 2021, 2023, Geert JM Vanderkelen
 */

package boxed

// Row is a collection of cells. It can be either header or data.
type Row struct {
	isHeader bool
	cells    []any
}

// NewRow instantiates a new Row with given cells as data.
func NewRow(cells ...any) Row {

	return Row{cells: cells}
}

// Append can be used to add more data to r as it comes available.
func (r *Row) Append(cells ...any) {

	r.cells = append(r.cells, cells...)
}

// Len returns the number of cells within r.
func (r *Row) Len() int {

	return len(r.cells)
}
