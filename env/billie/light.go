package billie

import (
	"fmt"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/lightid"
)

type WaterLight struct {
}

func NewWaterLight() *WaterLight {
	return &WaterLight{}
}

func (w WaterLight) GenerateEvents(ctx context.LightContext) evt.Events {

	return evt.NewEvents(
		evt.NewChromaLighting(ctx, evt.OptType(Water1)),
		evt.NewChromaLighting(ctx, evt.OptType(Water2)),
		evt.NewChromaLighting(ctx, evt.OptType(Water3)),
		evt.NewChromaLighting(ctx, evt.OptType(Water4)),
	)
}

func (w WaterLight) LightLen() int {
	return 1
}

func (w WaterLight) Name() []string {
	return []string{"WaterLight"}
}

type CompositeWaterLight struct {
	s lightid.Set
}

func (c CompositeWaterLight) GenerateEvents(ctx context.LightContext) evt.Events {
	events := evt.NewEvents()
	lightID := c.s.Index(ctx.LightOrdinal())
	for _, id := range lightID {
		var typ evt.Type
		switch id {
		case 1:
			typ = Water1
		case 2:
			typ = Water2
		case 3:
			typ = Water3
		case 4:
			typ = Water4
		default:
			continue
		}
		events.Add(evt.NewChromaLighting(ctx, evt.OptType(typ)))
	}
	return events
}

func (c CompositeWaterLight) LightIDSetTransform(tfer func(lightid.Set) lightid.Set) context.Light {
	return CompositeWaterLight{
		s: tfer(c.s),
	}
}

func (c CompositeWaterLight) LightLen() int {
	return c.s.Len()
}

func (c CompositeWaterLight) Name() []string {
	return []string{
		fmt.Sprintf("CompositeWaterLight-%v", c.s),
	}
}

func (w WaterLight) LightIDTransform(tfer func(lightid.ID) lightid.Set) context.Light {
	s := tfer(lightid.NewFromInterval(1, 4))
	return &CompositeWaterLight{s}
}

func NewLeftSunbeamLight() light.Basic {
	return light.NewBasic(LeftSunbeam, 11)
}

func NewRightSunbeamLight() light.Basic {
	return light.NewBasic(RightSunbeam, 11)
}

func NewSunLight() light.Basic {
	return light.NewBasic(Sun, 1)
}

func NewLeftBottomLasersLight() light.Basic {
	return light.NewBasic(LeftBottomLasers, 9)
}

func NewRightBottomLasersLight() light.Basic {
	return light.NewBasic(RightBottomLasers, 9)
}
