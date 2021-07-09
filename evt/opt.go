package evt

type Opt interface {
	apply(e Event)
}

func Apply(e Event, opts ...Opt) {
	for _, opt := range opts {
		opt.apply(e)
	}
}

type Opts []Opt

func NewOpts(opts ...Opt) Opts {
	return opts
}

func (o *Opts) Add(opts ...Opt) {
	*o = append(*o, opts...)
}

func (o Opts) apply(e Event) {
	for _, opt := range o {
		opt.apply(e)
	}
}

func (o Opts) applyRGBLighting(e *RGBLighting) {
	for _, opt := range o {
		lo, ok := opt.(RGBLightingOpt)
		if !ok {
			continue
		}
		lo.applyRGBLighting(e)
	}
}

func (o Opts) applyPreciseLaser(e *PreciseLaser) {
	for _, opt := range o {
		lo, ok := opt.(PreciseLaserOpt)
		if !ok {
			continue
		}
		lo.applyPreciseLaser(e)
	}
}

func (o Opts) applyPreciseRotation(e *PreciseRotation) {
	for _, opt := range o {
		lo, ok := opt.(PreciseRotationOpt)
		if !ok {
			continue
		}
		lo.applyPreciseRotation(e)
	}
}
