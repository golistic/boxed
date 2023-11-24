/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package boxed

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestStripNoneGraphic(t *testing.T) {
	t.Run("strip text formatting", func(t *testing.T) {
		var cases = []struct {
			got string
			exp string
		}{
			{
				got: "\033[1mbold\033[0m",
				exp: "bold",
			},
			{
				got: "\033[2Jclearing screen",
				exp: "clearing screen",
			},
			{
				got: "\033[no code",
				exp: "o code", // corrupts the string, but that's how it works
			},
			{
				got: "\033rogue escape",
				exp: "rogue escape",
			},
			{
				got: "\033]0;Window title change\a",
				exp: "]0;Window title change",
			},
		}

		for _, c := range cases {
			t.Run(c.exp, func(t *testing.T) {
				xt.Eq(t, c.exp, StripNoneGraphic(c.got))
			})
		}
	})
}
