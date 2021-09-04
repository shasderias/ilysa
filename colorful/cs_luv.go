package colorful

import "math"

// http://en.wikipedia.org/wiki/CIELUV#XYZ_.E2.86.92_CIELUV_and_CIELUV_.E2.86.92_XYZ_conversions
// For L*u*v*, we need to L*u*v*<->XYZ<->RGB and the first one is device dependent.

func XYZToLuv(x, y, z float64) (l, a, b float64) {
	// Use D65 white as reference point by default.
	// http://www.fredmiranda.com/forum/topic/1035332
	// http://en.wikipedia.org/wiki/Standard_illuminant
	return XYZToLuvWhiteRef(x, y, z, D65)
}

func XYZToLuvWhiteRef(x, y, z float64, wref [3]float64) (l, u, v float64) {
	if y/wref[1] <= 6.0/29.0*6.0/29.0*6.0/29.0 {
		l = y / wref[1] * (29.0 / 3.0 * 29.0 / 3.0 * 29.0 / 3.0) / 100.0
	} else {
		l = 1.16*math.Cbrt(y/wref[1]) - 0.16
	}
	ubis, vbis := xyz_to_uv(x, y, z)
	un, vn := xyz_to_uv(wref[0], wref[1], wref[2])
	u = 13.0 * l * (ubis - un)
	v = 13.0 * l * (vbis - vn)
	return
}

// For this part, we do as R's graphics.hcl does, not as wikipedia does.
// Or is it the same?
func xyz_to_uv(x, y, z float64) (u, v float64) {
	denom := x + 15.0*y + 3.0*z
	if denom == 0.0 {
		u, v = 0.0, 0.0
	} else {
		u = 4.0 * x / denom
		v = 9.0 * y / denom
	}
	return
}

func LuvToXYZ(l, u, v float64) (x, y, z float64) {
	// D65 white (see above).
	return LuvToXYZWhiteRef(l, u, v, D65)
}

func LuvToXYZWhiteRef(l, u, v float64, wref [3]float64) (x, y, z float64) {
	//y = wref[1] * lab_finv((l + 0.16) / 1.16)
	if l <= 0.08 {
		y = wref[1] * l * 100.0 * 3.0 / 29.0 * 3.0 / 29.0 * 3.0 / 29.0
	} else {
		y = wref[1] * cub((l+0.16)/1.16)
	}
	un, vn := xyz_to_uv(wref[0], wref[1], wref[2])
	if l != 0.0 {
		ubis := u/(13.0*l) + un
		vbis := v/(13.0*l) + vn
		x = y * 9.0 * ubis / (4.0 * vbis)
		z = y * (12.0 - 3.0*ubis - 20.0*vbis) / (4.0 * vbis)
	} else {
		x, y = 0.0, 0.0
	}
	return
}

// Converts the given color to CIE L*u*v* space using D65 as reference white.
// L* is in [0..1] and both u* and v* are in about [-1..1]
func (col Color) Luv() (l, u, v float64) {
	return XYZToLuv(col.XYZ())
}

// Converts the given color to CIE L*u*v* space, taking into account
// a given reference white. (i.e. the monitor's white)
// L* is in [0..1] and both u* and v* are in about [-1..1]
func (col Color) LuvWhiteRef(wref [3]float64) (l, u, v float64) {
	x, y, z := col.XYZ()
	return XYZToLuvWhiteRef(x, y, z, wref)
}

// Generates a color by using data given in CIE L*u*v* space using D65 as reference white.
// L* is in [0..1] and both u* and v* are in about [-1..1]
// WARNING: many combinations of `l`, `u`, and `v` values do not have corresponding
// valid RGB values, check the FAQ in the README if you're unsure.
func Luv(l, u, v float64) Color {
	return XYZ(LuvToXYZ(l, u, v))
}

// Generates a color by using data given in CIE L*u*v* space, taking
// into account a given reference white. (i.e. the monitor's white)
// L* is in [0..1] and both u* and v* are in about [-1..1]
func LuvWhiteRef(l, u, v float64, wref [3]float64) Color {
	return XYZ(LuvToXYZWhiteRef(l, u, v, wref))
}

// DistanceLuv is a good measure of visual similarity between two colors!
// A result of 0 would mean identical colors, while a result of 1 or higher
// means the colors differ a lot.
func (col Color) DistanceLuv(c2 Color) float64 {
	l1, u1, v1 := col.Luv()
	l2, u2, v2 := c2.Luv()
	return math.Sqrt(sq(l1-l2) + sq(u1-u2) + sq(v1-v2))
}

// BlendLUV blends the color with c2 in the CIE-L*u*v* color space.
func (col Color) BlendLUV(c2 Color, t float64) Color {
	return BlendLUV(col, c2, t)
}

// BlendLUV blends c1 and c2 in the CIE-L*u*v* color space.
func BlendLUV(c1, c2 Color, t float64) Color {
	l1, u1, v1 := c1.Luv()
	l2, u2, v2 := c2.Luv()
	return Luv(l1+t*(l2-l1),
		u1+t*(u2-u1),
		v1+t*(v2-v1))
}
