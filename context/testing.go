package context

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/timer"
)

type RefTiming struct {
	B              float64
	BWOffset       float64
	T              float64
	SeqT           float64
	SeqNextB       float64
	SeqNextBOffset float64
	LightID        lightid.ID
	LightIDT       float64
}

func refTimingFromContext(ctx Context) RefTiming {
	e := ctx.NewRGBLighting()

	return RefTiming{
		B:              ctx.B(),
		BWOffset:       e.Beat(),
		T:              ctx.T(),
		SeqT:           ctx.SeqT(),
		SeqNextB:       ctx.SeqNextB(),
		SeqNextBOffset: ctx.SeqNextBOffset(),
	}
}

func refTimingFromLightContext(ctx LightRGBLightingContext) RefTiming {
	e := ctx.NewRGBLighting()

	return RefTiming{
		B:              ctx.B(),
		BWOffset:       e.Beat(),
		T:              ctx.T(),
		SeqT:           ctx.SeqT(),
		SeqNextB:       ctx.SeqNextB(),
		SeqNextBOffset: ctx.SeqNextBOffset(),
		LightIDT:       ctx.LightIDT(),
	}
}

type MockProject struct {
	t          *testing.T
	maxLightID int
	events     []evt.Event
	refTimings []RefTiming
}

func NewMockProject(t *testing.T, maxLightID int) *MockProject {
	return &MockProject{
		t:          t,
		maxLightID: maxLightID,
		events:     make([]evt.Event, 0),
	}
}

func (p *MockProject) MaxLightID(t evt.LightType) int {
	return p.maxLightID
}

func (p *MockProject) AddEvents(events ...evt.Event) {
	p.events = append(p.events, events...)
}

func (p *MockProject) MockLight() MockLight {
	return newMockLight(p)
}

func (p *MockProject) addRefTiming(t RefTiming) {
	p.refTimings = append(p.refTimings, t)
}

func (p *MockProject) AddRefTimingFromCtx(ctx Context) {
	ref := refTimingFromContext(ctx)

	e := ctx.NewRGBLighting()
	ref.BWOffset = e.Beat()

	p.addRefTiming(ref)
}

type MockLight struct {
	*MockProject
}

func newMockLight(p *MockProject) MockLight {
	return MockLight{p}
}

func (l MockLight) NewRGBLighting(ctx LightRGBLightingContext) evt.RGBLightingEvents {
	t := refTimingFromLightContext(ctx)

	e := ctx.NewRGBLighting(evt.WithLightID(lightid.New(ctx.LightIDCur())))
	t.LightID = lightid.ID(e.LightID)

	l.addRefTiming(t)

	return evt.RGBLightingEvents{e}
}

func (l MockLight) LightIDLen() int {
	return l.MockProject.maxLightID
}

func (l MockLight) Name() []string {
	return []string{fmt.Sprintf("MockLight")}
}

func (p *MockProject) Cmp(t []RefTiming) {
	if diff := cmp.Diff(t, p.refTimings, cmpopts.EquateApprox(0.000001, 0)); diff != "" {
		p.t.Fatal(diff)
	}
}

func (p *MockProject) Events() *[]evt.Event {
	return &p.events
}

func (p *MockProject) RefTimings() []RefTiming {
	return p.refTimings
}

func (p *MockProject) Sequence(s timer.Sequencer, callback func(ctx Context)) {
	Base(p).Sequence(s, callback)
}

func (p *MockProject) Range(r timer.Ranger, callback func(ctx Context)) {
	Base(p).Range(r, callback)
}

func (p *MockProject) FilterEvents(f func(e evt.Event) bool) *[]evt.Event {
	filteredEvents := make([]evt.Event, 0)
	for _, e := range p.events {
		if f(e) {
			filteredEvents = append(filteredEvents, e)
		}
	}
	p.events = filteredEvents
	return &filteredEvents
}

func (p *MockProject) MapEvents(f func(e evt.Event) evt.Event) {
	for i := range p.events {
		p.events[i] = f(p.events[i])
	}
}
