package colorful

import "math"

// DistanceRiemersma is a color distance algorithm developed by Thiadmer Riemersma.
// It uses RGB coordinates, but he claims it has similar results to CIELUV.
// This makes it both fast and accurate.
//
// Sources:
//
//     https://www.compuphase.com/cmetric.htm
//     https://github.com/lucasb-eyer/go-colorful/issues/52
func (col Color) DistanceRiemersma(c2 Color) float64 {
	rAvg := (col.R + c2.R) / 2.0
	// Deltas
	dR := col.R - c2.R
	dG := col.G - c2.G
	dB := col.B - c2.B

	return math.Sqrt((2+rAvg)*dR*dR + 4*dG*dG + (2+(1-rAvg))*dB*dB)
}

// DistanceCIEDE2000 uses the Delta E 2000 formula to calculate color
// distance. It is more expensive but more accurate than both DistanceLAB
// and DistanceCIE94.
func (col Color) DistanceCIEDE2000(cr Color) float64 {
	return col.DistanceCIEDE2000klch(cr, 1.0, 1.0, 1.0)
}
