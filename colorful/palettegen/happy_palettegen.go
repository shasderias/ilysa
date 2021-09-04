package palettegen

import (
	"math/rand"

	"github.com/shasderias/ilysa/colorful"
)

// Uses the HSV color space to generate colors with similar S,V but distributed
// evenly along their Hue. This is fast but not always pretty.
// If you've got time to spare, use LAB (the non-fast below).
func FastHappyPalette(colorsCount int) (colors []colorful.Color) {
	colors = make([]colorful.Color, colorsCount)

	for i := 0; i < colorsCount; i++ {
		colors[i] = colorful.HSV(float64(i)*(360.0/float64(colorsCount)), 0.8+rand.Float64()*0.2, 0.65+rand.Float64()*0.2)
	}
	return
}

func HappyPalette(colorsCount int) ([]colorful.Color, error) {
	pimpy := func(l, a, b float64) bool {
		_, c, _ := colorful.LABToLABLCH(l, a, b)
		return 0.3 <= c && 0.4 <= l && l <= 0.8
	}
	return SoftPaletteEx(colorsCount, SoftPaletteSettings{pimpy, 50, true})
}
