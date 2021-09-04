package colorful

import (
	"math"
)

// Converts the given color to LuvLCh space using D65 as reference white.
// h values are in [0..360], C and L values are in [0..1] although C can overshoot 1.0
func (col Color) LuvLCh() (l, c, h float64) {
	return col.LuvLChWhiteRef(D65)
}

func LuvToLuvLCh(L, u, v float64) (l, c, h float64) {
	// Oops, floating point workaround necessary if u ~= v and both are very small (i.e. almost zero).
	if math.Abs(v-u) > 1e-4 && math.Abs(u) > 1e-4 {
		h = math.Mod(57.29577951308232087721*math.Atan2(v, u)+360.0, 360.0) // Rad2Deg
	} else {
		h = 0.0
	}
	l = L
	c = math.Sqrt(sq(u) + sq(v))
	return
}

// Converts the given color to LuvLCh space, taking into account
// a given reference white. (i.e. the monitor's white)
// h values are in [0..360], c and l values are in [0..1]
func (col Color) LuvLChWhiteRef(wref [3]float64) (l, c, h float64) {
	return LuvToLuvLCh(col.LuvWhiteRef(wref))
}

// Generates a color by using data given in LuvLCh space using D65 as reference white.
// h values are in [0..360], C and L values are in [0..1]
// WARNING: many combinations of `l`, `c`, and `h` values do not have corresponding
// valid RGB values, check the FAQ in the README if you're unsure.
func LuvLCh(l, c, h float64) Color {
	return LuvLChWhiteRef(l, c, h, D65)
}

func LuvLChToLuv(l, c, h float64) (L, u, v float64) {
	H := 0.01745329251994329576 * h // Deg2Rad
	u = c * math.Cos(H)
	v = c * math.Sin(H)
	L = l
	return
}

// Generates a color by using data given in LuvLCh space, taking
// into account a given reference white. (i.e. the monitor's white)
// h values are in [0..360], C and L values are in [0..1]
func LuvLChWhiteRef(l, c, h float64, wref [3]float64) Color {
	L, u, v := LuvLChToLuv(l, c, h)
	return LuvWhiteRef(L, u, v, wref)
}

// BlendLUVLCh blends the color with col2 in the cylindrical CIELUV color space.
func (col Color) BlendLUVLCh(col2 Color, t float64) Color {
	return BlendLUVLCh(col, col2, t)
}

// BlendLUVLCh blends col1 and col2 in the cylindrical CIELUV color space.
func BlendLUVLCh(col1, col2 Color, t float64) Color {
	l1, c1, h1 := col1.LuvLCh()
	l2, c2, h2 := col2.LuvLCh()

	// We know that h are both in [0..360]
	return LuvLCh(l1+t*(l2-l1), c1+t*(c2-c1), lerpAngle(h1, h2, t))
}
