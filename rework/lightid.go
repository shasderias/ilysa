package rework

type LightID []int

func NewLightID(ids ...int) LightID {
	return append(LightID{}, ids...)
}

func NewLightIDFromInterval(startID, endID int) LightID {
	lightID := make(LightID, 0, endID-startID+1)
	for i := startID; i <= endID; i++ {
		lightID = append(lightID, i)
	}
	return lightID
}
