package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

// DeleteBetween deletes all events between startB and endB (inclusive). If
// startB is negative, DeleteBetween deletes all events before endB. If endB
// is negative, DeleteBetween deletes all events after startB.
func DeleteBetween(ctx context.Context, startB, endB float64) {
	ctx.FilterEvents(func(e evt.Event) bool {
		if (e.Beat() >= startB || startB < 0) && (e.Beat() <= endB || endB < 0) {
			return false
		}
		return true
	})
}

// DeleteLightBetween deletes all events between startB and endB (inclusive)
// generated by light l.
func DeleteLightBetween(ctx context.Context, startB, endB float64, l context.Light) {
	ctx.FilterEvents(func(e evt.Event) bool {
		if (e.Beat() >= startB || startB < 0) && (e.Beat() <= endB || endB < 0) &&
			e.HasTag(l.Name()...) {
			return false
		}
		return true
	})
}
