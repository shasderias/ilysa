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
