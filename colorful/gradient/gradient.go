package gradient

import (
	"math/rand"

	"github.com/mitchellh/copystructure"
	"github.com/shasderias/ilysa/colorful"
)

// New returns a gradient with colors equidistant from each other.
func New(colors ...colorful.Color) Table {
	if len(colors) < 2 {
		panic("gradient.New(): gradient must contain at least two colors")
	}

	table := Table{}

	for i, color := range colors {
		table = append(table, Point{
			Col: color,
			Pos: float64(i) / float64(len(colors)-1),
		})
	}

	return table
}

// NewPingPong returns a gradient consisting of:
// (1) colors in order; followed by
// (2) count copies of colors, alternating between:
//   (a) colors in reverse order, with the first color omitted; and
//   (b) colors in order, with the first color omitted.
func NewPingPong(count int, colors ...colorful.Color) Table {
	if count < 1 {
		panic("gradient.NewPingPong(): count must be 1 or greater")
	}
	if len(colors) < 2 {
		panic("gradient.NewPingPong(): at least 2 colors must be provided")
	}

	reverseColors := make([]colorful.Color, len(colors)-1)

	for i := 0; i < len(reverseColors); i++ {
		reverseColors[i] = colors[len(colors)-2-i]
	}

	gradColors := append([]colorful.Color{}, colors...)

	colors = colors[1:]

	for i := 0; i < count; i++ {
		if i%2 == 0 {
			gradColors = append(gradColors, reverseColors...)
		} else {
			gradColors = append(gradColors, colors...)
		}
	}

	return New(gradColors...)
}

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type Table []Point

type Point struct {
	Col colorful.Color
	Pos float64
}

// Lerp returns interpolated color at t. The interpolation is done in the Oklab
// colorspace
func (gt Table) Lerp(t float64) colorful.Color {
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

func FromSet(s colorful.Set) Table {
	return New(s.Colors()...)
}

func (gt Table) Reverse() Table {
	reversedTable := copystructure.Must(copystructure.Copy(gt))
	rt := reversedTable.(Table)

	for i := len(rt)/2 - 1; i >= 0; i-- {
		opp := len(rt) - 1 - i
		rt[i].Col, rt[opp].Col = rt[opp].Col, rt[i].Col
	}

	return rt
}

func (gt Table) Rotate(n int) Table {
	rotatedTable := copystructure.Must(copystructure.Copy(gt))
	rt := rotatedTable.(Table)

	n %= len(gt)

	rt = append(rt, rt[0:n]...)
	rt = rt[n:]

	for i := range rt {
		rt[i].Pos = gt[i].Pos
	}

	return rt
}

func (gt Table) RotateRand() Table {
	n := rand.Intn(len(gt))
	return gt.Rotate(n)
}

func (gt Table) ToSet() colorful.Set {
	colors := make([]colorful.Color, len(gt))
	for i := 0; i < len(gt); i++ {
		colors[i] = gt[i].Col
	}
	return colorful.NewSet(colors...)
}

func (gt Table) Colors() []colorful.Color {
	return gt.ToSet().Colors()
}

// Boost returns a new gradient with each color's RGB value multiplied by m.
// Boost does not affect a color's A or position.
func (gt Table) Boost(m float64) Table {
	newTable := make(Table, 0)

	for _, c := range gt {
		newTable = append(newTable, Point{Col: c.Col.Boost(m), Pos: c.Pos})
	}

	return newTable
}
