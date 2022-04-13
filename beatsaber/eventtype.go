package beatsaber

// DEPRECATED: use the environment specific event types instead
type EventType int

const (
	EventTypeInvalid                          EventType = -1
	EventTypeBackLasers                       EventType = 0
	EventTypeRingLights                       EventType = 1
	EventTypeLeftRotatingLasers               EventType = 2
	EventTypeRightRotatingLasers              EventType = 3
	EventTypeCenterLights                     EventType = 4
	EventTypeBoostLights                      EventType = 5
	EventTypeInterscopeLeftLights             EventType = 6
	EventTypeInterscopeRightLights            EventType = 7
	EventTypeRingSpin                         EventType = 8
	EventTypeRingZoom                         EventType = 9
	EventTypeBPMChange                        EventType = 10
	EventTypeUnused11                         EventType = 11
	EventTypeLeftRotatingLasersRotationSpeed  EventType = 12
	EventTypeRightRotatingLasersRotationSpeed EventType = 13
	EventTypeEarlyRotation                    EventType = 14
	EventTypeLateRotation                     EventType = 15
	EventTypeInterscopeLowerHydraulics        EventType = 16
	EventTypeInterscopeRaiseHydraulics        EventType = 17
)
