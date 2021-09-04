// Package gradient implements gradients that can be used with Ilysa.
package gradient

import (
	"math/rand"

	"github.com/mitchellh/copystructure"
	"github.com/shasderias/ilysa/colorful"
)

type LerpFn func(c1, c2 colorful.Color, t float64) colorful.Color

// DefaultLerpFn is the function used by Lerp to interpolate between two colors.
var DefaultLerpFn LerpFn = colorful.BlendOklab

// Table represents a gradient and contains its "points". The position
// of each keypoint must lie in the range [0,1] and points must be sorted
// in ascending order.
type Table []Point

type Point struct {
	Col colorful.Color
	Pos float64
}

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

	registerGrad(table)

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

// FromSet creates a gradient from a colorful.Set.
func FromSet(s colorful.Set) Table {
	return New(s.Colors()...)
}

// Lerp returns interpolated color at t. By default, the interpolation is done
// in the Oklab colorspace. Change DefaultLerpFn to change the default. Use
// LerpIn to interpolate in a specific colorspace.
func (gt Table) Lerp(t float64) colorful.Color {
	return gt.LerpIn(t, DefaultLerpFn)
}

// LerpIn returns the interpolated color at t, with the interpolation done
// by fn.
func (gt Table) LerpIn(t float64, fn LerpFn) colorful.Color {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return fn(c1.Col, c2.Col, t).Clamped()
		}
	}

	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return gt[len(gt)-1].Col
}

// Reverse returns a copy of the gradient with its colors' order reversed.
func (gt Table) Reverse() Table {
	reversedTable := copystructure.Must(copystructure.Copy(gt))
	rt := reversedTable.(Table)

	for i := len(rt)/2 - 1; i >= 0; i-- {
		opp := len(rt) - 1 - i
		rt[i].Col, rt[opp].Col = rt[opp].Col, rt[i].Col
	}

	return rt
}

// Rotate returns a copy of the gradient with its colors rotated to the left
// by n.
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

// RotateRand returns a new gradient rotated by a random number.
func (gt Table) RotateRand() Table {
	n := rand.Intn(len(gt))
	return gt.Rotate(n)
}

// ToSet returns the gradient's colors as a colorful.Set.
//
// Deprecated: use Set() instead.
func (gt Table) ToSet() colorful.Set {
	return gt.Set()
}

// Set returns the gradient's colors as a colorful.Set.
func (gt Table) Set() colorful.Set {
	colors := make([]colorful.Color, len(gt))
	for i := 0; i < len(gt); i++ {
		colors[i] = gt[i].Col
	}
	return colorful.NewSet(colors...)
}

// Colors returns the gradient's colors as a slice of colorful.Colors.
func (gt Table) Colors() []colorful.Color {
	return gt.Set().Colors()
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
