/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package boxed

import "github.com/golistic/boxed/style"

type options struct {
	style *style.Style
}

func newOptions() *options {
	return &options{style: style.Basic}
}

func (o *options) apply(fOpts ...Option) {
	for _, f := range fOpts {
		f.apply(o)
	}
}

type Option struct {
	f func(*options)
}

func (fo *Option) apply(opts *options) {
	fo.f(opts)
}

func newBoxedOption(f func(*options)) *Option {
	return &Option{
		f: f,
	}
}

// WithStyle sets the styling used to draw the boxes (lines, corners, ...).
// The default style is boxed.Basic
func WithStyle(s *style.Style) *Option {
	return newBoxedOption(func(o *options) {
		o.style = s
	})
}
