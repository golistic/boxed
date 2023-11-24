// Copyright (c) 2021, Geert JM Vanderkelen

package boxed

import (
	"strings"
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

	t.Run("with ANSI codes (1)", func(t *testing.T) {

		exp := strings.Join([]string{
			"",
			"┌────┬───────────────────────────────────────────────┬────┐",
			"│ \u001B[32m1A\u001B[0m │ 1B                                            │ 1C │",
			"│ 2A │ \u001B[1mThe world's most popular open source database\u001B[0m │ 2C │",
			"└────┴───────────────────────────────────────────────┴────┘",
			"",
		}, "\n")

		box := New()
		xt.OK(t, box.Append([]Row{
			NewRow("\u001B[32m1A\u001B[0m", "1B", "1C"),
			NewRow("2A", "\033[1mThe world's most popular open source database\033[0m", "2C"),
		}...))

		xt.Eq(t, exp, "\n"+box.RenderAsString()) // extra newline for pretty output
	})

	t.Run("with ANSI codes (2)", func(t *testing.T) {

		exp := strings.Join([]string{
			"",
			"┌────┬───────────────────────────────────────────────┐",
			"│ 2A │ \u001B[1mThe world's most popular open source database\u001B[0m │",
			"│ \u001B[32m1A\u001B[0m │ 1B                                            │",
			"└────┴───────────────────────────────────────────────┘",
			"",
		}, "\n")

		box := New()
		xt.OK(t, box.Append([]Row{
			NewRow("2A", "\033[1mThe world's most popular open source database\u001B[0m"),
			NewRow("\u001B[32m1A\u001B[0m", "1B"),
		}...))

		xt.Eq(t, exp, "\n"+box.RenderAsString()) // extra newline for pretty output
	})

	t.Run("with ANSI codes (3)", func(t *testing.T) {

		exp := strings.Join([]string{
			"",
			"┌────┬───────────────────────────────────────────────┐",
			"│ \u001B[2J2A │ \u001B[1mThe world's most popular open source database\u001B[0m │",
			"│ 1A │ \u001B[32m1B\u001B[0m                                            │",
			"└────┴───────────────────────────────────────────────┘",
			"",
		}, "\n")

		box := New()
		xt.OK(t, box.Append([]Row{
			NewRow("\033[2J2A", "\033[1mThe world's most popular open source database\u001B[0m"),
			NewRow("1A", "\u001B[32m1B\u001B[0m"),
		}...))

		have := "\n" + box.RenderAsString() // extra newline for pretty output

		xt.Eq(t, exp, have)
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
