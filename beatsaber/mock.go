package beatsaber

import (
	"encoding/json"
	"fmt"
	"math"
)

type MockMap struct {
	bpm        float64
	bpmRegions []bpmRegion
	envProfile *EnvProfile
}

func NewMockMap(envName EnvironmentName, bpm float64, bpmRegionsJSON string) (MockMap, error) {
	mockMap := MockMap{
		bpm: bpm,
	}

	bpmChanges := []BPMChange{}
	if err := json.Unmarshal([]byte(bpmRegionsJSON), &bpmChanges); err != nil {
		return MockMap{}, err
	}

	mockMap.calculateBPMRegions(bpmChanges)

	envProfile, err := LoadEnv(string(envName))
	if err != nil {
		return MockMap{}, err
	}

	mockMap.envProfile = envProfile

	return mockMap, nil
}

func (m *MockMap) calculateBPMRegions(bpmChanges []BPMChange) {
	startBPM := m.bpm

	bpmRegions := []bpmRegion{{
		0,
		startBPM,
		0,
	}}

	for i, change := range bpmChanges {
		t := float64(change.Time)
		prevRegion := bpmRegions[i]

		bpmRegions = append(bpmRegions, bpmRegion{
			float64(change.Time),
			change.BPM,
			math.Ceil((t-prevRegion.start)*startBPM/change.BPM) + prevRegion.startBeat,
		})
	}

	m.bpmRegions = bpmRegions
}

func (m MockMap) UnscaleTime(beat float64) Time {
	startBPM := m.bpm
	bpmRegions := m.bpmRegions

	for i := len(bpmRegions) - 1; i >= 0; i-- {
		region := bpmRegions[i]
		if beat >= region.startBeat {
			diff := beat - region.startBeat

			return Time(region.start + (diff / region.bpm * startBPM))
		}
	}

	panic(fmt.Sprintf("UnscaleTime(): unreachable code, beat: %f", beat))
}

func (MockMap) SaveEvents(events []Event) error {
	return nil
}

func (m MockMap) ActiveDifficulty() *Difficulty {
	return nil
}

func (m MockMap) ActiveDifficultyProfile() *EnvProfile {
	return m.envProfile
}

func (m MockMap) SetActiveDifficulty(c Characteristic, d BeatmapDifficulty) error {
	return nil
}
