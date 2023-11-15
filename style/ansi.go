/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package style

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
