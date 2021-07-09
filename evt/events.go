package evt

type Option interface {
}

type Events []Event

func NewEvents(events ...Event) Events {
	return append([]Event{}, events...)
}

func (events Events) Apply(opts ...Opt) {
	for _, e := range events {
		for _, opt := range opts {
			opt.apply(e)
		}
	}
}

type RGBLightingEvents []*RGBLighting

func (events RGBLightingEvents) Apply(opts ...RGBLightingOpt) {
	for i := range events {
		for _, opt := range opts {
			opt.applyRGBLighting(events[i])
		}
	}
}

func (events *RGBLightingEvents) Add(newEvents ...*RGBLighting) {
	*events = append(*events, newEvents...)
}
