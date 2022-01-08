package nice

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
)

const (
	BackLasers evt.Type = iota
	BigRings
	LeftRotatingLasers
	RightRotatingLasers
	CenterLights
	BoostLights
	_
	_
	RingRotation
	RingZoom
	_
	_
	LeftLaserSpeed
	RightLaserSpeed
)

func NewBackLasers() light.Basic          { return light.NewBasic(BackLasers, 8) }
func NewBigRings() light.Basic            { return light.NewBasic(BigRings, 40) }
func NewLeftRotatingLasers() light.Basic  { return light.NewBasic(LeftRotatingLasers, 7) }
func NewRightRotatingLasers() light.Basic { return light.NewBasic(RightRotatingLasers, 7) }
func NewCenterLights() light.Basic        { return light.NewBasic(CenterLights, 12) }

func NewBigRingsStep45() light.Custom {
	return light.NewCustomFromLightIDs(BigRings,
		1, 11, 18, 28, 33,
		7, 14, 24, 29, 39,
		3, 10, 20, 25, 35,
		6, 16, 21, 31, 38,
		2, 12, 17, 27, 34,
		8, 13, 23, 30, 40,
		4, 9, 19, 26, 36,
		5, 15, 22, 32, 37,
	)
}
