/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package boxed

var ANSI = &Style{
	Horizontal:       "-",
	HorizontalTop:    "-",
	HorizontalBottom: "-",
	Vertical:         "|",
	VerticalLeft:     "|",
	VerticalRight:    "|",
	Cross:            "+",
	LeftCross:        "+",
	RightCross:       "+",
	TopLeft:          "+",
	TopRight:         "+",
	TopCross:         "+",
	BottomLeft:       "+",
	BottomRight:      "+",
	BottomCross:      "+",
}

var ANSINoOuter = &Style{
	Horizontal:       "-",
	HorizontalTop:    "",
	HorizontalBottom: "",
	Vertical:         "|",
	Cross:            "+",
	Header: &Style{
		HorizontalBottom: "-",
		Vertical:         "|",
		Cross:            "+",
	},
}
