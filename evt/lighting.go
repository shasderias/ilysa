package evt

import "encoding/json"

type Lighting struct {
	Base
}

type LightingOpt interface {
	applyLighting(event *Lighting)
}

func NewLighting(opts ...LightingOpt) Lighting {
	e := Lighting{NewBase(WithInvalidDefaults())}
	for _, opt := range opts {
		opt.applyLighting(&e)
	}
	return e
}

func (e *Lighting) CustomData() (json.RawMessage, error) { return nil, nil }

func (e *Lighting) Apply(opts ...LightingOpt) {
	for _, opt := range opts {
		opt.applyLighting(e)
	}
}
