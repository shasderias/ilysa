package fitbeat

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/lightid"
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

func NewBackLasers() context.Light          { return light.NewBasic(BackLasers, 30) }
func NewBigRings() context.Light            { return light.NewBasic(BigRings, 30) }
func NewLeftRotatingLasers() context.Light  { return light.NewBasic(LeftRotatingLasers, 8) }
func NewRightRotatingLasers() context.Light { return light.NewBasic(RightRotatingLasers, 8) }
func NewCenterLights() context.Light        { return light.NewBasic(CenterLights, 2) }

func NewBackLasersEx() context.Light {
	return light.NewComposite(BackLasers, lightid.NewSet(
		lightid.New(
			1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29,
			2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30)))
}
