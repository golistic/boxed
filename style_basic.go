/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package boxed

var Basic = &Style{
	Horizontal:       "─",
	HorizontalTop:    "─",
	HorizontalBottom: "─",
	Vertical:         "│",
	VerticalLeft:     "│",
	VerticalRight:    "│",
	Cross:            "┼",
	LeftCross:        "├",
	RightCross:       "┤",
	TopLeft:          "┌",
	TopRight:         "┐",
	TopCross:         "┬",
	BottomLeft:       "└",
	BottomRight:      "┘",
	BottomCross:      "┴",
}

var BasicNoOuter = &Style{
	Horizontal: "─",
	Vertical:   "│",
	Cross:      "┼",

	Header: &Style{
		HorizontalBottom: "═",
		Vertical:         "│",
		Cross:            "╪", // double horizontal
	},
}
