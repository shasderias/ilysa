package colorful

import "math"

// References
// http://www.sjbrown.co.uk/2004/05/14/gamma-correct-rendering/
// http://www.brucelindbloom.com/Eqn_RGB_to_XYZ.html

func linearize(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

// A much faster and still quite precise linearization using a 6th-order Taylor approximation.
// See the accompanying Jupyter notebook for derivation of the constants.
func linearizeFast(v float64) float64 {
	v1 := v - 0.5
	v2 := v1 * v1
	v3 := v2 * v1
	v4 := v2 * v2
	//v5 := v3*v2
	return -0.248750514614486 + 0.925583310193438*v + 1.16740237321695*v2 + 0.280457026598666*v3 - 0.0757991963780179*v4 //+ 0.0437040411548932*v5
}

func delinearize(v float64) float64 {
	if v <= 0.0031308 {
		return 12.92 * v
	}
	return 1.055*math.Pow(v, 1.0/2.4) - 0.055
}

func delinearizeFast(v float64) float64 {
	// This function (fractional root) is much harder to linearize, so we need to split.
	if v > 0.2 {
		v1 := v - 0.6
		v2 := v1 * v1
		v3 := v2 * v1
		v4 := v2 * v2
		v5 := v3 * v2
		return 0.442430344268235 + 0.592178981271708*v - 0.287864782562636*v2 + 0.253214392068985*v3 - 0.272557158129811*v4 + 0.325554383321718*v5
	} else if v > 0.03 {
		v1 := v - 0.115
		v2 := v1 * v1
		v3 := v2 * v1
		v4 := v2 * v2
		v5 := v3 * v2
		return 0.194915592891669 + 1.55227076330229*v - 3.93691860257828*v2 + 18.0679839248761*v3 - 101.468750302746*v4 + 632.341487393927*v5
	} else {
		v1 := v - 0.015
		v2 := v1 * v1
		v3 := v2 * v1
		v4 := v2 * v2
		v5 := v3 * v2
		// You can clearly see from the involved constants that the low-end is highly nonlinear.
		return 0.0519565234928877 + 5.09316778537561*v - 99.0338180489702*v2 + 3484.52322764895*v3 - 150028.083412663*v4 + 7168008.42971613*v5
	}
}

// LinearRGB returns the color's RGB triplet in the linear RGB color space.
func (col Color) LinearRGB() (r, g, b float64) {
	r = linearize(col.R)
	g = linearize(col.G)
	b = linearize(col.B)
	return
}

// FastLinearRGB is a faster LinearRGB that is almost as accurate.
// FastLinearRGB is only accurate for valid RGB colors, i.e. colors
// with RGB values in the [0,1] range.
func (col Color) FastLinearRGB() (r, g, b float64) {
	r = linearizeFast(col.R)
	g = linearizeFast(col.G)
	b = linearizeFast(col.B)
	return
}

// LinearRGB creates a Color using the provided linear RGB triplets.
func LinearRGB(r, g, b float64) Color {
	return Color{delinearize(r), delinearize(g), delinearize(b), 1.0}
}

// FastLinearRGB is a faster LinearRGB that is almost as accurate.
// FastLinearRGB is only accurate for valid RGB colors, i.e. colors
// with RGB values in the [0,1] range.
func FastLinearRGB(r, g, b float64) Color {
	return Color{delinearizeFast(r), delinearizeFast(g), delinearizeFast(b), 1.0}
}

// BlendLinearRGB blends the color with c2 in the linear RGB color space.
func (col Color) BlendLinearRGB(c2 Color, t float64) Color {
	return BlendLinearRGB(col, c2, t)
}

// BlendLinearRGB blends c1 and c2 in the linear RGB color space.
func BlendLinearRGB(c1, c2 Color, t float64) Color {
	r1, g1, b1 := c1.LinearRGB()
	r2, g2, b2 := c2.LinearRGB()
	return LinearRGB(
		r1+t*(r2-r1),
		g1+t*(g2-g1),
		b1+t*(b2-b1),
	)
}

// DistanceLinearRGB computes the distance between the color and c2 in the linear
// RGB color space. This is not useful for measuring how humans perceive color,
// but might be useful for other things, like dithering.
func (col Color) DistanceLinearRGB(c2 Color) float64 {
	r1, g1, b1 := col.LinearRGB()
	r2, g2, b2 := c2.LinearRGB()
	return math.Sqrt(sq(r1-r2) + sq(g1-g2) + sq(b1-b2))
}
