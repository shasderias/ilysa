package colorful

// Converts the given color to CIE xyY space using D65 as reference white.
// (Note that the reference white is only used for black input.)
// x, y and Y are in [0..1]

// XYY returns the coordinates of the color in the CIE xyY color space.
func (col Color) XYY() (x, y, Y float64) {
	return XYZToXYY(col.XYZ())
}

// XYYWhiteRef returns the coordinates of the color in the CIE xyY color
// space, calculated using the white reference wref.
func (col Color) XYYWhiteRef(wRef [3]float64) (x, y, Y float64) {
	X, Y2, Z := col.XYZ()
	return XYZToXYYWhiteRef(X, Y2, Z, wRef)
}

// XYY creates a Color with the provided CIE xyY color space coordinates.
func XYY(x, y, Y float64) Color {
	return XYZ(XYYToXYZ(x, y, Y))
}
