package class

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"maps"
	"strconv"
	"strings"
	"sync"
)

const (
	classSelectorPfx   = "."
	classNameSep       = "-"
	customPropertyPfx  = "--"
	rowGapPfx          = "rg"
	columnGapPfx       = "cg"
	alignContentPfx    = "ac"
	alignItemsPfx      = "ai"
	justifyContentPfx  = "jc"
	justifyItemsPfx    = "ji"
	flexDirectionPfx   = "fd"
	backgroundColorPfx = "bg"
	foregroundColorPfx = "fg"
	markerColorPfx     = "cm"
	fontSizePfx        = "fs"
	fontWeightPfx      = "fw"
	marginBlockEndPfx  = "mbe"
	gridTemplateRows   = "gtr"
)

var setClasses = make(map[string]any)
var mtx = sync.Mutex{}

func joinClassName(parts ...string) string {
	cn := strings.Join(parts, classNameSep)
	mtx.Lock()
	defer mtx.Unlock()
	setClasses[cn] = nil
	return cn
}

func classSelector(className string) string {
	return classSelectorPfx + className
}

func customProperty(className string) string {
	return customPropertyPfx + className
}

func RowGap(s size.Size) string {
	return joinClassName(rowGapPfx, s.String())
}

func ColumnGap(s size.Size) string {
	return joinClassName(columnGapPfx, s.String())
}

func AlignContent(a align.Align) string {
	return joinClassName(alignContentPfx, a.String())
}

func AlignItems(a align.Align) string {
	return joinClassName(alignItemsPfx, a.String())
}

func JustifyContent(a align.Align) string {
	return joinClassName(justifyContentPfx, a.String())
}

func JustifyItems(a align.Align) string {
	return joinClassName(justifyItemsPfx, a.String())
}

func FlexDirection(d direction.Direction) string {
	return joinClassName(flexDirectionPfx, d.String())
}

func BackgroundColor(c color.Color) string {
	return joinClassName(backgroundColorPfx, c.String())
}

func ForegroundColor(c color.Color) string {
	return joinClassName(foregroundColorPfx, c.String())
}

func MarkerColor(c color.Color) string {
	return joinClassName(markerColorPfx, c.String())
}

func FontSize(s size.Size) string {
	return joinClassName(fontSizePfx, s.String())
}

func FontWeight(w font_weight.Weight) string {
	return joinClassName(fontWeightPfx, w.String())
}

func MarginBlockEnd(s size.Size) string {
	return joinClassName(marginBlockEndPfx, s.String())
}

func GridTemplateRows(s size.Size) string {
	return joinClassName(gridTemplateRows, s.String())
}

func GridTemplateRowsPixels(px float64) string {
	return joinClassName(gridTemplateRows, strconv.FormatFloat(px, 'f', -1, 64))
}

func StyleClasses() []byte {
	sb := &strings.Builder{}
	for className := range maps.Keys(setClasses) {
		property, value := parsePropertyValue(className)
		sb.WriteString(classSelector(className) + "{")
		sb.WriteString(property + ":" + value + "}")
	}
	return []byte(sb.String())
}

func parsePropertyValue(className string) (string, string) {
	abbrParts := strings.Split(className, classNameSep)
	if len(abbrParts) != 2 {
		return "", ""
	}
	pfx, sfx := abbrParts[0], abbrParts[1]
	property := customProperty(pfx)
	value := ""

	switch pfx {
	case alignContentPfx:
		fallthrough
	case alignItemsPfx:
		fallthrough
	case justifyContentPfx:
		fallthrough
	case justifyItemsPfx:
		al := align.Parse(sfx)
		value = al.String()
	case fontSizePfx:
		sz := size.Parse(sfx)
		value = sz.FontSizeCssValue()
	case marginBlockEndPfx:
		fallthrough
	case columnGapPfx:
		fallthrough
	case rowGapPfx:
		sz := size.Parse(sfx)
		value = sz.SizeCssValue()
	case gridTemplateRows:
		if _, err := strconv.ParseFloat(sfx, 64); err == nil {
			value = sfx + "px"
		} else {
			sz := size.Parse(sfx)
			value = sz.SizeCssValue()
		}
	case flexDirectionPfx:
		dr := direction.Parse(sfx)
		value = dr.String()
	case markerColorPfx:
		fallthrough
	case foregroundColorPfx:
		fallthrough
	case backgroundColorPfx:
		cl := color.Parse(sfx)
		value = cl.CssValue()
	case fontWeightPfx:
		wt := font_weight.Parse(sfx)
		value = wt.CssValue()
	default:
		panic("class support not implemented for " + pfx)
	}

	return property, value
}
