// Package opt provides handy helper methods for options.
package opt

import "github.com/shasderias/ilysa/evt"

type alpha interface{ SetAlpha(a float64) }
type alphaOpt struct{ a float64 }

func (o alphaOpt) Apply(e evt.Event) {
	if a, ok := e.(alpha); ok {
		a.SetAlpha(o.a)
	}
}

func Alpha(a float64) evt.Option {
	return alphaOpt{a}
}

type Set []evt.Option

func NewSet(opts ...evt.Option) *Set {
	set := Set{}
	set = append(Set{}, opts...)
	return &set
}

func (s *Set) Add(opts ...evt.Option) *Set {
	*s = append(*s, opts...)
	return s
}

func (s *Set) Apply(e evt.Event) {
	for _, o := range *s {
		o.Apply(e)
	}
}

//func Successive(ctx context.Context, opts ...Option) Option {
//	return opts[ctx.Ordinal()%len(opts)]
//}
//
//type beatOpt struct {
//	b float64
//}
//
//type beatEvt interface {
//	SetBeat(b float64)
//}
//
//func Beat(b float64) evt.Option {
//	return beatOpt{b}
//}
//
//func (o beatOpt) Apply(e evt.Event) {
//	be, ok := e.(beatEvt)
//	if ok {
//		be.SetBeat(o.b)
//	}
//}
//
//type BaseOpt interface {
//	applyBase(e *Base)
//}
//
//// WithBeat sets the event's beat to b.
//// WithBeat can be used with all events.
//func WithBeat(b float64) Opt {
//	return withBeatOpt{b}
//}
//
//type withBeatOpt struct{ b float64 }
//
//func (o withBeatOpt) applyBase(e *Base) {
//	e.SetBeat(o.b)
//}
//func (o withBeatOpt) apply(e Event) {
//	e.SetBeat(o.b)
//}
//
//// WithB is an alias for WithBeat.
//func WithB(b float64) Opt {
//	return WithBeat(b)
//}
//
//// WithBOffset adds o to the event's beat.
//// WithBOffset can be used with all events.
//func WithBOffset(o float64) Opt {
//	return withBeatOffsetOpt{o}
//}
//
//type withBeatOffsetOpt struct {
//	o float64
//}
//
//func (o withBeatOffsetOpt) applyBase(e *Base) {
//	e.SetBeat(e.Beat() + o.o)
//}
//func (o withBeatOffsetOpt) apply(e Event) {
//	e.SetBeat(e.Beat() + o.o)
//}
//
//// WithTag tags the event with tags.
//// WithTag can be used with all events.
//func WithTag(tag ...string) Opt {
//	return withTagOpt{tag}
//}
//
//type withTagOpt struct {
//	tag []string
//}
//
//func (o withTagOpt) applyBase(e *Base) {
//	for _, t := range o.tag {
//		e.SetTag(t)
//	}
//}
//func (o withTagOpt) apply(e Event) {
//	for _, t := range o.tag {
//		e.SetTag(t)
//	}
//}
//
//// WithType sets the event's type.
//// WithType can be used with all events.
//func WithType(t beatsaber.EventType) Opt {
//	return withTypeOpt{t}
//}
//
//type withTypeOpt struct{ t beatsaber.EventType }
//
//func (o withTypeOpt) applyBase(e *Base) {
//	e.SetType(o.t)
//}
//func (o withTypeOpt) apply(e Event) {
//	e.SetType(o.t)
//}
//
//// WithValue sets the event's value.
//// WithValue can be used with all events.
//func WithValue(t beatsaber.EventValue) Opt {
//	return withValueOpt{t}
//}
//
//type withValueOpt struct{ t beatsaber.EventValue }
//
//func (o withValueOpt) applyBase(e *Base) {
//	e.SetValue(o.t)
//}
//func (o withValueOpt) apply(e Event) {
//	e.SetValue(o.t)
//}
//
//// WithIntValue sets the value of the event to the integer v.
//// WithIntValue can be used with:
////   - NewLaserSpeed()
////   - NewChromaLaserSpeed()
//func WithIntValue(v int) Opt {
//	return withIntValueOpt{v}
//}
//
//type withIntValueOpt struct{ v int }
//
//func (o withIntValueOpt) applyBase(e *Base) {
//	e.SetValue(beatsaber.EventValue(o.v))
//}
//func (o withIntValueOpt) apply(e Event) {
//	switch l := e.(type) {
//	case *Laser:
//		l.SetValue(beatsaber.EventValue(o.v))
//	case *PreciseLaser:
//		l.SetValue(beatsaber.EventValue(o.v))
//	}
//}
//
//type withInvalidBaseOpt struct{}
//
//func WithInvalidDefaults() Opt {
//	return withInvalidBaseOpt{}
//}
//
//func (o withInvalidBaseOpt) applyBase(e *Base) {
//	e.SetBeat(-1)
//	e.SetType(beatsaber.EventTypeInvalid)
//	e.SetValue(beatsaber.EventValueInvalid)
//}
//func (o withInvalidBaseOpt) apply(e Event) {
//	e.SetBeat(-1)
//	e.SetType(beatsaber.EventTypeInvalid)
//	e.SetValue(beatsaber.EventValueInvalid)
//}
//
//type withRGBLightingDefault struct{}
//
//func WithRGBLightingDefaults() Opt {
//	return withRGBLightingDefault{}
//}
//
//func (o withRGBLightingDefault) applyBase(e *Base) {
//	e.SetBeat(-1)
//	e.SetType(beatsaber.EventTypeInvalid)
//	e.SetValue(beatsaber.EventValueLightRedOn)
//}
//func (o withRGBLightingDefault) apply(e Event) {
//	e.SetBeat(-1)
//	e.SetType(beatsaber.EventTypeInvalid)
//	e.SetValue(beatsaber.EventValueLightRedOn)
//
//}
//func (o withLightOpt) applyChromaGradient(g *ChromaGradient) {
//	g.SetType(o.l)
//}
//
//type withDurationOpt struct {
//	d float64
//}
//
//// WithDuration sets the duration of a Chroma 2.0 gradient.
//// WithDuration can be used with:
////   - NewChromaGradient()
//func WithDuration(d float64) Option {
//	return withDurationOpt{d: d}
//}
//
//func (o withDurationOpt) apply(e Event) {
//	switch cg := e.(type) {
//	case *ChromaGradient:
//		cg.LightGradient.Duration = o.d
//	}
//}
//
//type withStartColorOpt struct {
//	c color.Color
//}
//
//// WithStartColor sets the initial color of a Chroma 2.0 gradient.
//// WithStartColor can be used with:
////   - NewChromaGradient()
//func WithStartColor(c color.Color) Option {
//	return withStartColorOpt{c}
//}
//
//func (o withStartColorOpt) apply(e Event) {
//	switch cg := e.(type) {
//	case *ChromaGradient:
//		cg.LightGradient.StartColor = o.c
//	}
//}
//
//type withEndColorOpt struct {
//	c color.Color
//}
//
//// WithEndColor sets the final color of a Chroma 2.0 gradient.
//// WithEndColor can be used with:
////   - NewChromaGradient()
//func WithEndColor(c color.Color) Option {
//	return withEndColorOpt{c}
//}
//
//func (o withEndColorOpt) apply(e Event) {
//	switch cg := e.(type) {
//	case *ChromaGradient:
//		cg.LightGradient.EndColor = o.c
//	}
//}
//
//type withEasingOpt struct {
//	e string
//}
//
//func nameOf(f interface{}) string {
//	v := reflect.ValueOf(f)
//	if v.Kind() == reflect.Func {
//		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
//			return rf.Name()
//		}
//	}
//	return v.String()
//}
//
//// WithEasing sets the easing function used to ease the Chroma 2.0 gradient.
//// easeFn must be a function from the ease package. Custom easing functions
//// will not work.
//// WithDuration can be used with:
////   - NewChromaGradient()
//func WithEasing(easeFn ease.Func) Option {
//	easeFnName := nameOf(easeFn)
//
//	switch {
//	case strings.Contains(easeFnName, "ease.Linear"):
//		easeFnName = "easeLinear"
//	case strings.Contains(easeFnName, "ease.Step"):
//		easeFnName = "easeStep"
//	case strings.Contains(easeFnName, "ease.InQuad"):
//		easeFnName = "easeInQuad"
//	case strings.Contains(easeFnName, "ease.OutQuad"):
//		easeFnName = "easeOutQuad"
//	case strings.Contains(easeFnName, "ease.InOutQuad"):
//		easeFnName = "easeInOutQuad"
//	case strings.Contains(easeFnName, "ease.InCubic"):
//		easeFnName = "easeInCubic"
//	case strings.Contains(easeFnName, "ease.OutCubic"):
//		easeFnName = "easeOutCubic"
//	case strings.Contains(easeFnName, "ease.InOutCubic"):
//		easeFnName = "easeInOutCubic"
//	case strings.Contains(easeFnName, "ease.InQuart"):
//		easeFnName = "easeInQuart"
//	case strings.Contains(easeFnName, "ease.OutQuart"):
//		easeFnName = "easeOutQuart"
//	case strings.Contains(easeFnName, "ease.InOutQuart"):
//		easeFnName = "easeInOutQuart"
//	case strings.Contains(easeFnName, "ease.InQuint"):
//		easeFnName = "easeInQuint"
//	case strings.Contains(easeFnName, "ease.OutQuint"):
//		easeFnName = "easeOutQuint"
//	case strings.Contains(easeFnName, "ease.InOutQuint"):
//		easeFnName = "easeInOutQuint"
//	case strings.Contains(easeFnName, "ease.InSin"):
//		easeFnName = "easeInSine"
//	case strings.Contains(easeFnName, "ease.OutSin"):
//		easeFnName = "easeOutSine"
//	case strings.Contains(easeFnName, "ease.InOutSin"):
//		easeFnName = "easeInOutSine"
//	case strings.Contains(easeFnName, "ease.InExpo"):
//		easeFnName = "easeInExpo"
//	case strings.Contains(easeFnName, "ease.OutExpo"):
//		easeFnName = "easeOutExpo"
//	case strings.Contains(easeFnName, "ease.InOutExpo"):
//		easeFnName = "easeInOutExpo"
//	case strings.Contains(easeFnName, "ease.InCirc"):
//		easeFnName = "easeInCirc"
//	case strings.Contains(easeFnName, "ease.OutCirc"):
//		easeFnName = "easeOutCirc"
//	case strings.Contains(easeFnName, "ease.InOutCirc"):
//		easeFnName = "easeInOutCirc"
//	case strings.Contains(easeFnName, "ease.InBack"):
//		easeFnName = "easeInBack"
//	case strings.Contains(easeFnName, "ease.OutBack"):
//		easeFnName = "easeOutBack"
//	case strings.Contains(easeFnName, "ease.InOutBack"):
//		easeFnName = "easeInOutBack"
//	case strings.Contains(easeFnName, "ease.InElastic"):
//		easeFnName = "easeInElastic"
//	case strings.Contains(easeFnName, "ease.OutElastic"):
//		easeFnName = "easeOutElastic"
//	case strings.Contains(easeFnName, "ease.InOutElastic"):
//		easeFnName = "easeInOutElastic"
//	case strings.Contains(easeFnName, "ease.InBounce"):
//		easeFnName = "easeInBounce"
//	case strings.Contains(easeFnName, "ease.OutBounce"):
//		easeFnName = "easeOutBounce"
//	case strings.Contains(easeFnName, "ease.InOutBounce"):
//		easeFnName = "easeInOutBounce"
//	default:
//		panic(fmt.Sprintf("WithEasing: unsupported ease: %v", easeFn))
//	}
//
//	return withEasingOpt{e: easeFnName}
//}
//
//func (o withEasingOpt) apply(e Event) {
//	switch cg := e.(type) {
//	case *ChromaGradient:
//		cg.LightGradient.Easing = o.e
//	}
//}
//
//// WithZoomStep dictates how much position offset is added between each ring
//// WithZoomStep accepts the following options:
//func WithZoomStep(s float64) withZoomStepOpt {
//	return withZoomStepOpt{s}
//}
//
//type withZoomStepOpt struct {
//	s float64
//}
//
//func (o withZoomStepOpt) apply(e Event) {
//	switch z := e.(type) {
//	case *PreciseZoom:
//		z.Step = o.s
//	}
//}
//
//// WithNameFilter causes event to only affect rings with the name filter (e.g.
//// SmallTrackLaneRings, BigTrackLaneRings).
//// WithNameFilter can be used with:
////   - NewChromaLighting()
//func WithNameFilter(filter string) Option {
//	return withNameFilterOpt{filter}
//}
//
//type withNameFilterOpt struct {
//	nameFilter string
//}
//
//func (o withNameFilterOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.NameFilter = o.nameFilter
//	}
//}
//
//// WithReset resets the rings when set to true (overwrites other values below)
//// WithReset can be used with:
////   - NewChromaLighting()
//func WithReset(r bool) Option {
//	return withResetOpt{r}
//}
//
//type withResetOpt struct {
//	reset bool
//}
//
//func (o withResetOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.Reset = o.reset
//	}
//}
//
//// WithRotation dictates how far the first ring will spin
//// WithRotation can be used with:
////   - NewChromaLighting()
//func WithRotation(r float64) Option {
//	return withRotationOpt{r}
//}
//
//type withRotationOpt struct {
//	rotation float64
//}
//
//func (o withRotationOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.Rotation = o.rotation
//	}
//}
//
//// WithRotationStep dictates how much rotation is added between each ring
//// WithRotationStep can be used with:
////   - NewChromaLighting()
//func WithRotationStep(s float64) Option {
//	return withRotationStepOpt{s}
//}
//
//type withRotationStepOpt struct {
//	step float64
//}
//
//func (o withRotationStepOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.Step = o.step
//	}
//}
//
//// WithProp dictates the rate at which rings behind the first one have physics
//// applied to them. High value makes all rings move simultaneously, low value
//// gives them significant delay.
//// WithProp can be used with:
////   - NewChromaLighting()
//func WithProp(p float64) Option {
//	return withPropOpt{p}
//}
//
//type withPropOpt struct {
//	prop float64
//}
//
//func (o withPropOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.Prop = o.prop
//	}
//}
//
//// WithRotationSpeed dictates the s multiplier of the rings
//// WithRotationSpeed can be used with:
////   - NewChromaLighting()
//func WithRotationSpeed(s float64) Option {
//	return withRotationSpeedOpt{s}
//}
//
//type withRotationSpeedOpt struct {
//	s float64
//}
//
//func (o withRotationSpeedOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.Speed = o.s
//	}
//}
//
//// WithRotationDirection dictates the direction to spin the rings
//// WithRotationDirection can be used with:
////   - NewChromaLighting()
//func WithRotationDirection(d chroma.SpinDirection) Option {
//	return withRotationDirectionOpt{d}
//}
//
//type withRotationDirectionOpt struct {
//	d chroma.SpinDirection
//}
//
//func (o withRotationDirectionOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.Direction = o.d
//	}
//}
//
//// WithCounterSpin causes the smaller ring to spin in the opposite direction
//// WithCounterSpin can be used with:
////   - NewChromaLighting()
//func WithCounterSpin(c bool) Option {
//	return withCounterSpinOpt{c}
//}
//
//type withCounterSpinOpt struct {
//	counterSpin bool
//}
//
//func (o withCounterSpinOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.CounterSpin = o.counterSpin
//	}
//}
//
//const IlysaRotationResetTag = "_ilysa_rotation_reset"
//
//func WithRotationReset() Option {
//	return withRotationReset{}
//}
//
//type withRotationReset struct{}
//
//func (o withRotationReset) apply(e Event) {
//	switch r := e.(type) {
//	case *PreciseRotation:
//		r.SetTag(IlysaRotationResetTag)
//	}
//}
//
//// WithLightID causes the event to only affect the light IDs ids.
//// WithLightID can be used with:
////   - NewChromaLighting()
//func WithLightID(id lightid.ID) Option {
//	return withLightIDOpt{id}
//}
//
//type withLightIDOpt struct {
//	l lightid.ID
//}
//
//func (o withLightIDOpt) apply(e Event) {
//	switch l := e.(type) {
//	case *RGBLighting:
//		l.SetLightID(o.l)
//	}
//}
//
//// WithColor sets the color of the event.
//// WithColor can be used with:
////   - NewChromaLighting()
//func WithColor(c color.Color) withColorOpt {
//	return withColorOpt{c}
//}
//
//type withColorOpt struct {
//	c color.Color
//}
//
//func (o withColorOpt) apply(e Event) {
//	switch r := e.(type) {
//	case *RGBLighting:
//		r.SetColor(o.c)
//	}
//}
//
//// WithAlpha sets the alpha of the event's color. Use WithColor to set the
//// event's color before using WithAlpha.
//// WithAlpha can be used with:
////   - NewChromaLighting()
//func WithAlpha(a float64) withAlphaOpt {
//	return withAlphaOpt{a}
//}
//
//type withAlphaOpt struct {
//	a float64
//}
//
//func (o withAlphaOpt) apply(e Event) {
//	switch te := e.(type) {
//	case *RGBLighting:
//		te.SetAlpha(o.a)
//	}
//}
//
//func WithFloatValue(v float64) withFloatValueOpt {
//	return withFloatValueOpt{v}
//}
//
//type withFloatValueOpt struct {
//	v float64
//}
//
//func (o withFloatValueOpt) apply(e Event) {
//	switch te := e.(type) {
//	case *Lighting:
//		te.SetFloatValue(o.v)
//	case *RGBLighting:
//		te.SetFloatValue(o.v)
//	}
//}
//
//type DirectionalLaser int
//
//const (
//	LeftLaser  DirectionalLaser = 0
//	RightLaser DirectionalLaser = 1
//)
//
//// WithLockPosition will not reset laser positions when true
//// WithLockPosition can be used with:
////   - NewChromaLaserSpeed()
//func WithLockPosition(b bool) Option {
//	return withLockPositionOpt{b}
//}
//
//type withLockPositionOpt struct {
//	b bool
//}
//
//func (o withLockPositionOpt) apply(e Event) {
//	plse, ok := e.(*ChromaLaserSpeed)
//	if !ok {
//		return
//	}
//	plse.LockPosition = o.b
//}
//
//func (o withLockPositionOpt) applyPreciseLaser(e *ChromaLaserSpeed) {
//	e.LockPosition = o.b
//}
//
//// WithLaserSpeed sets the laser speed.
//// WithLaserSpeed can be used with:
////   - NewLaserSpeed()
////   - NewChromaLaserSpeed()
//func WithLaserSpeed(v int) Option {
//	return withLaserSpeedOpt{v}
//}
//
//type withLaserSpeedOpt struct {
//	v int
//}
//
//func (o withLaserSpeedOpt) apply(e Event) {
//	switch lse := e.(type) {
//	case *LaserSpeed:
//		lse.SetValue(beatsaber.EventValue(o.v))
//	case *ChromaLaserSpeed:
//		lse.SetValue(beatsaber.EventValue(o.v))
//	}
//}
//
//// WithPreciseLaserSpeed is identical to just setting value, but allows for decimals.
//// Will overwrite value (Because the game will randomize laser position on
//// anything other than value 0, a small trick you can do is set value to 1 and
//// _preciseSpeed to 0, creating 0 s lasers with a randomized position).
//// WithPreciseLaserSpeed can be used with:
////   - NewChromaLaserSpeed()
//func WithPreciseLaserSpeed(s float64) Option {
//	return withPreciseLaserSpeedOpt{s}
//}
//
//type withPreciseLaserSpeedOpt struct {
//	s float64
//}
//
//func (o withPreciseLaserSpeedOpt) apply(e Event) {
//	lse, ok := e.(*ChromaLaserSpeed)
//	if !ok {
//		return
//	}
//	lse.Speed = o.s
//}
//
//func (o withPreciseLaserSpeedOpt) applyPreciseLaser(e *ChromaLaserSpeed) {
//	e.Speed = o.s
//}
//
//// WithLaserDirection set the spin direction of the laser.
//// WithLaserDirection can be used with:
////   - NewChromaLaserSpeed()
//func WithLaserDirection(d chroma.SpinDirection) Option {
//	return withDirectionOpt{d}
//}
//
//type withDirectionOpt struct {
//	direction chroma.SpinDirection
//}
//
//func (o withDirectionOpt) apply(e Event) {
//	ple, ok := e.(*ChromaLaserSpeed)
//	if !ok {
//		return
//	}
//	ple.Direction = o.direction
//}
//
//func (o withDirectionOpt) applyPreciseLaser(e *ChromaLaserSpeed) {
//	e.Direction = o.direction
//}
//
//// WithDirectionalLaser sets the laser that the laser event will apply to.
//// WIthDirectionalLaser can be used with:
////   - NewLaserSpeed()
////   - NewChromaLaserSpeed()
//func WithDirectionalLaser(dl DirectionalLaser) Option {
//	switch dl {
//	case LeftLaser:
//		return withDirectionalLaserOpt{beatsaber.EventTypeLeftRotatingLasersRotationSpeed}
//	case RightLaser:
//		return withDirectionalLaserOpt{beatsaber.EventTypeRightRotatingLasersRotationSpeed}
//	default:
//		panic(fmt.Sprintf("WithDirectionalLaser: unsupported direction %v", dl))
//	}
//}
//
//type withDirectionalLaserOpt struct {
//	t beatsaber.EventType
//}
//
//func (o withDirectionalLaserOpt) apply(e Event) {
//	switch l := e.(type) {
//	case *LaserSpeed:
//		l.SetType(o.t)
//	case *ChromaLaserSpeed:
//		l.SetType(o.t)
//	}
//}
