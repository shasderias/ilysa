package beatsaber

import (
	"encoding/json"
	"os"
	"reflect"
	"sort"
	"strconv"

	"github.com/shasderias/ilysa/internal/swallowjson"
)

func OpenDifficultyV250(info *Info, path string) (Difficulty, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var diff DifficultyV250

	err = json.Unmarshal(f, &diff)
	if err != nil {
		return nil, err
	}

	diff.info = info
	diff.filepath = path
	diff.calculateBPMRegions()

	return &diff, nil
}

type DifficultyV250 struct {
	info       *Info
	filepath   string
	bpmRegions []bpmRegion

	Version string      `json:"_version"`
	Notes   []Note      `json:"_notes"`
	Events  []EventV250 `json:"_events"`

	CustomData DifficultyCustomData `json:"_customData"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (d *DifficultyV250) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(d, "Extra", raw)
}

func (d DifficultyV250) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(d, "Extra")
}

func (d *DifficultyV250) Save() error {
	for i := range d.Events {
		d.Events[i].unscaleBeatToTime(d.UnscaleTime)
	}

	sort.Slice(d.Events, func(i, j int) bool {
		return d.Events[i].Time < d.Events[j].Time
	})

	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return os.WriteFile(d.filepath, bytes, 0644)
}

func (d *DifficultyV250) SetEvents(events interface{}) {
	type EventV250er interface{ EventV250() EventV250 }

	typ := reflect.TypeOf(events)
	if typ.Kind() != reflect.Slice {
		panic("events must be a slice")
	}

	val := reflect.ValueOf(events)
	l := val.Len()
	v250Events := make([]EventV250, l)
	for i := 0; i < l; i++ {
		elem := val.Index(i).Interface()
		e, ok := elem.(EventV250er)
		if !ok {
			panic("cannot cast index " + strconv.Itoa(i) + " to EventV250er")
		}
		v250Events[i] = e.EventV250()
	}
	d.Events = v250Events
}

func (d *DifficultyV250) calculateBPMRegions() {
	d.bpmRegions = calculateBPMRegions(d.info.BPM, d.CustomData.BPMChanges)
}

func (d *DifficultyV250) UnscaleTime(beat float64) Time {
	return unscaleTime(d.info.BPM, d.bpmRegions, beat)
}

func (d *DifficultyV250) DifficultyVersion() DifficultyVersion {
	return DifficultyVersion_2_5_0
}

type EventV250 struct {
	beat float64

	Time       Time    `json:"_time"`
	Type       int     `json:"_type"`
	Value      int     `json:"_value"`
	FloatValue float64 `json:"_floatValue"`

	CustomData json.RawMessage `json:"_customData,omitempty"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func NewEventV250(beat float64, typ int, value int, floatValue float64, customData json.RawMessage) EventV250 {
	return EventV250{
		beat:       beat,
		Type:       typ,
		Value:      value,
		FloatValue: floatValue,
		CustomData: customData,
	}
}

func (e *EventV250) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(e, "Extra", raw)
}

func (e EventV250) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(e, "Extra")
}

func (e *EventV250) unscaleBeatToTime(unscaleFn func(beat float64) Time) {
	e.Time = unscaleFn(e.beat)
}
