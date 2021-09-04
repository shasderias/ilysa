package colorful

import "math"

// DistanceRGB computes the distance between two colors in RGB space.
// This is not a good measure! Rather do it in LAB space.
func (col Color) DistanceRGB(c2 Color) float64 {
	return math.Sqrt(sq(col.R-c2.R) + sq(col.G-c2.G) + sq(col.B-c2.B))
}

// BlendRGB blends the color with c2 in the sRGB color space. Prefer BlendLAB,
// BlendLUV or BlendLABLCH over this.
func (col Color) BlendRGB(c2 Color, t float64) Color {
	return BlendRGB(col, c2, t)
}

// BlendRGB blends c1 with c2 in the sRGB color space. Prefer BlendLAB,
// BlendLUV or BlendLABLCH over this.
func BlendRGB(c1, c2 Color, t float64) Color {
	return Color{c1.R + t*(c2.R-c1.R),
		c1.G + t*(c2.G-c1.G),
		c1.B + t*(c2.B-c1.B),
		c1.A + t*(c2.A-c1.A),
	}
}
