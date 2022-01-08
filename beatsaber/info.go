package beatsaber

import (
	"encoding/json"
	"os"

	"github.com/shasderias/ilysa/internal/swallowjson"
)

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

type Info struct {
	BPM float64 `json:"_beatsPerMinute"`

	BeatmapSets     []BeatmapSet    `json:"_difficultyBeatmapSets"`
	EnvironmentName EnvironmentName `json:"_environmentName"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (in *Info) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(in, "Extra", raw)
}

func (in Info) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(in, "Extra")
}

type BeatmapSet struct {
	Characteristic Characteristic `json:"_beatmapCharacteristicName"`
	Beatmaps       []Beatmap      `json:"_difficultyBeatmaps"`
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

type Beatmap struct {
	Difficulty     BeatmapDifficulty `json:"_difficulty"`
	DifficultyRank int               `json:"_difficultyRank"`
	Filename       string            `json:"_beatmapFilename"`
	NJS            float64           `json:"_noteJumpMovementSpeed"`
	Offset         float64           `json:"_noteJumpStartBeatOffset"`
}

type BeatmapDifficulty string

const (
	BeatmapDifficultyEasy       BeatmapDifficulty = "Easy"
	BeatmapDifficultyNormal     BeatmapDifficulty = "Normal"
	BeatmapDifficultyHard       BeatmapDifficulty = "Hard"
	BeatmapDifficultyExpert     BeatmapDifficulty = "Expert"
	BeatmapDifficultyExpertPlus BeatmapDifficulty = "ExpertPlus"
)
