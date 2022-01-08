package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

// Delete deletes all events that match all the given filterFuncs.
func Delete(ctx context.Context, filterFuncs ...func(e evt.Event) bool) {
	n := 0
	events := ctx.Events()

	for _, evt := range *events {
		for _, fn := range filterFuncs {
			if !fn(evt) {
				goto keep
			}
		}
		// not keeping
		continue
	keep:
		(*events)[n] = evt
		n++
	}

	*events = (*events)[:n]
}

func DelAfterB(b float64) func(evt.Event) bool {
	return func(e evt.Event) bool {
		return e.Beat() > b
	}
}

func DelBetweenB(startB, endB float64) func(evt.Event) bool {
	return func(e evt.Event) bool {
		if (e.Beat() >= startB || startB < 0) && (e.Beat() <= endB || endB < 0) {
			return true
		}
		return false
	}
}

func DelWithType(typ evt.Type) func(evt.Event) bool {
	return func(e evt.Event) bool {
		return e.Type() == typ
	}
}
