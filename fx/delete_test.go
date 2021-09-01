package fx_test

import (
	"testing"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
)

func TestFilterEvents(t *testing.T) {
	mockMap, err := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	if err != nil {
		t.Fatal(err)
	}

	ctx := ilysa.New(mockMap)

	assertProjectHasNEvents(t, ctx, 0)

	ctx.BeatRange(0, 1, 5, ease.Linear, func(ctx context.Context) {
		ctx.NewRGBLighting(evt.WithLight(evt.RingLights))
	})

	assertProjectHasNEvents(t, ctx, 5)

	// 0, 0.25, 0.5, 0.75, 1
	fx.DeleteBetween(ctx, 0.5, -1)

	assertProjectHasNEvents(t, ctx, 2)

	ringLights := light.NewBasic(ctx, evt.RingLights)

	ctx.BeatRange(2, 4, 9, ease.Linear, func(ctx context.Context) {
		ctx.Light(ringLights, func(ctx context.LightContext) {
			ctx.NewRGBLighting()
		})
	})

	assertProjectHasNEvents(t, ctx, 11)

	fx.DeleteLightBetween(ctx, -1, 3, ringLights)

	// 2, 2.25, 2.50, 2.75, 3
	//    3.25, 3.50, 3.75, 4
	assertProjectHasNEvents(t, ctx, 6)
}

func assertProjectHasNEvents(t *testing.T, ctx *ilysa.Project, n int) {
	t.Helper()
	if l := len(*ctx.Events()); l != n {
		t.Fatalf("got map with %d events; want %d", l, n)
	}
}
