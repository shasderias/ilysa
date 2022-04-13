package beatsaber

type DifficultyVersion string

const (
	DifficultyVersionNil    DifficultyVersion = ""
	DifficultyVersion_2_2_0                   = "2.2.0"
	DifficultyVersion_2_5_0                   = "2.5.0"
	DifficultyVersion_2_6_0                   = "2.6.0"
	DifficultyVersion_3_0_0                   = "3.0.0"
)

// DifficultyVersions are the known difficulty versions
var DifficultyVersions = []string{
	DifficultyVersion_2_2_0,
	DifficultyVersion_2_5_0,
	DifficultyVersion_2_6_0,
	DifficultyVersion_3_0_0,
}

func NewDifficultyVersion(s string) DifficultyVersion {
	for _, v := range DifficultyVersions {
		if v == s {
			return DifficultyVersion(s)
		}
	}
	return DifficultyVersionNil
}

var SupportedDifficultyVersions = map[DifficultyVersion]DifficultyVersionSupport{
	DifficultyVersion_2_2_0: {
		OpenFunc: OpenDifficultyV220,
	},
	DifficultyVersion_2_5_0: {
		OpenFunc: OpenDifficultyV250,
	},
}

type DifficultyVersionSupport struct {
	OpenFunc func(info *Info, path string) (Difficulty, error)
}
