package colorful

import "math"

// XYZToLinearRGB converts from CIE XYZ-space to Linear RGB space.
func XYZToLinearRGB(x, y, z float64) (r, g, b float64) {
	r = 3.2409699419045214*x - 1.5373831775700935*y - 0.49861076029300328*z
	g = -0.96924363628087983*x + 1.8759675015077207*y + 0.041555057407175613*z
	b = 0.055630079696993609*x - 0.20397695888897657*y + 1.0569715142428786*z
	return
}

func LinearRGBToXYZ(r, g, b float64) (x, y, z float64) {
	x = 0.41239079926595948*r + 0.35758433938387796*g + 0.18048078840183429*b
	y = 0.21263900587151036*r + 0.71516867876775593*g + 0.072192315360733715*b
	z = 0.019330818715591851*r + 0.11919477979462599*g + 0.95053215224966058*b
	return
}

// http://www.brucelindbloom.com/Eqn_XYZ_to_xyY.html

func XYZToXYY(X, Y, Z float64) (x, y, Yout float64) {
	return XYZToXYYWhiteRef(X, Y, Z, D65)
}

func XYZToXYYWhiteRef(X, Y, Z float64, wref [3]float64) (x, y, Yout float64) {
	Yout = Y
	N := X + Y + Z
	if math.Abs(N) < 1e-14 {
		// When we have black, Bruce Lindbloom recommends to use
		// the reference white's chromacity for x and y.
		x = wref[0] / (wref[0] + wref[1] + wref[2])
		y = wref[1] / (wref[0] + wref[1] + wref[2])
	} else {
		x = X / N
		y = Y / N
	}
	return
}

func XYYToXYZ(x, y, Y float64) (X, Yout, Z float64) {
	Yout = Y

	if -1e-14 < y && y < 1e-14 {
		X = 0.0
		Z = 0.0
	} else {
		X = Y / y * x
		Z = Y / y * (1.0 - x - y)
	}

	return
}

func LABToLABLCH(L, a, b float64) (h, c, l float64) {
	// Oops, floating point workaround necessary if a ~= b and both are very small (i.e. almost zero).
	if math.Abs(b-a) > 1e-4 && math.Abs(a) > 1e-4 {
		h = math.Mod(57.29577951308232087721*math.Atan2(b, a)+360.0, 360.0) // Rad2Deg
	} else {
		h = 0.0
	}
	c = math.Sqrt(sq(a) + sq(b))
	l = L
	return
}

func LABLCHToLAB(h, c, l float64) (L, a, b float64) {
	H := 0.01745329251994329576 * h // Deg2Rad
	a = c * math.Cos(H)
	b = c * math.Sin(H)
	L = l
	return
}
