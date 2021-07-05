package context_test

import (
	"testing"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/timer"
)

func TestContext(t *testing.T) {
	offset1xCtx := context.WithOffset(context.Base(), 3.0)
	if offset1xCtx.Offset() != 3 {
		t.Fatalf("got %f; want %f", offset1xCtx.Offset(), 3.0)
	}

	offset2xCtx := context.WithOffset(offset1xCtx, 3.0)
	if offset2xCtx.Offset() != 6.0 {
		t.Fatalf("got %f; want %f", offset2xCtx.Offset(), 6.0)
	}
}

func TestSeqCtx(t *testing.T) {
	ctx := context.WithOffset(context.Base(), 2)
	context.WithSequence(ctx, timer.SequencerFromSlice([]float64{0, 2, 4, 6, 8}), func(ctx context.Context) {
		context.WithRange(ctx, timer.NewRanger(0, 1, 3, ease.Linear), func(ctx context.Context) {
		})
	})
}
