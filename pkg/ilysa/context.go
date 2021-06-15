package ilysa

import (
	"math/rand"

	"ilysa/pkg/chroma"
	"ilysa/pkg/chroma/lightid"
	"ilysa/pkg/light"
	"ilysa/pkg/util"
)

type Context struct {
	*Project
	Timing

	RandFloat64 float64
	modifiers   []EventModifier
}

type Timing struct {
	StartBeat float64
	EndBeat   float64
	Duration  float64
	B         float64
	Pos       float64
	Ordinal   int
	Last      bool
}

func newContext(p *Project) Context {
	return Context{
		Project:     p,
		RandFloat64: rand.Float64(),
	}
}

func (c *Context) newTiming(startBeat, endBeat, beat float64, ordinal int) {
	c.Timing = NewTiming(startBeat, endBeat, beat, ordinal)
}

func NewTiming(startBeat, endBeat, beat float64, ordinal int) Timing {
	return Timing{
		StartBeat: startBeat,
		EndBeat:   endBeat,
		B:         beat,
		Duration:  endBeat - startBeat,
		Pos:       util.ScaleToUnitInterval(startBeat, endBeat)(beat),
		Ordinal:   ordinal,
		Last:      endBeat == beat,
	}
}

func (c Context) WithModifier(modifiers ...EventModifier) Context {
	return Context{
		Project:     c.Project,
		Timing:      c.Timing,
		RandFloat64: c.RandFloat64,
		modifiers:   append([]EventModifier{}, modifiers...),
	}
}

func (c Context) addEvent(e Event) {
	c.applyModifiers(e)
	c.Project.events = append(c.Project.events, e)
}

func (c Context) applyModifiers(e Event) {
	if len(c.modifiers) == 0 {
		return
	}

	for _, m := range c.modifiers {
		m(e)
	}
}

type RangeLightIDContext struct {
	Context

	MinLightID     int
	MaxLightID     int
	CurLightID     light.ID
	PreLightID     light.ID
	LightIDPos     float64
	LightIDOrdinal int
	LightIDSet     lightid.Set
	LightIDSetLen  int
}

func (c Context) RangeLightIDs(light Light, picker lightid.Picker, callback func(ctx RangeLightIDContext)) {
	set := picker(light)

	for i, lightID := range set {
		ctx := c.WithModifier(func(e Event) {
			le, ok := e.(*RGBLightingEvent)
			if !ok {
				return
			}
			le.BaseEvent.Type = light.EventType()
			le.LightID = lightID
		})
		callback(RangeLightIDContext{
			Context:        ctx,
			MinLightID:     light.MinLightID(),
			MaxLightID:     light.MaxLightID(),
			CurLightID:     lightID,
			PreLightID:     set.Pick(i - 1),
			LightIDPos:     float64(i) / float64(len(set)),
			LightIDOrdinal: i,
			LightIDSet:     set,
			LightIDSetLen:  len(set),
		})
	}
}

func (c Context) RangeLight(l light.Light, callback func(ctx RangeLightIDContext)) {
	for i := l.LightIDMin(); i <= l.LightIDMax(); i++ {
		ctx := c.WithModifier(func(e Event) {
			le, ok := e.(*RGBLightingEvent)
			if !ok {
				return
			}
			le.BaseEvent.Type = l.EventType(i)
			le.SetLightID(l.LightID(i))
		})
		callback(RangeLightIDContext{
			Context: ctx,
			MinLightID: l.LightIDMin(),
			MaxLightID: l.LightIDMax(),
			CurLightID: l.LightID(i),
			LightIDPos: float64(i)/float64(l.LightIDLen()),
		})
	}
}

type EventModifier func(e Event)
type EventGenerator func(ctx Context)
type EventModder func(ctx Context, event Event)
type EventFilter func(event Event) bool


// time - 0.25 - 0.50
// light - back lasers
// color rainbow, varying with time
// alpha