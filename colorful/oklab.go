package colorful

import "math"

func (col Color) Oklab() (l, a, b float64) {
	l = 0.4122214708*col.R + 0.5363325363*col.G + 0.0514459929*col.B
	m := 0.2119034982*col.R + 0.6806995451*col.G + 0.1073969566*col.B
	s := 0.0883024619*col.R + 0.2817188376*col.G + 0.6299787005*col.B

	l_ := math.Cbrt(l)
	m_ := math.Cbrt(m)
	s_ := math.Cbrt(s)

	return 0.2104542553*l_ + 0.7936177850*m_ - 0.0040720468*s_,
		1.9779984951*l_ - 2.4285922050*m_ + 0.4505937099*s_,
		0.0259040371*l_ + 0.7827717662*m_ - 0.8086757660*s_
}

func Oklab(L, a, b float64) Color {
	l_ := L + 0.3963377774*a + 0.2158037573*b
	m_ := L - 0.1055613458*a - 0.0638541728*b
	s_ := L - 0.0894841775*a - 1.2914855480*b

	l := l_ * l_ * l_
	m := m_ * m_ * m_
	s := s_ * s_ * s_

	return Color{
		+4.0767416621*l - 3.3077115913*m + 0.2309699292*s,
		-1.2684380046*l + 2.6097574011*m - 0.3413193965*s,
		-0.0041960863*l - 0.7034186147*m + 1.7076147010*s,
		1.0,
	}
}

func (col Color) BlendOklab(c2 Color, t float64) Color {
	r1, g1, b1 := col.Oklab()
	r2, g2, b2 := c2.Oklab()

	a1 := col.A
	a2 := c2.A

	c := Oklab(
		r1+t*(r2-r1),
		g1+t*(g2-g1),
		b1+t*(b2-b1),
	)

	c.A = a1 + t*(a2-a1)
	return c
}

// https://www.shadertoy.com/view/ttcyRS
func (col Color) BlendOklabPlus(c2 Color, t float64) Color {
	r1, g1, b1 := col.Oklab()
	r2, g2, b2 := c2.Oklab()

	r := r1 + t*(r2-r1)
	g := g1 + t*(g2-g1)
	b := b1 + t*(b2-b1)

	mut := 1.0 + 0.2*t*(1.0-t)

	return Oklab(
		r*mut,
		g*mut,
		b*mut,
	)
}
