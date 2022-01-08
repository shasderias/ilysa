package context

import (
	"testing"

	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
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
	return RefTiming{
		B:              ctx.B(),
		BWOffset:       ctx.B() + ctx.BOffset(),
		T:              ctx.T(),
		SeqT:           ctx.SeqT(),
		SeqNextB:       ctx.SeqNextB(),
		SeqNextBOffset: ctx.SeqNextBOffset(),
	}
}

func refTimingFromLightContext(ctx LightContext) RefTiming {
	return RefTiming{
		B:              ctx.B(),
		BWOffset:       ctx.B() + ctx.BOffset(),
		T:              ctx.T(),
		SeqT:           ctx.SeqT(),
		SeqNextB:       ctx.SeqNextB(),
		SeqNextBOffset: ctx.SeqNextBOffset(),
		LightIDT:       ctx.LightT(),
	}
}

type MockProject struct {
	t          *testing.T
	events     *evt.Events
	refTimings []RefTiming
}

func NewMockProject(t *testing.T) *MockProject {
	events := evt.NewEvents()

	return &MockProject{
		t:      t,
		events: &events,
	}
}
func (p *MockProject) addRefTiming(t RefTiming) {
	p.refTimings = append(p.refTimings, t)
}

func (p *MockProject) AddRefTimingFromCtx(ctx Context) {
	ref := refTimingFromContext(ctx)
	p.addRefTiming(ref)
}

func (p *MockProject) RefTimings() []RefTiming {
	return p.refTimings
}

func (p *MockProject) Events() *evt.Events {
	return p.events
}

func (p *MockProject) MockLight(maxLightID int) Light {
	return &mockLight{
		MockProject: p,
		maxLightID:  maxLightID,
	}
}

type mockLight struct {
	*MockProject
	maxLightID int
}

func (m mockLight) GenerateEvents(ctx LightContext) evt.Events {
	refTiming := refTimingFromLightContext(ctx)
	refTiming.LightID = lightid.New(ctx.LightCur())
	m.addRefTiming(refTiming)
	return evt.NewEvents()
}

func (m mockLight) LightLen() int {
	return m.maxLightID
}

func (m mockLight) Name() []string {
	return []string{"MockLight"}
}
