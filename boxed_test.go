// Copyright (c) 2021, Geert JM Vanderkelen

package boxed

import (
	"testing"

	"github.com/golistic/xgo/xt"

	"github.com/golistic/boxed/style"
)

func TestBoxed_AddRows(t *testing.T) {
	t.Run("3 colums, 2 rows, dynamic", func(t *testing.T) {
		exp := `
┌────┬───────────────────────────────────────────────┬────┐
│ 1A │ 1B                                            │ 1C │
│ 2A │ The world's most popular open source database │ 2C │
└────┴───────────────────────────────────────────────┴────┘
`

		box := New()
		xt.OK(t, box.Append([]Row{
			NewRow("1A", "1B", "1C"),
			NewRow("2A", "The world's most popular open source database", "2C"),
		}...))

		xt.Eq(t, exp, "\n"+box.RenderAsString()) // extra newline for pretty output
	})
}

func TestBoxed_AddHeader(t *testing.T) {
	t.Run("header, now rows", func(t *testing.T) {
		exp := `
┌───┬─────┬───────┐
│ A │ B B │ C C C │
└───┴─────┴───────┘
`

		box := New()
		xt.OK(t, box.AddHeader("A", "B B", "C C C"))

		xt.Eq(t, exp, "\n"+box.RenderAsString()) // extra newline for pretty output
		xt.Eq(t, []rune(exp), []rune("\n"+box.RenderAsString()))
	})

	t.Run("header, 2 rows", func(t *testing.T) {

		exp := `
┌────┬─────┬───────┐
│ A  │ B B │ C C C │
├────┼─────┼───────┤
│ 1A │ 1B  │ 1C    │
│ 2A │ 2B  │ 2C    │
└────┴─────┴───────┘
`

		box := New()
		xt.OK(t, box.AddHeader("A", "B B", "C C C"))
		xt.OK(t, box.Append(
			NewRow("1A", "1B", "1C"),
			NewRow("2A", "2B", "2C"),
		))

		res := "\n" + box.RenderAsString()

		xt.Eq(t, exp, res) // extra newline for pretty output
		xt.Eq(t, []rune(exp), []rune("\n"+box.RenderAsString()))
	})

	t.Run("header, 2 rows plus different header style", func(t *testing.T) {

		exp := `
 A  │ B B │ C C C 
════╪═════╪═══════
 1A │ 1B  │ 1C    
 2A │ 2B  │ 2C    
`

		g := New(WithStyle(style.BasicNoOuter))

		xt.OK(t, g.AddHeader("A", "B B", "C C C"))
		xt.OK(t, g.Append(
			NewRow("1A", "1B", "1C"),
			NewRow("2A", "2B", "2C"),
		))

		res := "\n" + g.RenderAsString()

		xt.Eq(t, exp, res) // extra newline for pretty output
		xt.Eq(t, []rune(exp), []rune("\n"+g.RenderAsString()))
	})

	t.Run("without header or boxing, just colon", func(t *testing.T) {

		exp := `
Active : true                   
API    : https://api.example.com
`

		g := New(WithStyle(style.Record))

		xt.OK(t, g.Append(
			NewRow("Active", true),
			NewRow("API", "https://api.example.com"),
		))

		res := "\n" + g.RenderAsString()

		xt.Eq(t, exp, res) // extra newline for pretty output
	})
}
