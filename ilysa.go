package ilysa

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"sort"
	"strings"

	"github.com/shasderias/ilysa/colorful"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/internal/imath"
	"github.com/shasderias/ilysa/internal/swallowjson"
)

type Project struct {
	Map
	context.Context

	events []evt.Event
}

type Map interface {
	ActiveDifficultyProfile() *beatsaber.EnvProfile
	AppendEvents([]beatsaber.Event) error
	Events() []beatsaber.Event
	SaveEvents([]beatsaber.Event) error
	SetActiveDifficulty(c beatsaber.Characteristic, difficulty beatsaber.BeatmapDifficulty) error
	UnscaleTime(beat float64) beatsaber.Time
}

func New(bsMap Map) *Project {
	p := &Project{
		Map:    bsMap,
		events: []evt.Event{},
	}

	p.Context = context.Base(p)

	return p
}

func (p *Project) BOffset(offset float64) context.Context {
	return context.Base(p).BOffset(offset)
}

func (p *Project) MaxLightID(t evt.LightType) int {
	return p.Map.ActiveDifficultyProfile().MaxLightID(beatsaber.EventType(t))
}

func (p *Project) AddEvents(events ...evt.Event) {
	p.events = append(p.events, events...)
}

func (p *Project) Events() *[]evt.Event {
	return &p.events
}

func (p *Project) FilterEvents(f func(e evt.Event) bool) *[]evt.Event {
	filteredEvents := make([]evt.Event, 0)

	for _, e := range p.events {
		if f(e) {
			filteredEvents = append(filteredEvents, e)
		}
	}

	p.events = filteredEvents

	return &filteredEvents
}

func (p *Project) MapEvents(f func(e evt.Event) evt.Event) {
	for i := range p.events {
		p.events[i] = f(p.events[i])
	}
}

func (p *Project) sortEvents() {
	sort.Slice(p.events, func(i, j int) bool {
		return p.events[i].Beat() < p.events[j].Beat()
	})
}

func (p *Project) generateBeatSaberEvents() ([]beatsaber.Event, error) {
	events := []beatsaber.Event{}

	for _, e := range p.events {
		roundedTime := beatsaber.Time(imath.Round(float64(p.Map.UnscaleTime(e.Beat())), 5))
		event := beatsaber.Event{
			Time:  roundedTime,
			Type:  e.Type(),
			Value: e.Value(),
		}

		cd, err := e.CustomData()
		if err != nil {
			return nil, err
		}
		event.CustomData = cd

		events = append(events, event)
	}

	return events, nil
}

func (p *Project) Save() error {
	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(events))

	return p.Map.SaveEvents(events)
}

// GenerateGradientReference generates a PNG with all gradients declared at the
// top level and saves it to path. Each gradient is labelled with its name.
// Useful as a reference.
//
// GenerateGradientReference only works on the machine the program is compiled on.
func (p *Project) GenerateGradientReference(path string) error {
	const (
		topBottomPadding    = 16
		leftRightPadding    = 16
		textColWidth        = 160
		gradientHeight      = 36
		gradientWidth       = 480
		textGradientSpacing = 16
		gradientSpacing     = 8
	)

	var (
		backgroundColor = colorful.Hex("#333")
	)

	grads := gradient.DeclaredGradients
	l := len(grads)

	h := topBottomPadding + (gradientHeight * l) + (gradientSpacing * (l - 1)) + topBottomPadding
	w := leftRightPadding + textColWidth + textGradientSpacing + gradientWidth + leftRightPadding

	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	draw.Draw(img, image.Rect(0, 0, w, h), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	i := 0

	type namedGrad struct {
		name string
		grad gradient.Table
	}
	sortedGrads := []namedGrad{}

	for name, grad := range grads {
		sortedGrads = append(sortedGrads, namedGrad{
			name: name,
			grad: grad,
		})
	}

	sort.Slice(sortedGrads, func(i, j int) bool {
		return strings.Compare(sortedGrads[i].name, sortedGrads[j].name) < 0
	})

	for _, sg := range sortedGrads {
		name := sg.name
		grad := sg.grad

		// label
		x0 := leftRightPadding
		y0 := topBottomPadding + ((gradientHeight + gradientSpacing) * i)

		addLabel(img, x0, y0+(gradientHeight/2), name)

		// gradient
		x0 = leftRightPadding + textColWidth + textGradientSpacing
		y1 := y0 + gradientHeight

		for x := 0; x < gradientWidth; x++ {
			c := grad.Lerp(float64(x) / float64(gradientWidth))
			draw.Draw(img, image.Rect(x0+x, y0, x0+x+1, y1), &image.Uniform{c}, image.Point{}, draw.Src)
		}

		i++
	}
	outpng, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outpng.Close()

	return png.Encode(outpng, img)
}

func addLabel(img *image.NRGBA, x, y int, label string) {
	col := colorful.Hex("#fff")
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

type EventCustomData struct {
	IlysaDev bool                   `json:"ilysaDev"`
	Rest     map[string]interface{} `json:"-"`
}

func (p *Project) getNonIlysaEvents() ([]beatsaber.Event, error) {
	mapEvents := p.Map.Events()

	fmt.Printf("map has %d events\n", len(mapEvents))

	niEvents := []beatsaber.Event{}

	for _, e := range mapEvents {
		if len(e.CustomData) == 0 {
			niEvents = append(niEvents, e)
			continue
		}

		cd := EventCustomData{}
		if err := json.Unmarshal(e.CustomData, &cd); err != nil {
			return nil, err
		}
		if !cd.IlysaDev {
			niEvents = append(niEvents, e)
		}
	}

	fmt.Printf("map has %d non-Ilysa events\n", len(niEvents))

	return niEvents, nil
}

func (p *Project) SaveDev() error {
	newEvents, err := p.getNonIlysaEvents()
	if err != nil {
		return err
	}

	generatedEvents, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(generatedEvents))

	for i, e := range generatedEvents {
		cd := EventCustomData{}
		if err := swallowjson.UnmarshalWith(&cd, "Rest", e.CustomData); err != nil {
			return err
		}
		cd.IlysaDev = true
		newRawMessage, err := swallowjson.MarshalWith(cd, "Rest")
		if err != nil {
			return err
		}

		generatedEvents[i].CustomData = newRawMessage
	}

	newEvents = append(newEvents, generatedEvents...)

	fmt.Printf("saving %d events\n", len(newEvents))

	return p.Map.SaveEvents(newEvents)
}

func (p *Project) SaveProd() error {
	niEvents, err := p.getNonIlysaEvents()
	if err != nil {
		return err
	}

	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(events))

	newEvents := append(niEvents, events...)

	fmt.Printf("saving %d events\n", len(newEvents))

	return p.Map.SaveEvents(newEvents)
}

func (p *Project) Append() error {
	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(events))

	return p.Map.AppendEvents(events)
}

func (p *Project) Dump() error {
	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	eventsJSON, err := json.Marshal(events)
	if err != nil {
		return err
	}

	fmt.Print(string(eventsJSON))

	return nil
}
