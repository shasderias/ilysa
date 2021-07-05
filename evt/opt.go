package evt

type Opt interface {
	apply(e Event)
}

func Apply(e Event, opts ...Opt) {
	for _, opt := range opts {
		opt.apply(e)
	}
}
