package ease

import "math"

var (
	pow  = math.Pow
	sqrt = math.Sqrt
	sin  = math.Sin
	cos  = math.Cos
)

const (
	pi = math.Pi
	// http: //void.heteml.jp/blog/archives/2014/05/easing_magicnumber.html
	c1 = 1.70158
	c2 = c1 * 1.525
	c3 = c1 + 1
	c4 = (2 * pi) / 3
	c5 = (2 * pi) / 4.5
)

type Func func(x float64) float64

func bounceOut(x float64) float64 {
	const n1 = 7.5625
	const d1 = 2.75

	if x < 1/d1 {
		return n1 * x * x
	} else if x < 2/d1 {
		x -= 1.5 / d1
		return n1*x*x + 0.75
	} else if x < 2.5/d1 {
		x -= 2.25 / d1
		return n1*x*x + 0.9375
	} else {
		x -= 2.625 / d1
		return n1*x*x + 0.984375
	}
}

func Linear(x float64) float64 { return x }

func InQuad(x float64) float64  { return x * x }
func OutQuad(x float64) float64 { return 1 - (1-x)*(1-x) }
func InOutQuad(x float64) float64 {
	if x < 0.5 {
		return 2 * x * x
	} else {
		return 1 - pow(-2*x+2, 2)/2
	}
}

func InCubic(x float64) float64  { return x * x * x }
func OutCubic(x float64) float64 { return 1 - pow(1-x, 3) }
func InOutCubic(x float64) float64 {
	if x < 0.5 {
		return 4 * x * x * x
	} else {
		return 1 - pow(-2*x+2, 3)/2
	}
}

func InQuart(x float64) float64  { return x * x * x * x }
func OutQuart(x float64) float64 { return 1 - pow(1-x, 4) }
func InOutQuart(x float64) float64 {
	if x < 0.5 {
		return 8 * x * x * x * x
	} else {
		return 1 - pow(-2*x+2, 4)/2
	}
}

func InQuint(x float64) float64  { return x * x * x * x * x }
func OutQuint(x float64) float64 { return 1 - pow(1-x, 5) }
func InOutQuint(x float64) float64 {
	if x < 0.5 {
		return 16 * x * x * x * x * x
	} else {
		return 1 - pow(-2*x+2, 5)/2
	}
}
func InSin(x float64) float64    { return 1 - cos((x*pi)/2) }
func OutSin(x float64) float64   { return sin((x * pi) / 2) }
func InOutSin(x float64) float64 { return -(cos(pi*x) - 1) / 2 }

func InExpo(x float64) float64 {
	if x == 0 {
		return 0
	} else {
		return pow(2, 10*x-10)
	}
}
func OutExpo(x float64) float64 {
	if x == 1 {
		return 1
	} else {
		return 1 - pow(2, -10*x)
	}
}
func InOutExpo(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	case x < 0.5:
		return pow(2, 20*x-10) / 2
	default:
		return (2 - pow(2, -20*x+10)) / 2
	}
}

func InCirc(x float64) float64  { return 1 - sqrt(1-pow(x, 2)) }
func OutCirc(x float64) float64 { return sqrt(1 - pow(x-1, 2)) }
func InOutCirc(x float64) float64 {
	if x < 0.5 {
		return (1 - sqrt(1-pow(2*x, 2))) / 2
	} else {
		return (sqrt(1-pow(-2*x+2, 2)) + 1) / 2
	}
}

func InBack(x float64) float64  { return c3*x*x*x - c1*x*x }
func OutBack(x float64) float64 { return 1 + c3*pow(x-1, 3) + c1*pow(x-1, 2) }
func InOutBack(x float64) float64 {
	if x < 0.5 {
		return (pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
	} else {
		return (pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
	}
}

func InElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return -pow(2, 10*x-10) * sin((x*10-10.75)*c4)
	}
}
func OutElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return pow(2, -10*x)*sin((x*10-0.75)*c4) + 1
	}
}
func InOutElastic(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	case x < 0.5:
		return -(pow(2, 20*x-10) * sin((20*x-11.125)*c5)) / 2
	default:
		return pow(2, -20*x+10)*sin((20*x-11.125)*c5)/2 + 1
	}
}

func InBounce(x float64) float64 {
	return 1 - bounceOut(1-x)
}
func OutBounce(x float64) float64 { return bounceOut(x) }
func InOutBounce(x float64) float64 {
	if x < 0.5 {
		return (1 - bounceOut(1-2*x)) / 2
	} else {
		return (1 + bounceOut(2*x-1)) / 2
	}
}
