package beatsaber

import (
	"encoding/json"
	"os"
	"reflect"
	"sort"
	"strconv"

	"github.com/shasderias/ilysa/internal/swallowjson"
)

func OpenDifficultyV220(info *Info, path string) (Difficulty, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var diff DifficultyV220

	err = json.Unmarshal(f, &diff)
	if err != nil {
		return nil, err
	}

	diff.info = info
	diff.filepath = path
	diff.calculateBPMRegions()

	return &diff, nil
}

type DifficultyV220 struct {
	info       *Info
	filepath   string
	bpmRegions []bpmRegion

	Version string      `json:"_version"`
	Notes   []Note      `json:"_notes"`
	Events  []EventV220 `json:"_events"`

	CustomData DifficultyCustomData `json:"_customData"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (d *DifficultyV220) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(d, "Extra", raw)
}

func (d DifficultyV220) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(d, "Extra")
}

func (d *DifficultyV220) Save() error {
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

func (d *DifficultyV220) SetEvents(events interface{}) {
	type EventV220er interface{ EventV220() EventV220 }

	typ := reflect.TypeOf(events)
	if typ.Kind() != reflect.Slice {
		panic("events must be a slice")
	}

	val := reflect.ValueOf(events)
	l := val.Len()
	v220Events := make([]EventV220, l)
	for i := 0; i < l; i++ {
		elem := val.Index(i).Interface()
		e, ok := elem.(EventV220er)
		if !ok {
			panic("cannot cast index " + strconv.Itoa(i) + " to EventV220er")
		}
		v220Events[i] = e.EventV220()
	}
	d.Events = v220Events
}

func (d *DifficultyV220) calculateBPMRegions() {
	d.bpmRegions = calculateBPMRegions(d.info.BPM, d.CustomData.BPMChanges)
}

func (d *DifficultyV220) UnscaleTime(beat float64) Time {
	return unscaleTime(d.info.BPM, d.bpmRegions, beat)
}

func (d *DifficultyV220) DifficultyVersion() DifficultyVersion {
	return DifficultyVersion_2_2_0
}

type EventV220 struct {
	beat float64

	Time  Time `json:"_time"`
	Type  int  `json:"_type"`
	Value int  `json:"_value"`

	CustomData json.RawMessage `json:"_customData,omitempty"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func NewEventV220(beat float64, typ int, value int, customData json.RawMessage) EventV220 {
	return EventV220{
		beat:       beat,
		Type:       typ,
		Value:      value,
		CustomData: customData,
	}
}

func (e *EventV220) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(e, "Extra", raw)
}

func (e EventV220) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(e, "Extra")
}

func (e *EventV220) unscaleBeatToTime(unscaleFn func(float64) Time) {
	e.Time = unscaleFn(e.beat)
}
