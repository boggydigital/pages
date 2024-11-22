package color

import (
	_ "embed"
	"golang.org/x/exp/maps"
)

type Color int

const (
	Unset Color = iota
	Black
	White
	Red
	Orange
	Yellow
	Green
	Mint
	Teal
	Cyan
	Blue
	Indigo
	Purple
	Pink
	Brown
	Background
	Foreground
	Gray
	Highlight
	Transparent
)

var colorStrings = map[Color]string{
	Black:       "black",
	White:       "white",
	Red:         "red",
	Orange:      "orange",
	Yellow:      "yellow",
	Green:       "green",
	Mint:        "mint",
	Teal:        "teal",
	Cyan:        "cyan",
	Blue:        "blue",
	Indigo:      "indigo",
	Purple:      "purple",
	Pink:        "pink",
	Brown:       "brown",
	Gray:        "gray",
	Background:  "background",
	Foreground:  "foreground",
	Highlight:   "highlight",
	Transparent: "transparent",
}

func (c Color) String() string {
	return colorStrings[c]
}

func (c Color) CssValue() string {
	return "var(--c-" + c.String() + ")"
}

func Parse(s string) Color {
	for c, str := range colorStrings {
		if s == str {
			return c
		}
	}
	return Unset
}

func All() []Color {
	return maps.Keys(colorStrings)
}
