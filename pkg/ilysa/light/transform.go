package light

import "ilysa/pkg/ilysa"

func Divide(l ilysa.Light) ilysa.Light{
	bl, ok := l.(ilysa.BasicLight)
	if !ok { return nil }


	maxLightID := bl.project.ActiveDifficultyProfile().LightIDMax(l.eventType)
	return ilysa.CompoundLight{
		project:   l.project,
		eventType: l.eventType,
		set:       tfer(NewLightIDFromInterval(1, maxLightID)),
	}
}
