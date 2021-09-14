package main

import (
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
)

var (
	magnetRainbowPale = gradient.New(
		colorful.MustParseHex("#F48DB4"),
		colorful.MustParseHex("#BCA2D8"),
		colorful.MustParseHex("#70B5D8"),
		colorful.MustParseHex("#44BFB4"),
		colorful.MustParseHex("#6DBE81"),
		colorful.MustParseHex("#A5B559"),
		colorful.MustParseHex("#D6A454"),
		colorful.MustParseHex("#F49472"),
	)

	magnetRainbow = gradient.New(
		colorful.MustParseHex("#FF0000"),
		colorful.MustParseHex("#FF8000"),
		colorful.MustParseHex("#FFFF00"),
		colorful.MustParseHex("#00FF00"),
		colorful.MustParseHex("#00FFFF"),
		colorful.MustParseHex("#0000FF"),
		colorful.MustParseHex("#8000FF"),
		colorful.MustParseHex("#FF00FF"),
	)
)

var (
	shirayukiGold   = colorful.MustParseHex("#F5CA1C")
	shirayukiPurple = colorful.MustParseHex("#a51fce")
	sukoyaPink      = colorful.MustParseHex("#F521CF")
	sukoyaWhite     = colorful.MustParseHex("#FFFCFF")
)

var (
	magnetRed       = colorful.MustParseHex("#600F45")
	magnetPurpleRed = colorful.MustParseHex("#8A317C")
	magnetPurple    = colorful.MustParseHex("#B241BA")
	magnetPink      = colorful.MustParseHex("#C856D9")
	magnetWhite     = colorful.MustParseHex("#FFBEFF")
)

var (
	allColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
		sukoyaWhite,
		sukoyaPink,
		magnetPurple,
		magnetPink,
		magnetWhite,
	)

	magnetColors = colorful.NewSet(
		magnetRed,
		magnetPurpleRed,
		magnetPurple,
		magnetPink,
		magnetWhite,
	)

	shirayukiColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
	)

	sukoyaColors = colorful.NewSet(
		sukoyaPink,
		sukoyaWhite,
	)

	crossickColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
		sukoyaWhite,
		sukoyaPink,
	)
)

var (
	magnetGradient = gradient.Table{
		{magnetRed, 0.0},
		{magnetPurpleRed, 0.25},
		{magnetPurple, 0.50},
		{magnetPink, 0.75},
		{magnetWhite, 1.00},
	}

	allColorsGradient = gradient.Table{
		{shirayukiGold, 0.0},
		{shirayukiPurple, 0.167},
		{sukoyaPink, 0.167 * 2},
		{sukoyaWhite, 0.167 * 3},
		{magnetPurple, 0.167 * 4},
		{magnetPink, 0.167 * 5},
		{magnetWhite, 1.0},
	}

	//shirayukiGradient = gradient.Table{
	//	{shirayukiPurple, 0},
	//	{shirayukiGold, 0.33},
	//	{shirayukiPurple, 0.5},
	//	{shirayukiGold, 0.66},
	//	{shirayukiPurple, 1},
	//}

	shirayukiGradient = gradient.New(
		shirayukiPurple,
		shirayukiGold,
		shirayukiGold,
		shirayukiPurple,
	)

	shirayukiSingleGradient = gradient.New(
		shirayukiPurple,
		shirayukiGold,
	)

	sukoyaGradient = gradient.New(
		sukoyaPink,
		sukoyaWhite,
		sukoyaWhite,
		sukoyaPink,
	)

	shirayukiWhiteGradient = gradient.New(
		shirayukiPurple,
		magnetWhite,
	)

	sukoyaSingleGradient = gradient.New(
		sukoyaPink,
		sukoyaWhite,
	)

	sukoyaWhiteGradient = gradient.New(
		sukoyaPink,
		magnetWhite,
	)

	sukoyaWing           = gradient.New(sukoyaPink, sukoyaWhite, sukoyaWhite, sukoyaPink)
	sukoyaWingInverse    = gradient.New(sukoyaWhite, sukoyaPink, sukoyaPink, sukoyaWhite)
	shirayukiWing        = gradient.New(shirayukiPurple, shirayukiGold, shirayukiPurple)
	shirayukiWingInverse = gradient.New(shirayukiGold, shirayukiPurple, shirayukiGold)

	shirayukiRipple = gradient.Table{
		{shirayukiPurple, 0.0},
		{shirayukiGold, 0.3},
		{shirayukiGold, 0.7},
		{shirayukiPurple, 1.0},
	}

	sukoyaRipple = gradient.Table{
		{sukoyaPink, 0.0},
		{sukoyaWhite, 0.3},
		{sukoyaWhite, 0.7},
		{sukoyaPink, 1.0},
	}
)
