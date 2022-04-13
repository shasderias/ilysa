package evt

func OptB(b float64) Option          { return NewFuncOpt(func(e Event) { e.SetBeat(b) }) }
func OptShiftB(b float64) Option     { return NewFuncOpt(func(e Event) { e.SetBeat(e.Beat() + b) }) }
func OptType(t Type) Option          { return NewFuncOpt(func(e Event) { e.SetType(t) }) }
func OptValue(v Value) Option        { return NewFuncOpt(func(e Event) { e.SetValue(v) }) }
func OptIntValue(v int) Option       { return NewFuncOpt(func(e Event) { e.SetValue(Value(v)) }) }
func OptFloatValue(f float64) Option { return NewFuncOpt(func(e Event) { e.SetFloatValue(f) }) }

type FuncOpt struct {
	applyFn func(Event)
}

func NewFuncOpt(applyFn func(e Event)) FuncOpt { return FuncOpt{applyFn} }
func (o FuncOpt) Apply(e Event)                { o.applyFn(e) }
