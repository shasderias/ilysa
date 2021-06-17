package beatsaber

type EventType int

const (
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

func (t EventType) EventType() EventType {
	return t
}

type EventTyper interface {
	EventType() EventType
}

type EventTypeSet []EventType

func NewEventTypeSet(eventTypes ...EventType) EventTypeSet {
	return append(EventTypeSet{}, eventTypes...)
}

func (s EventTypeSet) Pick(n int) EventType {
	return s[n%len(s)]
}

func (s *EventTypeSet) Add(eventTypes ...EventType) {
	for _, et := range eventTypes {
		if !s.Has(et) {
			*s = append(*s, et)
		}
	}
}

func (s EventTypeSet) Union(sets ...EventTypeSet) EventTypeSet {
	unionSet := NewEventTypeSet(s...)
	for _, set := range sets {
		unionSet.Add(set...)
	}
	return unionSet
}

func (s EventTypeSet) Has(et EventType) bool {
	for _, e := range s {
		if e == et {
			return true
		}
	}
	return false
}
