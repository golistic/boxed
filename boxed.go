// Copyright (c) 2021, 2023, Geert JM Vanderkelen

package boxed

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	style2 "github.com/golistic/boxed/style"
)

// Boxed defines a headers, rows with cells and how to display them.
type Boxed struct {
	writeTo   io.Writer
	rows      []Row
	headers   []Row
	cellWidth []int
	style     *style2.Style
}

// New instantiates a box instance.
func New(opts ...*Option) *Boxed {

	o := newOptions()

	for _, opt := range opts {
		opt.apply(o)
	}

	return &Boxed{
		writeTo: os.Stdout,
		style:   o.style,
	}
}

func (bx *Boxed) addRow(dst *[]Row, row Row) error {

	if row.Len() == 0 {
		// nothing to do
		return nil
	}

	if len(bx.cellWidth) == 0 {
		bx.cellWidth = make([]int, row.Len())
	}

	if row.Len() != len(bx.cellWidth) {
		return fmt.Errorf("row length must be %d", len(bx.cellWidth))
	}

	for i, col := range row.cells {
		v := fmt.Sprintf("%v", col)
		l := len(StripNoneGraphic(v))
		if bx.cellWidth[i] < l {
			bx.cellWidth[i] = l
		}
	}
	*dst = append(*dst, row)

	return nil
}

func (bx *Boxed) AddHeader(head ...any) error {

	row := NewRow(head...)
	row.isHeader = true

	return bx.addRow(&bx.headers, row)
}

func (bx *Boxed) Append(rows ...Row) error {

	for _, row := range rows {
		if err := bx.addRow(&bx.rows, row); err != nil {
			return err
		}
	}

	return nil
}

func (bx *Boxed) haveHeader() bool {
	return len(bx.headers) > 0
}

func (bx *Boxed) top(w io.Writer, header bool) {

	b := bx.style
	if header && b.Header != nil {
		b = b.Header
	}

	top := b.TopLeft

	for i, width := range bx.cellWidth {
		top += strings.Repeat(b.HorizontalTop, width+2)
		if i == len(bx.cellWidth)-1 {
			top += b.TopRight
		} else {
			top += b.TopCross
		}
	}

	if strings.TrimSpace(top) == "" {
		return
	}

	Fprint(w, top+"\n")
}

func (bx *Boxed) bottom(w io.Writer) {

	b := bx.style
	bottom := b.BottomLeft

	for i, width := range bx.cellWidth {
		bottom += strings.Repeat(b.HorizontalBottom, width+2)
		if i == len(bx.cellWidth)-1 {
			bottom += b.BottomRight
		} else {
			bottom += b.BottomCross
		}
	}

	if strings.TrimSpace(bottom) == "" {
		return
	}

	Fprint(w, bottom+"\n")
}

func (bx *Boxed) headerBottom(w io.Writer) {

	b := bx.style
	if b.Header != nil {
		b = b.Header
	}

	bottom := b.LeftCross

	for i, width := range bx.cellWidth {
		bottom += strings.Repeat(b.HorizontalBottom, width+2)
		if i == len(bx.cellWidth)-1 {
			bottom += b.RightCross
		} else {
			bottom += b.Cross
		}
	}

	if strings.TrimSpace(bottom) == "" {
		return
	}

	Fprint(w, bottom+"\n")
}

func (bx *Boxed) row(w io.Writer, row Row, header bool) error {

	b := bx.style
	if header && b.Header != nil {
		b = b.Header
	}

	for j, col := range row.cells {
		if j == 0 {
			Fprint(w, b.VerticalLeft)
			if b.VerticalLeft != "" || bx.haveHeader() {
				Fprint(w, " ")
			}
		} else {
			Fprint(w, " ")
		}

		v := fmt.Sprintf("%v", col)
		content := v + strings.Repeat(" ", bx.cellWidth[j]-len(StripNoneGraphic(v)))
		Fprint(w, content)

		if j == len(row.cells)-1 {
			if b.VerticalRight != "" || bx.haveHeader() {
				Fprint(w, " ")
			}
			Fprint(w, b.VerticalRight+"\n")
		} else {
			Fprint(w, " ")
			Fprint(w, b.Vertical)
		}
	}

	return nil
}

func (bx *Boxed) render(w io.Writer) error {

	if len(bx.headers) > 0 {
		for i := -1; i < len(bx.headers)+1; i++ {
			if i == -1 {
				bx.top(w, true)
				continue
			}

			if i == len(bx.headers) {
				if len(bx.rows) > 0 {
					bx.headerBottom(w)
				} else {
					bx.bottom(w)
				}
				continue
			}

			if err := bx.row(w, bx.headers[i], true); err != nil {
				return err
			}
		}
	}

	if len(bx.rows) > 0 {
		for i := -1; i < len(bx.rows)+1; i++ {
			if i == -1 {
				if len(bx.headers) == 0 {
					bx.top(w, false)
				}
				continue
			}

			if i == len(bx.rows) {
				bx.bottom(w)
				continue
			}

			if err := bx.row(w, bx.rows[i], false); err != nil {
				return err
			}
		}
	}

	return nil
}

// Render will print the result to STDOUT. Use RenderTo to write to an io.Writer,
// or RenderAsString to get the string.
func (bx *Boxed) Render() error {
	return bx.render(os.Stdout)
}

// RenderTo will write the result to w. Use Render to directly write to STDOUT,
// or RenderAsString to get the string.
func (bx *Boxed) RenderTo(w io.Writer) error {

	return bx.render(w)
}

// RenderAsString will return the result as string. Use Render to directly write to STDOUT,
// or RenderTo to write to an io.Writer.
// Note that errors are made part of the string.
func (bx *Boxed) RenderAsString() string {

	w := new(bytes.Buffer)
	if err := bx.render(w); err != nil {
		return "(boxed) " + err.Error()
	}

	buf, err := io.ReadAll(w)
	if err != nil {
		return "(boxed) " + err.Error()
	}

	return string(buf)
}
