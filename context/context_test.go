package context_test

import (
	"testing"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/timer"
)

func TestSeqCtx(t *testing.T) {
	t.Run("Sequence", func(t *testing.T) {
		proj := context.NewMockProject(t, 3)
		proj.Sequence(timer.SequencerFromSlice([]float64{0, 2, 4}), func(ctx context.Context) {
			proj.AddRefTimingFromCtx(ctx)
		})

		want := []context.RefTiming{
			{0, 0, 0, 0, 2, 2, nil, 0},
			{2, 2, 1, 1, 0, 2, nil, 0},
		}

		proj.Cmp(want)
	})
	t.Run("Range", func(t *testing.T) {
		proj := context.NewMockProject(t, 3)
		proj.Range(timer.NewRanger(1, 2, 5, ease.Linear), func(ctx context.Context) {
			proj.AddRefTimingFromCtx(ctx)
		})

		want := []context.RefTiming{
			{1, 1, 0, 0, 1.2, 0.2, nil, 0},
			{1.25, 1.25, 0.25, 0.25, 1.4, 0.2, nil, 0},
			{1.5, 1.5, 0.5, 0.5, 1.6, 0.2, nil, 0},
			{1.75, 1.75, 0.75, 0.75, 1.8, 0.2, nil, 0},
			{2, 2, 1, 1, 2, 0.2, nil, 0},
		}

		proj.Cmp(want)
	})
	t.Run("Offset/Sequence/Range/Light", func(t *testing.T) {
		proj := context.NewMockProject(t, 3)
		light := proj.MockLight()
		ctx := context.WithOffset(context.Base(proj), 2)
		ctx.Sequence(timer.SequencerFromSlice([]float64{0, 2, 4}), func(ctx context.Context) {
			ctx.Range(timer.NewRanger(0, 1, 3, ease.Linear), func(ctx context.Context) {
				ctx.Light(light, func(ctx context.LightContext) {
					ctx.NewRGBLighting()
				})
			})
		})

		want := []context.RefTiming{
			{0, 2, 0, 0, 2, 2, lightid.ID{1}, 0},
			{0, 2, 0, 0, 2, 2, lightid.ID{2}, 0.5},
			{0, 2, 0, 0, 2, 2, lightid.ID{3}, 1},
			{0.5, 2.5, 0.5, 0, 2, 2, lightid.ID{1}, 0},
			{0.5, 2.5, 0.5, 0, 2, 2, lightid.ID{2}, 0.5},
			{0.5, 2.5, 0.5, 0, 2, 2, lightid.ID{3}, 1},
			{1, 3, 1, 0, 2, 2, lightid.ID{1}, 0},
			{1, 3, 1, 0, 2, 2, lightid.ID{2}, 0.5},
			{1, 3, 1, 0, 2, 2, lightid.ID{3}, 1},
			{0, 4, 0, 1, 0, 2, lightid.ID{1}, 0},
			{0, 4, 0, 1, 0, 2, lightid.ID{2}, 0.5},
			{0, 4, 0, 1, 0, 2, lightid.ID{3}, 1},
			{0.5, 4.5, 0.5, 1, 0, 2, lightid.ID{1}, 0},
			{0.5, 4.5, 0.5, 1, 0, 2, lightid.ID{2}, 0.5},
			{0.5, 4.5, 0.5, 1, 0, 2, lightid.ID{3}, 1},
			{1, 5, 1, 1, 0, 2, lightid.ID{1}, 0},
			{1, 5, 1, 1, 0, 2, lightid.ID{2}, 0.5},
			{1, 5, 1, 1, 0, 2, lightid.ID{3}, 1},
		}
		proj.Cmp(want)
	})
}
