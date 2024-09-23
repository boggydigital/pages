package color

import (
	_ "embed"
	"iter"
	"maps"
)

//go:embed "style/colors.css"
var StyleSheet []byte

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
	Subtle
	Highlight
)

var colorStrings = map[Color]string{
	Black:      "black",
	White:      "white",
	Red:        "red",
	Orange:     "orange",
	Yellow:     "yellow",
	Green:      "green",
	Mint:       "mint",
	Teal:       "teal",
	Cyan:       "cyan",
	Blue:       "blue",
	Indigo:     "indigo",
	Purple:     "purple",
	Pink:       "pink",
	Brown:      "brown",
	Background: "background",
	Foreground: "foreground",
	Subtle:     "subtle",
	Highlight:  "highlight",
}

func (c Color) String() string {
	return colorStrings[c]
}

func (c Color) CssValue() string {
	return "var(--c-" + c.String() + ")"
}

func AllColors() iter.Seq[Color] {
	return maps.Keys(colorStrings)
}

func Parse(s string) Color {
	for c, str := range colorStrings {
		if s == str {
			return c
		}
	}
	return Unset
}
