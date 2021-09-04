// Package colorful implements functions for working with colors.
//
// colorful is adapted from github.com/lucasb-eyer/go-colorful.
// colorful has been adapted to favor brevity (e.g. certain functions panic
// instead of returning errors) and is probably not suited for general usage.
package colorful

import (
	"image/color"
	"math"
)

// Color holds the RGBA values of a color.
type Color struct {
	R, G, B, A float64
}

// RGBA implements color.Color.
func (col Color) RGBA() (r, g, b, a uint32) {
	rte := math.RoundToEven
	r = uint32(rte(col.R * 0xffff))
	g = uint32(rte(col.G * 0xffff))
	b = uint32(rte(col.B * 0xffff))
	a = uint32(rte(col.A * 0xffff))
	return
}

// FromColor constructs a colorful.Color from a color.Color
func FromColor(col color.Color) Color {
	c, ok := col.(Color)
	if ok {
		return c
	}

	r, g, b, a := col.RGBA()

	return Color{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff, float64(a) / 0xffff}
}

// RGB255 returns the RGB values of the color as uint8s.
func (col Color) RGB255() (r, g, b uint8) {
	r = uint8(col.R*255.0 + 0.5)
	g = uint8(col.G*255.0 + 0.5)
	b = uint8(col.B*255.0 + 0.5)
	return
}

// IsValid returns true if the color exists in RGB space, i.e. all values lie
// in the range [0,1].
func (col Color) IsValid() bool {
	return 0.0 <= col.R && col.R <= 1.0 &&
		0.0 <= col.G && col.G <= 1.0 &&
		0.0 <= col.B && col.B <= 1.0
}

// Clamped returns a copy of the color with its RGBA values clamped to the range [0,1].
func (col Color) Clamped() Color {
	return Color{clamp01(col.R), clamp01(col.G), clamp01(col.B), clamp01(col.A)}
}

// AlmostEqualRGB returns true if the colors are equal within the tolerance Delta.
func (col Color) AlmostEqualRGB(c2 Color) bool {
	return math.Abs(col.R-c2.R)+
		math.Abs(col.G-c2.G)+
		math.Abs(col.B-c2.B) < 3.0*Delta
}

// Used to simplify HSLuv testing.
func (col Color) values() (float64, float64, float64) {
	return col.R, col.G, col.B
}
