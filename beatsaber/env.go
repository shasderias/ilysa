package beatsaber

import (
	"embed"
	"encoding/json"
	"strconv"
)

//go:embed env/*.json
var fs embed.FS

type EnvPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type EnvColor struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
}

type EnvProfile struct {
	ColorScheme EnvColorScheme        `json:"colorScheme"`
	Props       map[string][]EnvProp  `json:"props"`
	PropGroup   map[int]map[int][]int `json:"-"`
}

type EnvColorScheme struct {
	ColorLeft          EnvColor `json:"_colorLeft"`
	ColorRight         EnvColor `json:"_colorRight"`
	EnvColorLeft       EnvColor `json:"_envColorLeft"`
	EnvColorRight      EnvColor `json:"_envColorRight"`
	ObstacleColor      EnvColor `json:"_obstacleColor"`
	EnvColorLeftBoost  EnvColor `json:"_envColorLeftBoost"`
	EnvColorRightBoost EnvColor `json:"_envColorRightBoost"`
}

type EnvProp struct {
	Type     int         `json:"_type"`
	PropID   int         `json:"_propId"`
	LightID  int         `json:"_lightId"`
	Name     string      `json:"name"`
	Position EnvPosition `json:"position"`
}

func LoadEnv(envName string) (*EnvProfile, error) {
	f, err := fs.Open("env/" + envName + ".json")
	if err != nil {
		return nil, err
	}

	var info EnvProfile

	dec := json.NewDecoder(f)
	err = dec.Decode(&info)
	if err != nil {
		return nil, err
	}

	info.PropGroup = make(map[int]map[int][]int)

	for eventTypeStr, prop := range info.Props {
		eventType, err := strconv.Atoi(eventTypeStr)
		if err != nil {
			continue
		}

		if _, ok := info.PropGroup[eventType]; !ok {
			info.PropGroup[eventType] = make(map[int][]int)
		}

		for lightIDMinusOne, p := range prop {
			lightID := lightIDMinusOne + 1
			if _, ok := info.PropGroup[eventType][p.PropID]; !ok {
				info.PropGroup[eventType][p.PropID] = make([]int, 0, 2)
			}
			info.PropGroup[eventType][p.PropID] = append(info.PropGroup[eventType][p.PropID], lightID)
		}
	}

	return &info, nil
}

func (p *EnvProfile) MaxLightID(eventType EventType) int {
	if p == nil {
		return -1
	}
	lights, ok := p.Props[strconv.Itoa(int(eventType))]
	if !ok {
		return -1
	}
	return len(lights) // lightIDs are 1-indexed
}
