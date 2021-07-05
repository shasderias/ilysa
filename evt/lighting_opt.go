package evt

import "github.com/shasderias/ilysa/beatsaber"

type LightType int

const (
	BackLasers            LightType = 0
	RingLights            LightType = 1
	LeftRotatingLasers    LightType = 2
	RightRotatingLasers   LightType = 3
	CenterLights          LightType = 4
	BoostLights           LightType = 5
	InterscopeLeftLights  LightType = 6
	InterscopeRightLights LightType = 7
)

func WithLight(l LightType) withLightOpt {
	return withLightOpt{beatsaber.EventType(l)}
}

type withLightOpt struct {
	l beatsaber.EventType
}

func (o withLightOpt) applyLighting(e *Lighting) {
	e.SetType(o.l)
}

func (o withLightOpt) applyRGBLighting(e *RGBLighting) {
	e.SetType(o.l)
}

type LightValue int

const (
	LightOff       LightValue = 0
	LightBlueOn    LightValue = 1
	LightBlueFlash LightValue = 2
	LightBlueFade  LightValue = 3
	LightUnused4   LightValue = 4
	LightRedOn     LightValue = 5
	LightRedFlash  LightValue = 6
	LightRedFade   LightValue = 7
)

func WithLightValue(v LightValue) withLightValueOpt {
	return withLightValueOpt{beatsaber.EventValue(v)}
}

type withLightValueOpt struct {
	v beatsaber.EventValue
}

func (o withLightValueOpt) apply(e Event) {
	switch le := e.(type) {
	case *Lighting:
		le.Apply(o)
	case *RGBLighting:
		le.Apply(o)
	}
}

func (o withLightValueOpt) applyLighting(e *Lighting) {
	e.SetValue(o.v)
}

func (o withLightValueOpt) applyRGBLighting(e *RGBLighting) {
	e.SetValue(o.v)
}
