package colorful

import "math"

// http://en.wikipedia.org/wiki/LAB_color_space#CIELAB-CIEXYZ_conversions
// For L*a*b*, we need to L*a*b*<->XYZ->RGB and the first one is device dependent.

func lab_f(t float64) float64 {
	if t > 6.0/29.0*6.0/29.0*6.0/29.0 {
		return math.Cbrt(t)
	}
	return t/3.0*29.0/6.0*29.0/6.0 + 4.0/29.0
}

func XYZToLAB(x, y, z float64) (l, a, b float64) {
	// Use D65 white as reference point by default.
	// http://www.fredmiranda.com/forum/topic/1035332
	// http://en.wikipedia.org/wiki/Standard_illuminant
	return XYZToLABWhiteRef(x, y, z, D65)
}

func XYZToLABWhiteRef(x, y, z float64, wref [3]float64) (l, a, b float64) {
	fy := lab_f(y / wref[1])
	l = 1.16*fy - 0.16
	a = 5.0 * (lab_f(x/wref[0]) - fy)
	b = 2.0 * (fy - lab_f(z/wref[2]))
	return
}

func lab_finv(t float64) float64 {
	if t > 6.0/29.0 {
		return t * t * t
	}
	return 3.0 * 6.0 / 29.0 * 6.0 / 29.0 * (t - 4.0/29.0)
}

func LABToXYZ(l, a, b float64) (x, y, z float64) {
	// D65 white (see above).
	return LABToXYZWhiteRef(l, a, b, D65)
}

func LABToXYZWhiteRef(l, a, b float64, wref [3]float64) (x, y, z float64) {
	l2 := (l + 0.16) / 1.16
	x = wref[0] * lab_finv(l2+a/5.0)
	y = wref[1] * lab_finv(l2)
	z = wref[2] * lab_finv(l2-b/2.0)
	return
}

// Converts the given color to CIE L*a*b* space using D65 as reference white.
func (col Color) LAB() (l, a, b float64) {
	return XYZToLAB(col.XYZ())
}

// Converts the given color to CIE L*a*b* space, taking into account
// a given reference white. (i.e. the monitor's white)
func (col Color) LABWhiteRef(wref [3]float64) (l, a, b float64) {
	x, y, z := col.XYZ()
	return XYZToLABWhiteRef(x, y, z, wref)
}

// Generates a color by using data given in CIE L*a*b* space using D65 as reference white.
// WARNING: many combinations of `l`, `a`, and `b` values do not have corresponding
// valid RGB values, check the FAQ in the README if you're unsure.
func LAB(l, a, b float64) Color {
	return XYZ(LABToXYZ(l, a, b))
}

// Generates a color by using data given in CIE L*a*b* space, taking
// into account a given reference white. (i.e. the monitor's white)
func LABWhiteRef(l, a, b float64, wref [3]float64) Color {
	return XYZ(LABToXYZWhiteRef(l, a, b, wref))
}

// DistanceLAB is a good measure of visual similarity between two colors!
// A result of 0 would mean identical colors, while a result of 1 or higher
// means the colors differ a lot.
func (col Color) DistanceLAB(c2 Color) float64 {
	l1, a1, b1 := col.LAB()
	l2, a2, b2 := c2.LAB()
	return math.Sqrt(sq(l1-l2) + sq(a1-a2) + sq(b1-b2))
}

// DistanceCIE76 is the same as DistanceLAB.
func (col Color) DistanceCIE76(c2 Color) float64 {
	return col.DistanceLAB(c2)
}

// Uses the CIE94 formula to calculate color distance. More accurate than
// DistanceLAB, but also more work.
func (col Color) DistanceCIE94(cr Color) float64 {
	l1, a1, b1 := col.LAB()
	l2, a2, b2 := cr.LAB()

	// NOTE: Since all those formulas expect L,a,b values 100x larger than we
	//       have them in this library, we either need to adjust all constants
	//       in the formula, or convert the ranges of L,a,b before, and then
	//       scale the distances down again. The latter is less error-prone.
	l1, a1, b1 = l1*100.0, a1*100.0, b1*100.0
	l2, a2, b2 = l2*100.0, a2*100.0, b2*100.0

	kl := 1.0 // 2.0 for textiles
	kc := 1.0
	kh := 1.0
	k1 := 0.045 // 0.048 for textiles
	k2 := 0.015 // 0.014 for textiles.

	deltaL := l1 - l2
	c1 := math.Sqrt(sq(a1) + sq(b1))
	c2 := math.Sqrt(sq(a2) + sq(b2))
	deltaCab := c1 - c2

	// Not taking Sqrt here for stability, and it's unnecessary.
	deltaHab2 := sq(a1-a2) + sq(b1-b2) - sq(deltaCab)
	sl := 1.0
	sc := 1.0 + k1*c1
	sh := 1.0 + k2*c1

	vL2 := sq(deltaL / (kl * sl))
	vC2 := sq(deltaCab / (kc * sc))
	vH2 := deltaHab2 / sq(kh*sh)

	return math.Sqrt(vL2+vC2+vH2) * 0.01 // See above.
}

// DistanceCIEDE2000klch uses the Delta E 2000 formula with custom values
// for the weighting factors kL, kC, and kH.
func (col Color) DistanceCIEDE2000klch(cr Color, kl, kc, kh float64) float64 {
	l1, a1, b1 := col.LAB()
	l2, a2, b2 := cr.LAB()

	// As with CIE94, we scale up the ranges of L,a,b beforehand and scale
	// them down again afterwards.
	l1, a1, b1 = l1*100.0, a1*100.0, b1*100.0
	l2, a2, b2 = l2*100.0, a2*100.0, b2*100.0

	cab1 := math.Sqrt(sq(a1) + sq(b1))
	cab2 := math.Sqrt(sq(a2) + sq(b2))
	cabmean := (cab1 + cab2) / 2

	g := 0.5 * (1 - math.Sqrt(math.Pow(cabmean, 7)/(math.Pow(cabmean, 7)+math.Pow(25, 7))))
	ap1 := (1 + g) * a1
	ap2 := (1 + g) * a2
	cp1 := math.Sqrt(sq(ap1) + sq(b1))
	cp2 := math.Sqrt(sq(ap2) + sq(b2))

	hp1 := 0.0
	if b1 != ap1 || ap1 != 0 {
		hp1 = math.Atan2(b1, ap1)
		if hp1 < 0 {
			hp1 += math.Pi * 2
		}
		hp1 *= 180 / math.Pi
	}
	hp2 := 0.0
	if b2 != ap2 || ap2 != 0 {
		hp2 = math.Atan2(b2, ap2)
		if hp2 < 0 {
			hp2 += math.Pi * 2
		}
		hp2 *= 180 / math.Pi
	}

	deltaLp := l2 - l1
	deltaCp := cp2 - cp1
	dhp := 0.0
	cpProduct := cp1 * cp2
	if cpProduct != 0 {
		dhp = hp2 - hp1
		if dhp > 180 {
			dhp -= 360
		} else if dhp < -180 {
			dhp += 360
		}
	}
	deltaHp := 2 * math.Sqrt(cpProduct) * math.Sin(dhp/2*math.Pi/180)

	lpmean := (l1 + l2) / 2
	cpmean := (cp1 + cp2) / 2
	hpmean := hp1 + hp2
	if cpProduct != 0 {
		hpmean /= 2
		if math.Abs(hp1-hp2) > 180 {
			if hp1+hp2 < 360 {
				hpmean += 180
			} else {
				hpmean -= 180
			}
		}
	}

	t := 1 - 0.17*math.Cos((hpmean-30)*math.Pi/180) + 0.24*math.Cos(2*hpmean*math.Pi/180) + 0.32*math.Cos((3*hpmean+6)*math.Pi/180) - 0.2*math.Cos((4*hpmean-63)*math.Pi/180)
	deltaTheta := 30 * math.Exp(-sq((hpmean-275)/25))
	rc := 2 * math.Sqrt(math.Pow(cpmean, 7)/(math.Pow(cpmean, 7)+math.Pow(25, 7)))
	sl := 1 + (0.015*sq(lpmean-50))/math.Sqrt(20+sq(lpmean-50))
	sc := 1 + 0.045*cpmean
	sh := 1 + 0.015*cpmean*t
	rt := -math.Sin(2*deltaTheta*math.Pi/180) * rc

	return math.Sqrt(sq(deltaLp/(kl*sl))+sq(deltaCp/(kc*sc))+sq(deltaHp/(kh*sh))+rt*(deltaCp/(kc*sc))*(deltaHp/(kh*sh))) * 0.01
}

// BlendLAB blends the color with c2 in the L*a*b* color space.
func (col Color) BlendLAB(c2 Color, t float64) Color {
	return BlendLAB(col, c2, t)
}

// BlendLAB blends c1 and c2 in the L*a*b* color space.
func BlendLAB(c1, c2 Color, t float64) Color {
	l1, a1, b1 := c1.LAB()
	l2, a2, b2 := c2.LAB()

	return LAB(l1+t*(l2-l1),
		a1+t*(a2-a1),
		b1+t*(b2-b1))
}
