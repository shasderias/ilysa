package evt

import "github.com/shasderias/ilysa/chroma"

// WithNameFilter causes event to only affect rings with the name filter (e.g.
// SmallTrackLaneRings, BigTrackLaneRings).
func WithNameFilter(filter string) withNameFilterOpt {
	return withNameFilterOpt{filter}
}

type withNameFilterOpt struct {
	nameFilter string
}

func (o withNameFilterOpt) applyPreciseRotation(e *PreciseRotation) {
	e.NameFilter = o.nameFilter
}

// WithReset resets the rings when set to true (overwrites other values below)
func WithReset(r bool) withResetOpt {
	return withResetOpt{r}
}

type withResetOpt struct {
	reset bool
}

func (o withResetOpt) applyPreciseRotation(e *PreciseRotation) {
	e.Reset = o.reset
}

// WithRotation dictates how far the first ring will spin
func WithRotation(r float64) withRotationOpt {
	return withRotationOpt{r}
}

type withRotationOpt struct {
	rotation float64
}

func (o withRotationOpt) applyPreciseRotation(e *PreciseRotation) {
	e.Rotation = o.rotation
}

// WithRotationStep dictates how much rotation is added between each ring
func WithRotationStep(s float64) withRotationStep {
	return withRotationStep{s}
}

type withRotationStep struct {
	step float64
}

func (o withRotationStep) applyPreciseRotation(e *PreciseRotation) {
	e.Step = o.step
}

// WithProp dictates the rate at which rings behind the first one have physics
// applied to them. High value makes all rings move simultaneously, low value
// gives them significant delay.
func WithProp(p float64) withPropOpt {
	return withPropOpt{p}
}

type withPropOpt struct {
	prop float64
}

func (o withPropOpt) applyPreciseRotation(e *PreciseRotation) {
	e.Prop = o.prop
}

// WithRotationSpeed dictates the s multiplier of the rings
func WithRotationSpeed(s float64) withRotationSpeedOpt {
	return withRotationSpeedOpt{s}
}

type withRotationSpeedOpt struct {
	s float64
}

func (o withRotationSpeedOpt) applyPreciseRotation(e *PreciseRotation) {
	e.Speed = o.s
}

// WithRotationDirection dictates the direction to spin the rings
func WithRotationDirection(d chroma.SpinDirection) withRotationDirectionOpt {
	return withRotationDirectionOpt{d}
}

type withRotationDirectionOpt struct {
	d chroma.SpinDirection
}

func (o withRotationDirectionOpt) applyPreciseRotation(e *PreciseRotation) {
	e.Direction = o.d
}

// WithCounterSpin causes the smaller ring to spin in the opposite direction
func WithCounterSpin(c bool) withCounterSpinOpt {
	return withCounterSpinOpt{c}
}

type withCounterSpinOpt struct {
	counterSpin bool
}

func (o withCounterSpinOpt) applyPreciseRotation(e *PreciseRotation) {
	e.CounterSpin = o.counterSpin
}
