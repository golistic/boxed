/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package style

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
