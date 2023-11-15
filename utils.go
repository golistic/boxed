// Copyright (c) 2021, Geert JM Vanderkelen

package boxed

import (
	"fmt"
	"io"
)

// Fprint convenient wrapper around fmt.Fprint ignoring errors.
func Fprint(w io.Writer, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}
