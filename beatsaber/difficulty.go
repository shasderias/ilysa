package beatsaber

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/shasderias/ilysa/internal/swallowjson"
	"github.com/shasderias/ilysa/scale"
)

type Difficulty struct {
	info       *Info
	filepath   string
	bpmRegions []bpmRegion

	Version string  `json:"_version"`
	Notes   []Note  `json:"_notes"`
	Events  []Event `json:"_events"`

	CustomData DifficultyCustomData `json:"_customData"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (d *Difficulty) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(d, "Extra", raw)
}

func (d Difficulty) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(d, "Extra")
}

func (d *Difficulty) Save() error {
	sort.Slice(d.Events, func(i, j int) bool {
		return d.Events[i].Time < d.Events[j].Time
	})

	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return os.WriteFile(d.filepath, bytes, 0644)

	//f, err := os.OpenFile(d.filepath, os.O_RDWR|os.O_TRUNC, 0755)
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//
	//enc := json.NewEncoder(f)
	//
	//fmt.Println(d)
	//err = enc.Encode(d)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}

func (d *Difficulty) GetEvents() *[]Event {
	return &d.Events
}

type Note struct {
	Time         Time         `json:"_time"`
	LineIndex    int          `json:"_lineIndex"`
	LineLayer    int          `json:"_lineLayer"`
	Type         NoteType     `json:"_type"`
	CutDirection CutDirection `json:"_cutDirection"`

	Extra map[string]*json.RawMessage `json:"-"`
}

type Time float64
type NoteType int
type CutDirection int

type DifficultyCustomData struct {
	BPMChanges []BPMChange `json:"_BPMChanges"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (dcd *DifficultyCustomData) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(dcd, "Extra", raw)
}

func (dcd DifficultyCustomData) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(dcd, "Extra")
}

type Event struct {
	Time       Time    `json:"_time"`
	Type       int     `json:"_type"`
	Value      int     `json:"_value"`
	FloatValue float64 `json:"_floatValue"`

	CustomData json.RawMessage `json:"_customData,omitempty"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (e *Event) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(e, "Extra", raw)
}

func (e Event) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(e, "Extra")
}

type BPMChange struct {
	Time            Time    `json:"_time"`
	BPM             float64 `json:"_BPM"`
	BeatsPerBar     int     `json:"_beatsPerBar"`
	MetronomeOffset int     `json:"_metronomeOffset"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (bpmc *BPMChange) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(bpmc, "Extra", raw)
}

func (bpmc BPMChange) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(bpmc, "Extra")
}

type bpmRegion struct {
	start     float64
	bpm       float64
	startBeat float64
}

func (d *Difficulty) calculateBPMRegions() {
	startBPM := d.info.BPM

	bpmChanges := d.CustomData.BPMChanges

	bpmRegions := []bpmRegion{{
		0,
		startBPM,
		0,
	}}

	for i, change := range bpmChanges {
		t := float64(change.Time)
		prevRegion := bpmRegions[i]

		// adapted from https://github.com/Caeden117/ChroMapper/blob/41a64f50212de47b252a7d33881cfab8f78aea32/Assets/__Scripts/MapEditor/Grid/Collections/BPMChangesContainer.cs
		passedBeats := (t - prevRegion.start - 0.01) / startBPM * prevRegion.bpm
		scaledBeat := prevRegion.startBeat + math.Ceil(passedBeats)

		bpmRegions = append(bpmRegions, bpmRegion{
			t,
			change.BPM,
			scaledBeat,
		})
	}

	d.bpmRegions = bpmRegions
}

func (d *Difficulty) UnscaleTime(beat float64) Time {
	startBPM := d.info.BPM
	bpmRegions := d.bpmRegions

	for i := len(bpmRegions) - 1; i >= 0; i-- {
		region := bpmRegions[i]

		if beat >= region.startBeat {
			if i != len(bpmRegions)-1 {
				nextRegion := bpmRegions[i+1]

				if math.Floor(beat)+1.0 == nextRegion.startBeat {
					integer, frac := math.Modf(beat)
					unscaledCritStart := region.start + ((integer - region.startBeat) / region.bpm * startBPM)

					scaler := scale.FromUnitClamp(unscaledCritStart, nextRegion.start)
					return Time(scaler(frac))
				}
			}

			diff := beat - region.startBeat

			scaledBeat := Time(region.start + (diff / region.bpm * startBPM))

			return scaledBeat
		}
	}

	panic(fmt.Sprintf("UnscaleTime(): unreachable code, beat: %f", beat))
}
