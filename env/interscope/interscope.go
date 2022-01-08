package interscope

import "github.com/shasderias/ilysa/evt"

const (
	TypeLeftLights      evt.Type = 6
	TypeRightLights              = 7
	TypeLowerHydraulics          = 16
	TypeRaiseHydraulics          = 17
)

const (
	ValueAllCarsNoHydraulics evt.Value = iota
	ValueAllCars
	ValueLeftCars
	ValueRightCars
	ValueFrontMostCars
	ValueFrontMiddleCars
	ValueBackMiddleCars
	ValueBackMostCars
)
