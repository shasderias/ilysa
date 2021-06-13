package gradient

import "ilysa/pkg/colorful"

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type Table []struct {
	Col colorful.Color
	Pos float64
}

// This is the meat of the gradient computation. It returns a HCL-blend between
// the two colors around `t`.
// Note: It relies heavily on the fact that the gradient keypoints are sorted.
func (gt Table) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendOklab(c2.Col, t).Clamped()
		}
	}

	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return gt[len(gt)-1].Col
}

var Rainbow = Table{
	{colorful.MustParseHex("#9e0142"), 0.0},
	{colorful.MustParseHex("#d53e4f"), 0.1},
	{colorful.MustParseHex("#f46d43"), 0.2},
	{colorful.MustParseHex("#fdae61"), 0.3},
	{colorful.MustParseHex("#fee090"), 0.4},
	{colorful.MustParseHex("#ffffbf"), 0.5},
	{colorful.MustParseHex("#e6f598"), 0.6},
	{colorful.MustParseHex("#abdda4"), 0.7},
	{colorful.MustParseHex("#66c2a5"), 0.8},
	{colorful.MustParseHex("#3288bd"), 0.9},
	{colorful.MustParseHex("#5e4fa2"), 1.0},
}

