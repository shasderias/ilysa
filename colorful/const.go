package colorful

// Delta is the tolerance used when comparing colors using AlmostEqualRGB.
const Delta = 1.0 / 255.0

// D65 is the default reference white point.
var (
	D50 = [3]float64{0.96422, 1.00000, 0.82521}
	D65 = [3]float64{0.95047, 1.00000, 1.08883}
)

var (
	Red   = Hex("#FF0000")
	Green = Hex("#00FF00")
	Blue  = Hex("#0000FF")
	Black = Hex("#000000")
	White = Hex("#FFFFFF")
)
