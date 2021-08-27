package beatsaber

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"

	"github.com/shasderias/ilysa/internal/swallowjson"
	"github.com/shasderias/ilysa/scale"
)

type Map struct {
	workingDir string

	Info *Info

	activeDifficulty           *Difficulty
	activeDifficultyPath       string
	activeCharacteristic       Characteristic
	activeBeatmapDifficulty    BeatmapDifficulty
	activeDifficultyBPMRegions []bpmRegion
	activeEnvironmentProfile   *EnvProfile
}

func Open(dir string) (*Map, error) {
	infoPath := filepath.Join(dir, "Info.dat")

	info, err := openInfo(infoPath)
	if err != nil {
		return nil, err
	}

	return &Map{
		workingDir: dir,
		Info:       info,
	}, nil
}

func openInfo(filename string) (*Info, error) {
	inf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var info Info

	err = json.Unmarshal(inf, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

type EnvironmentName string

const (
	EnvironmentBigMirror       EnvironmentName = "BigMirrorEnvironment"
	EnvironmentBTS             EnvironmentName = "BTSEnvironment"
	EnvironmentCrabRave        EnvironmentName = "CrabRaveEnvironment"
	EnvironmentDefault         EnvironmentName = "DefaultEnvironment"
	EnvironmentDragons         EnvironmentName = "DragonsEnvironment"
	EnvironmentFitBeat         EnvironmentName = "FitBeatEnvironment"
	EnvironmentGlassDesert     EnvironmentName = "GlassDesertEnvironment"
	EnvironmentGreenDay        EnvironmentName = "GreenDayEnvironment"
	EnvironmentGreenDayGrenade EnvironmentName = "GreenDayGrenadeEnvironment"
	EnvironmentInterscope      EnvironmentName = "InterscopeEnvironment"
	EnvironmentKaleidoscope    EnvironmentName = "KaleidoscopeEnvironment"
	EnvironmentKDA             EnvironmentName = "KDAEnvironment"
	EnvironmentLinkinPark      EnvironmentName = "LinkinParkEnvironment"
	EnvironmentMonstercat      EnvironmentName = "MonstercatEnvironment"
	EnvironmentNice            EnvironmentName = "NiceEnvironment"
	EnvironmentOrigins         EnvironmentName = "OriginsEnvironment"
	EnvironmentPanic           EnvironmentName = "PanicEnvironment"
	EnvironmentRocket          EnvironmentName = "RocketEnvironment"
	EnvironmentTimbaland       EnvironmentName = "TimbalandEnvironment"
	EnvironmentTriangle        EnvironmentName = "TriangleEnvironment"
)

type Info struct {
	BPM float64 `json:"_beatsPerMinute"`

	BeatmapSets     []BeatmapSet    `json:"_difficultyBeatmapSets"`
	EnvironmentName EnvironmentName `json:"_environmentName"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (m *Map) ActiveDifficulty() *Difficulty {
	return m.activeDifficulty
}

func (m *Map) ActiveDifficultyProfile() *EnvProfile {
	return m.activeEnvironmentProfile
}

func (m *Map) SetActiveDifficulty(characteristic Characteristic, difficulty BeatmapDifficulty) error {
	var beatmapSet *BeatmapSet

	in := m.Info

	for i, set := range in.BeatmapSets {
		if set.Characteristic == characteristic {
			beatmapSet = &in.BeatmapSets[i]
			goto foundCharacteristic
		}
	}
	return fmt.Errorf("characteristic '%s' not found in info.dat", characteristic)
foundCharacteristic:

	var difficultyFilename string

	for _, beatmap := range beatmapSet.Beatmaps {
		if beatmap.Difficulty == difficulty {
			difficultyFilename = beatmap.Filename
			goto foundDifficulty
		}
	}
	return fmt.Errorf("difficulty '%s' not found in info.dat", difficulty)
foundDifficulty:

	difficultyPath := filepath.Join(m.workingDir, difficultyFilename)

	d, err := openDifficulty(difficultyPath)
	if err != nil {
		return err
	}

	m.activeDifficulty = d
	m.activeCharacteristic = characteristic
	m.activeBeatmapDifficulty = difficulty
	m.activeDifficultyPath = difficultyPath
	m.calculateBPMRegions()
	m.loadEnvironmentProfile()

	return nil
}

type bpmRegion struct {
	start     float64
	bpm       float64
	startBeat float64
}

func (m *Map) calculateBPMRegions() {
	startBPM := m.Info.BPM

	bpmChanges := m.ActiveDifficulty().CustomData.BPMChanges

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

	m.activeDifficultyBPMRegions = bpmRegions
}

func (m *Map) UnscaleTime(beat float64) Time {
	startBPM := m.Info.BPM
	bpmRegions := m.activeDifficultyBPMRegions

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

func (m *Map) loadEnvironmentProfile() {
	profile, err := LoadEnv(string(m.Info.EnvironmentName))
	if err != nil {
		return
	}
	m.activeEnvironmentProfile = profile
}

func (m *Map) AppendEvents(events []Event) error {
	m.activeDifficulty.Events = append(m.activeDifficulty.Events, events...)

	f, err := os.OpenFile(m.activeDifficultyPath, os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)

	err = enc.Encode(m.activeDifficulty)
	if err != nil {
		return err
	}

	return nil
}

func (m *Map) SaveEvents(events []Event) error {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Time < events[j].Time
	})

	m.activeDifficulty.Events = events

	f, err := os.OpenFile(m.activeDifficultyPath, os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)

	err = enc.Encode(m.activeDifficulty)
	if err != nil {
		return err
	}

	return nil
}

func (m *Map) Events() []Event {
	return m.activeDifficulty.Events
}

func (in *Info) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(in, "Extra", raw)
}

func (in Info) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(in, "Extra")
}

type Characteristic string

const (
	CharacteristicStandard  Characteristic = "Standard"
	CharacteristicNoArrows  Characteristic = "NoArrows"
	CharacteristicOneSaber  Characteristic = "OneSaber"
	Characteristic360Degree Characteristic = "360Degree"
	Characteristic90Degree  Characteristic = "90Degree"
	CharacteristicLightshow Characteristic = "Lightshow"
	CharacteristicLawless   Characteristic = "Lawless"
)

type BeatmapSet struct {
	Characteristic Characteristic `json:"_beatmapCharacteristicName"`
	Beatmaps       []Beatmap      `json:"_difficultyBeatmaps"`
}

type BeatmapDifficulty string

const (
	BeatmapDifficultyEasy       BeatmapDifficulty = "Easy"
	BeatmapDifficultyNormal     BeatmapDifficulty = "Normal"
	BeatmapDifficultyHard       BeatmapDifficulty = "Hard"
	BeatmapDifficultyExpert     BeatmapDifficulty = "Expert"
	BeatmapDifficultyExpertPlus BeatmapDifficulty = "ExpertPlus"
)

type Beatmap struct {
	Difficulty     BeatmapDifficulty `json:"_difficulty"`
	DifficultyRank int               `json:"_difficultyRank"`
	Filename       string            `json:"_beatmapFilename"`
	NJS            float64           `json:"_noteJumpMovementSpeed"`
	Offset         float64           `json:"_noteJumpStartBeatOffset"`
}

func openDifficulty(filename string) (*Difficulty, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var difficulty Difficulty

	err = json.Unmarshal(f, &difficulty)
	if err != nil {
		return nil, err
	}

	return &difficulty, nil
}

type Difficulty struct {
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

type Time float64

type NoteType int
type CutDirection int

type Note struct {
	Time         Time         `json:"_time"`
	LineIndex    int          `json:"_lineIndex"`
	LineLayer    int          `json:"_lineLayer"`
	Type         NoteType     `json:"_type"`
	CutDirection CutDirection `json:"_cutDirection"`

	Extra map[string]*json.RawMessage `json:"-"`
}

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

type BPMChange struct {
	Time            Time    `json:"_time"`
	BPM             float64 `json:"_BPM"`
	BeatsPerBar     int     `json:"_beatsPerBar"`
	MetronomeOffset int     `json:"_metronomeOffset"`

	Extra map[string]*json.RawMessage `json:"-"`
}
