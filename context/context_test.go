package context_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/timer"
)

func TestContextTimer(t *testing.T) {
	t.Run("Sequence", func(t *testing.T) {
		proj := context.NewMockProject(t)
		ctx := context.Base(proj)
		ctx.WSeq(timer.Seq(0, 2, 4), func(ctx context.Context) {
			proj.AddRefTimingFromCtx(ctx)
		})

		want := []context.RefTiming{
			{0, 0, 0, 0, 2, 2, nil, 0},
			{2, 2, 1, 1, 0, 2, nil, 0},
		}

		if diff := cmp.Diff(proj.RefTimings(), want); diff != "" {
			t.Fatal(diff)
		}
	})

	t.Run("Range", func(t *testing.T) {
		proj := context.NewMockProject(t)
		ctx := context.Base(proj)
		ctx.WRng(timer.Rng(1, 2, 5, ease.Linear), func(ctx context.Context) {
			proj.AddRefTimingFromCtx(ctx)
		})

		want := []context.RefTiming{
			{1.00, 1.00, 0.00, 0.00, 1.25, 0.25, nil, 0},
			{1.25, 1.25, 0.25, 0.25, 1.50, 0.25, nil, 0},
			{1.50, 1.50, 0.50, 0.50, 1.75, 0.25, nil, 0},
			{1.75, 1.75, 0.75, 0.75, 2.00, 0.25, nil, 0},
			{2.00, 2.00, 1.00, 1.00, 2.00, 0.25, nil, 0},
		}

		if diff := cmp.Diff(proj.RefTimings(), want); diff != "" {
			t.Fatal(diff)
		}
	})

	t.Run("RangeInterval1", func(t *testing.T) {
		proj := context.NewMockProject(t)
		ctx := context.Base(proj)

		ctx.WRng(timer.RngInterval(1, 2, 4, ease.Linear), func(ctx context.Context) {
			proj.AddRefTimingFromCtx(ctx)
		})

		want := []context.RefTiming{
			{1.00, 1.00, 0.00, 0.00, 1.25, 0.25, nil, 0},
			{1.25, 1.25, 0.25, 0.25, 1.50, 0.25, nil, 0},
			{1.50, 1.50, 0.50, 0.50, 1.75, 0.25, nil, 0},
			{1.75, 1.75, 0.75, 0.75, 2.00, 0.25, nil, 0},
			{2.00, 2.00, 1.00, 1.00, 2.00, 0.25, nil, 0},
		}

		if diff := cmp.Diff(proj.RefTimings(), want); diff != "" {
			t.Fatal(diff)
		}
	})

	t.Run("RangeInterval2", func(t *testing.T) {
		proj := context.NewMockProject(t)
		ctx := context.Base(proj)

		ctx.WRng(timer.RngInterval(1, 2, 8, ease.Linear), func(ctx context.Context) {
			proj.AddRefTimingFromCtx(ctx)
		})

		want := []context.RefTiming{
			{1.000, 1.000, 0.000, 0.000, 1.125, 0.125, nil, 0},
			{1.125, 1.125, 0.125, 0.125, 1.250, 0.125, nil, 0},
			{1.250, 1.250, 0.250, 0.250, 1.375, 0.125, nil, 0},
			{1.375, 1.375, 0.375, 0.375, 1.500, 0.125, nil, 0},
			{1.500, 1.500, 0.500, 0.500, 1.625, 0.125, nil, 0},
			{1.625, 1.625, 0.625, 0.625, 1.750, 0.125, nil, 0},
			{1.750, 1.750, 0.750, 0.750, 1.875, 0.125, nil, 0},
			{1.875, 1.875, 0.875, 0.875, 2.000, 0.125, nil, 0},
			{2.000, 2.000, 1.000, 1.000, 2.000, 0.125, nil, 0},
		}

		if diff := cmp.Diff(proj.RefTimings(), want); diff != "" {
			t.Fatal(diff)
		}
	})

	t.Run("BOffset/Sequence/Range/Light", func(t *testing.T) {
		proj := context.NewMockProject(t)
		light := proj.MockLight(3)

		ctx := context.Base(proj)
		ctx = ctx.WBOffset(2)
		ctx.WSeq(timer.Seq(0, 2, 4), func(ctx context.Context) {
			ctx.WRng(timer.Rng(0, 1, 3, ease.Linear), func(ctx context.Context) {
				ctx.WLight(light, func(ctx context.LightContext, e evt.Events) {

				})
			})
		})

		want := []context.RefTiming{
			{0.0, 2.0, 0.0, 0, 2, 2, lightid.ID{1}, 0.0},
			{0.0, 2.0, 0.0, 0, 2, 2, lightid.ID{2}, 0.5},
			{0.0, 2.0, 0.0, 0, 2, 2, lightid.ID{3}, 1.0},
			{0.5, 2.5, 0.5, 0, 2, 2, lightid.ID{1}, 0.0},
			{0.5, 2.5, 0.5, 0, 2, 2, lightid.ID{2}, 0.5},
			{0.5, 2.5, 0.5, 0, 2, 2, lightid.ID{3}, 1.0},
			{1.0, 3.0, 1.0, 0, 2, 2, lightid.ID{1}, 0.0},
			{1.0, 3.0, 1.0, 0, 2, 2, lightid.ID{2}, 0.5},
			{1.0, 3.0, 1.0, 0, 2, 2, lightid.ID{3}, 1.0},
			{0.0, 4.0, 0.0, 1, 0, 2, lightid.ID{1}, 0.0},
			{0.0, 4.0, 0.0, 1, 0, 2, lightid.ID{2}, 0.5},
			{0.0, 4.0, 0.0, 1, 0, 2, lightid.ID{3}, 1.0},
			{0.5, 4.5, 0.5, 1, 0, 2, lightid.ID{1}, 0.0},
			{0.5, 4.5, 0.5, 1, 0, 2, lightid.ID{2}, 0.5},
			{0.5, 4.5, 0.5, 1, 0, 2, lightid.ID{3}, 1.0},
			{1.0, 5.0, 1.0, 1, 0, 2, lightid.ID{1}, 0.0},
			{1.0, 5.0, 1.0, 1, 0, 2, lightid.ID{2}, 0.5},
			{1.0, 5.0, 1.0, 1, 0, 2, lightid.ID{3}, 1.0},
		}

		if diff := cmp.Diff(proj.RefTimings(), want); diff != "" {
			t.Fatal(diff)
		}
	})
}
