package colorful

import "math"

// From http://en.wikipedia.org/wiki/HSL_and_HSV
// Note that h is in [0..360] and s,v in [0..1]

// HSV returns the Hue [0..360], Saturation and Value [0..1] of the color.
func (col Color) HSV() (h, s, v float64) {
	min := math.Min(math.Min(col.R, col.G), col.B)
	v = math.Max(math.Max(col.R, col.G), col.B)
	C := v - min

	s = 0.0
	if v != 0.0 {
		s = C / v
	}

	h = 0.0 // We use 0 instead of undefined as in wp.
	if min != v {
		if v == col.R {
			h = math.Mod((col.G-col.B)/C, 6.0)
		}
		if v == col.G {
			h = (col.B-col.R)/C + 2.0
		}
		if v == col.B {
			h = (col.R-col.G)/C + 4.0
		}
		h *= 60.0
		if h < 0.0 {
			h += 360.0
		}
	}
	return
}

// HSV creates a new Color given a Hue in [0..360], a Saturation and a Value in [0..1]
func HSV(H, S, V float64) Color {
	Hp := H / 60.0
	C := V * S
	X := C * (1.0 - math.Abs(math.Mod(Hp, 2.0)-1.0))

	m := V - C
	r, g, b := 0.0, 0.0, 0.0

	switch {
	case 0.0 <= Hp && Hp < 1.0:
		r = C
		g = X
	case 1.0 <= Hp && Hp < 2.0:
		r = X
		g = C
	case 2.0 <= Hp && Hp < 3.0:
		g = C
		b = X
	case 3.0 <= Hp && Hp < 4.0:
		g = X
		b = C
	case 4.0 <= Hp && Hp < 5.0:
		r = X
		b = C
	case 5.0 <= Hp && Hp < 6.0:
		r = C
		b = X
	}

	return Color{m + r, m + g, m + b, 1.0}
}

// BlendHSV blends the color with c2 using the HSV color model.
func (col Color) BlendHSV(c2 Color, t float64) Color {
	return BlendHSV(col, c2, t)
}

// BlendHSV blends c1 with c2 using the HSV color model.
func BlendHSV(c1, c2 Color, t float64) Color {
	h1, s1, v1 := c1.HSV()
	h2, s2, v2 := c2.HSV()

	// We know that h are both in [0..360]
	return HSV(lerpAngle(h1, h2, t), s1+t*(s2-s1), v1+t*(v2-v1))
}
