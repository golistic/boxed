/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package boxed

import (
	"strings"
	"unicode"
)

// StripNoneGraphic removes any CSI ANSI escape code sequences and
// non-Graphic runes from the s.
//
// CSI, Control Sequence Introducer, are completely removed.
// OSC, Operating System Command, (or anything else) are not removed, but since we
// remove the escape code, the characters will be part of the resulting
// string.
//
// This function is not meant to cleanse malicious codes in strings. It is meant
// to be used to calculate string lengths within the boxed-package.
func StripNoneGraphic(s string) string {

	var stripped strings.Builder
	inCode := false

	runes := []rune(s)
	sl := len(runes)

	for i := 0; i < sl; i++ {
		r := runes[i]

		switch {
		case r == 0x1b: // or 27 or \033 or \u001b
			if sl-1 != i && (s[i+1] == '[') {
				inCode = true
				i++
			}
			// ignore the escape
		case inCode && (r >= 0x40 && r <= 0x7e):
			inCode = false
		case !inCode:
			if unicode.IsGraphic(r) {
				stripped.WriteRune(r)
			}
		}
	}

	return stripped.String()
}
