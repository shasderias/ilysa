// Package ease implements Robert Penner's easing functions.
//
// http://robertpenner.com/easing/
//
// See https://easings.net/ for visualizations of the easing functions.
package ease

import (
	"math"
)

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

var Linear Func = func(x float64) float64 { return x }
var Step Func = func(x float64) float64 { return math.Floor(x) }

var InQuad Func = func(x float64) float64 { return x * x }
var OutQuad Func = func(x float64) float64 { return 1 - (1-x)*(1-x) }
var InOutQuad Func = func(x float64) float64 {
	if x < 0.5 {
		return 2 * x * x
	} else {
		return 1 - pow(-2*x+2, 2)/2
	}
}

var InCubic Func = func(x float64) float64 { return x * x * x }
var OutCubic Func = func(x float64) float64 { return 1 - pow(1-x, 3) }
var InOutCubic Func = func(x float64) float64 {
	if x < 0.5 {
		return 4 * x * x * x
	} else {
		return 1 - pow(-2*x+2, 3)/2
	}
}

var InQuart Func = func(x float64) float64 { return x * x * x * x }
var OutQuart Func = func(x float64) float64 { return 1 - pow(1-x, 4) }
var InOutQuart Func = func(x float64) float64 {
	if x < 0.5 {
		return 8 * x * x * x * x
	} else {
		return 1 - pow(-2*x+2, 4)/2
	}
}

var InQuint Func = func(x float64) float64 { return x * x * x * x * x }
var OutQuint Func = func(x float64) float64 { return 1 - pow(1-x, 5) }
var InOutQuint Func = func(x float64) float64 {
	if x < 0.5 {
		return 16 * x * x * x * x * x
	} else {
		return 1 - pow(-2*x+2, 5)/2
	}
}
var InSin Func = func(x float64) float64 { return 1 - cos((x*pi)/2) }
var OutSin Func = func(x float64) float64 { return sin((x * pi) / 2) }
var InOutSin Func = func(x float64) float64 { return -(cos(pi*x) - 1) / 2 }

var InExpo Func = func(x float64) float64 {
	if x == 0 {
		return 0
	} else {
		return pow(2, 10*x-10)
	}
}
var OutExpo Func = func(x float64) float64 {
	if x == 1 {
		return 1
	} else {
		return 1 - pow(2, -10*x)
	}
}
var InOutExpo Func = func(x float64) float64 {
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

var InCirc Func = func(x float64) float64 { return 1 - sqrt(1-pow(x, 2)) }
var OutCirc Func = func(x float64) float64 { return sqrt(1 - pow(x-1, 2)) }
var InOutCirc Func = func(x float64) float64 {
	if x < 0.5 {
		return (1 - sqrt(1-pow(2*x, 2))) / 2
	} else {
		return (sqrt(1-pow(-2*x+2, 2)) + 1) / 2
	}
}

var InBack Func = func(x float64) float64 { return c3*x*x*x - c1*x*x }
var OutBack Func = func(x float64) float64 { return 1 + c3*pow(x-1, 3) + c1*pow(x-1, 2) }
var InOutBack Func = func(x float64) float64 {
	if x < 0.5 {
		return (pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
	} else {
		return (pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
	}
}

var InElastic Func = func(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return -pow(2, 10*x-10) * sin((x*10-10.75)*c4)
	}
}
var OutElastic Func = func(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return pow(2, -10*x)*sin((x*10-0.75)*c4) + 1
	}
}
var InOutElastic Func = func(x float64) float64 {
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

var InBounce Func = func(x float64) float64 {
	return 1 - bounceOut(1-x)
}
var OutBounce Func = func(x float64) float64 { return bounceOut(x) }
var InOutBounce Func = func(x float64) float64 {
	if x < 0.5 {
		return (1 - bounceOut(1-2*x)) / 2
	} else {
		return (1 + bounceOut(2*x-1)) / 2
	}
}
