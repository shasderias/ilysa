package colorful

// LABLCH is nothing else than L*a*b* in cylindrical coordinates!
// (this was wrong on English wikipedia, I fixed it, let's hope the fix stays.)
// But it is widely popular since it is a "correct HSV"
// http://www.hunterlab.com/appnotes/an09_96a.pdf

// LABLCH returns the coordinates of the color in the CIE LABLCh color space
// using D65 as the white reference.
// H values are in [0..360], C and L values are in [0..1] although C can overshoot 1.0
func (col Color) LABLCH() (h, c, l float64) {
	return col.LABLCHWhiteRef(D65)
}

// LABLCHWhiteRef returns the coordinates of the color in the CIE LABLCh color
// space using wref as the white reference.
// H values are in [0..360], C and L values are in [0..1]
func (col Color) LABLCHWhiteRef(wref [3]float64) (h, c, l float64) {
	L, a, b := col.LABWhiteRef(wref)
	return LABToLABLCH(L, a, b)
}

// Generates a color by using data given in LABLCH space using D65 as reference white.
// H values are in [0..360], C and L values are in [0..1]
func LABLCH(h, c, l float64) Color {
	return LABLCHWhiteRef(h, c, l, D65)
}

// Generates a color by using data given in LABLCH space, taking
// into account a given reference white. (i.e. the monitor's white)
// H values are in [0..360], C and L values are in [0..1]
func LABLCHWhiteRef(h, c, l float64, wref [3]float64) Color {
	L, a, b := LABLCHToLAB(h, c, l)
	return LABWhiteRef(L, a, b, wref)
}

// BlendLABLCH blends the color with col2 in the CIE-L*C*h° color space.
func (col Color) BlendLABLCH(col2 Color, t float64) Color {
	return BlendLABLCH(col, col2, t)
}

// BlendLABLCH blends col1 with col2 in the CIE-L*C*h° color space.
func BlendLABLCH(col1, col2 Color, t float64) Color {
	h1, c1, l1 := col1.LABLCH()
	h2, c2, l2 := col2.LABLCH()

	// We know that h are both in [0..360]
	return LABLCH(lerpAngle(h1, h2, t), c1+t*(c2-c1), l1+t*(l2-l1)).Clamped()
}
