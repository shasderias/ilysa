package ilysa

import (
	"math/rand"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type baseContext struct {
	*Project

	beatOffset float64
	fixedRand  float64

	defaultOpts []evt.Opt
}

func newBaseContext(p *Project) baseContext {
	return baseContext{
		Project:     p,
		beatOffset:  0,
		fixedRand:   rand.Float64(),
		defaultOpts: []evt.Opt{},
	}
}

func (c baseContext) FixedRand() float64 {
	return c.fixedRand
}

func (c baseContext) LightIDMax(typ beatsaber.EventType) int {
	return c.Project.ActiveDifficultyProfile().LightIDMax(typ)
}

func (c baseContext) withBeatOffset(o float64) baseContext {
	return baseContext{
		Project:     c.Project,
		beatOffset:  o,
		fixedRand:   c.fixedRand,
		defaultOpts: c.defaultOpts,
	}
}

func (c baseContext) WithBeatOffset(o float64) BaseContext {
	return baseContext{
		Project:     c.Project,
		beatOffset:  c.beatOffset + o,
		fixedRand:   rand.Float64(),
		defaultOpts: c.defaultOpts,
	}
}

func (c baseContext) withDefaultOpts(opts ...evt.Opt) baseContext {
	return baseContext{
		Project:     c.Project,
		beatOffset:  c.beatOffset,
		fixedRand:   c.fixedRand,
		defaultOpts: append(c.defaultOpts, opts...),
	}
}

func (c baseContext) addEvent(e evt.Event) {
	e.SetBeat(e.Beat() + c.beatOffset)
	c.Project.events = append(c.Project.events, e)
}

//func (c baseContext) rangeTimer(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(RangeContext)) {
//	startBeat += c.beatOffset
//	endBeat += c.beatOffset
//
//	tScaler := scale.ToUnitIntervalClamped(0, float64(steps-1))
//
//	for i := 0; i < steps; i++ {
//		beat := Ierp(startBeat, endBeat, tScaler(float64(i)), easeFunc)
//		callback(c.withTimer(beat, startBeat, endBeat, i).withBeatOffset(0))
//	}
//}

func (c baseContext) Sequence(s timer.Sequencer, callback func(ctx SequenceContext)) {
	newSequenceContext(c.withBeatOffset(0), s.Offset(c.beatOffset)).iterate(callback)
}

func (c baseContext) Range(startBeat, endBeat float64, steps int, fn ease.Func, callback func(ctx RangeContext)) {

}

//func (c baseContext) ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder func(ctx RangeContext, event Event)) {
//	p := c.Project
//	p.sortEvents()
//
//	startBeat += c.beatOffset
//	endBeat += c.beatOffset
//
//	startIdx, endIdx := 0, len(p.events)
//
//	for i := 0; i < len(p.events); i++ {
//		if p.events[i].Beat() >= startBeat {
//			startIdx = i
//			goto startFound
//		}
//	}
//	// past last event
//	return
//startFound:
//
//	for i := len(p.events) - 1; i >= startIdx; i-- {
//		if p.events[i].Beat() <= endBeat {
//			endIdx = i
//			break
//		}
//	}
//
//	events := p.events[startIdx : endIdx+1]
//
//	for i := range events {
//		if !filter(events[i]) {
//			continue
//		}
//		modder(c.withTimer(events[i].Beat(), startBeat, endBeat, i).withBeatOffset(0), events[i])
//	}
//}

//func (c baseContext) DeleteEvents(startBeat float64, filter EventFilter) {
//	p := c.Project
//	p.sortEvents()
//
//	startBeat += c.beatOffset
//
//	startIdx := 0
//
//	for i := 0; i < len(p.events); i++ {
//		if p.events[i].Beat() >= startBeat {
//			startIdx = i
//			goto startFound
//		}
//	}
//	// past last event
//	return
//startFound:
//
//	events := p.events[:startIdx]
//
//	for _, e := range p.events[startIdx:] {
//		if filter(e) {
//			events = append(events, e)
//		}
//	}
//
//	p.events = events
//}

func (c baseContext) NewLighting(opts ...evt.LightingOpt) *evt.Lighting {
	e := evt.NewLighting()
	evt.Apply(&e, c.defaultOpts...)
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	e := evt.NewRGBLighting()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewLaser(opts ...evt.LaserOpt) *evt.Laser {
	e := evt.NewLaser()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser {
	e := evt.NewPreciseLaser()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewRotation(opts ...evt.RotationOpt) *evt.Rotation {
	e := evt.NewRotation()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation {
	e := evt.NewPreciseRotation()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewZoom(opts ...evt.ZoomOpt) *evt.Zoom {
	e := evt.NewZoom()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}

func (c baseContext) NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom {
	e := evt.NewPreciseZoom()
	e.Apply(opts...)
	c.addEvent(&e)
	return &e
}
