package colorful

// http://www.sjbrown.co.uk/2004/05/14/gamma-correct-rendering/

// XYZ returns the color's coordinates in the CIE XYZ color space.
func (col Color) XYZ() (x, y, z float64) {
	return LinearRGBToXYZ(col.LinearRGB())
}

// XYZ creates a Color from the provided CIE XYZ color space coordinates.
func XYZ(x, y, z float64) Color {
	return LinearRGB(XYZToLinearRGB(x, y, z))
}
