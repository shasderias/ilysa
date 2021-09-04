package colorful

import "math"

// HSL returns the coordinates of the color in the HSL color model. The
// coordinates returned will fall in the following ranges.
// Hue        - [0,360]
// Saturation - [0,1]
// Luminance  - [0,1]
func (col Color) HSL() (h, s, l float64) {
	min := math.Min(math.Min(col.R, col.G), col.B)
	max := math.Max(math.Max(col.R, col.G), col.B)

	l = (max + min) / 2

	if min == max {
		s = 0
		h = 0
	} else {
		if l < 0.5 {
			s = (max - min) / (max + min)
		} else {
			s = (max - min) / (2.0 - max - min)
		}

		if max == col.R {
			h = (col.G - col.B) / (max - min)
		} else if max == col.G {
			h = 2.0 + (col.B-col.R)/(max-min)
		} else {
			h = 4.0 + (col.R-col.G)/(max-min)
		}

		h *= 60

		if h < 0 {
			h += 360
		}
	}

	return
}

// HSL creates a color from the provided HSL coordinates.
func HSL(h, s, l float64) Color {
	if s == 0 {
		return Color{l, l, l, 1.0}
	}

	var r, g, b float64
	var t1 float64
	var t2 float64
	var tr float64
	var tg float64
	var tb float64

	if l < 0.5 {
		t1 = l * (1.0 + s)
	} else {
		t1 = l + s - l*s
	}

	t2 = 2*l - t1
	h /= 360
	tr = h + 1.0/3.0
	tg = h
	tb = h - 1.0/3.0

	if tr < 0 {
		tr++
	}
	if tr > 1 {
		tr--
	}
	if tg < 0 {
		tg++
	}
	if tg > 1 {
		tg--
	}
	if tb < 0 {
		tb++
	}
	if tb > 1 {
		tb--
	}

	// Red
	if 6*tr < 1 {
		r = t2 + (t1-t2)*6*tr
	} else if 2*tr < 1 {
		r = t1
	} else if 3*tr < 2 {
		r = t2 + (t1-t2)*(2.0/3.0-tr)*6
	} else {
		r = t2
	}

	// Green
	if 6*tg < 1 {
		g = t2 + (t1-t2)*6*tg
	} else if 2*tg < 1 {
		g = t1
	} else if 3*tg < 2 {
		g = t2 + (t1-t2)*(2.0/3.0-tg)*6
	} else {
		g = t2
	}

	// Blue
	if 6*tb < 1 {
		b = t2 + (t1-t2)*6*tb
	} else if 2*tb < 1 {
		b = t1
	} else if 3*tb < 2 {
		b = t2 + (t1-t2)*(2.0/3.0-tb)*6
	} else {
		b = t2
	}

	return Color{r, g, b, 1.0}
}
