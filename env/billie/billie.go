package billie

import "github.com/shasderias/ilysa/evt"

const (
	Water4 evt.Type = iota
	Water1
	LeftSunbeam
	RightSunbeam
	Sun
	MoonAndBoostColors
	Water2
	Water3
	ToggleRain
	SunbeamMode
	LeftBottomLasers
	RightBottomLasers
	LeftSunbeamSpeed
	RightSunbeamSpeed
)

const (
	RainOff evt.Value = iota
	RainOn
)

const (
	SunbeamModeTogether evt.Value = iota
	SunbeamModeSpread
)
