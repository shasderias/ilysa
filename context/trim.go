package context

import "github.com/shasderias/ilysa/evt"

func trimEvents(events *[]evt.Event, trimPoint float64) {
	trimmedEvents := make([]evt.Event, 0)

	for _, e := range *events {
		if e.Beat() < trimPoint {
			trimmedEvents = append(trimmedEvents, e)
		}
	}

	*events = trimmedEvents
}
