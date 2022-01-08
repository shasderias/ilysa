package evt

type (
	Type  int
	Value int
)

const (
	TypeInvalid         Type = -1
	TypeBackLasers           = 0
	TypeRingLights           = 1
	TypeLeftLasers           = 2
	TypeRightLasers          = 3
	TypeCenterLights         = 4
	TypeBoostLights          = 5
	TypeRingRotation         = 8
	TypeRingZoom             = 9
	TypeBPMChange            = 10
	TypeLeftLaserSpeed       = 12
	TypeRightLaserSpeed      = 13
	TypeEarlyRotation        = 14
	TypeLateRotation         = 15
)

const (
	ValueInvalid        Value = -1
	ValueLightOff             = 0
	ValueLightBlueOn          = 1
	ValueLightBlueFlash       = 2
	ValueLightBlueFade        = 3
	ValueLightUnused4         = 4
	ValueLightRedOn           = 5
	ValueLightRedFlash        = 6
	ValueLightRedFade         = 7
)

const (
	ValueBoostOff Value = 0
	ValueBoostOn        = 1
)

const (
	Value60CCW Value = iota
	Value45CCW
	Value30CCW
	Value15CCW
	Value15CW
	Value30CW
	Value45CW
	Value60CW
)
