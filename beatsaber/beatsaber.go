package beatsaber

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Map struct {
	workingDir string

	Info *Info
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

func (m *Map) OpenDifficulty(characteristic Characteristic, difficulty BeatmapDifficulty) (*Difficulty, error) {
	var beatmapSet *BeatmapSet

	in := m.Info

	for i, set := range in.BeatmapSets {
		if set.Characteristic == characteristic {
			beatmapSet = &in.BeatmapSets[i]
			goto foundCharacteristic
		}
	}
	return nil, fmt.Errorf("characteristic '%s' not found in info.dat", characteristic)
foundCharacteristic:

	var difficultyFilename string

	for _, beatmap := range beatmapSet.Beatmaps {
		if beatmap.Difficulty == difficulty {
			difficultyFilename = beatmap.Filename
			goto foundDifficulty
		}
	}
	return nil, fmt.Errorf("difficulty '%s' not found in info.dat", difficulty)
foundDifficulty:

	difficultyPath := filepath.Join(m.workingDir, difficultyFilename)

	f, err := os.ReadFile(difficultyPath)
	if err != nil {
		return nil, err
	}

	var diff Difficulty

	err = json.Unmarshal(f, &diff)
	if err != nil {
		return nil, err
	}

	diff.info = m.Info
	diff.filepath = difficultyPath
	diff.calculateBPMRegions()

	return &diff, nil
}
