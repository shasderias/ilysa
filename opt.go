package ilysa

//import "github.com/shasderias/ilysa/evt"
//
//type Opt interface {
//	apply(e evt.Event)
//}
//
//func Apply(e evt.Event, opts ...Opt) {
//	for _, opt := range opts {
//		opt.Apply(e)
//	}
//}
//
//type Opts []Opt
//
//func NewOpts(opts ...Opt) Opts {
//	return opts
//}
//
//func (o *Opts) Add(opts ...Opt) {
//	*o = append(*o, opts...)
//}
//
//func (o Opts) apply(e evt.Event) {
//	for _, opt := range o {
//		opt.Apply(e)
//	}
//}
