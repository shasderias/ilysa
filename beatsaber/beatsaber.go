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

func (m *Map) OpenDifficulty(characteristic Characteristic, difficulty BeatmapDifficulty) (Difficulty, error) {
	var beatmapSet *BeatmapSet

	infoDat := m.Info

	for i, set := range infoDat.BeatmapSets {
		if set.Characteristic == characteristic {
			beatmapSet = &infoDat.BeatmapSets[i]
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

	version, err := parseDifficultyVersion(difficultyPath)
	if err != nil {
		return nil, err
	}

	versionSupport, ok := SupportedDifficultyVersions[version]
	if !ok {
		return nil, fmt.Errorf("unsupported difficulty version '%s'", version)
	}

	return versionSupport.OpenFunc(infoDat, difficultyPath)
}

func parseDifficultyVersion(path string) (DifficultyVersion, error) {
	type diffVersionOnly struct {
		Version string `json:"_version"`
	}

	f, err := os.Open(path)
	if err != nil {
		return DifficultyVersionNil, err
	}

	j := json.NewDecoder(f)

	var dvo diffVersionOnly

	if err := j.Decode(&dvo); err != nil {
		return DifficultyVersionNil, err
	}

	return NewDifficultyVersion(dvo.Version), nil
}
