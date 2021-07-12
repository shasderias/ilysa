package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
)

const mapPath = `D:\Beat Saber Data\CustomWIPLevels\Magnet`

func main() {
	if err := do(); err != nil {
		fmt.Println(err)
	}
}

func do() error {
	magnet, err := beatsaber.Open(mapPath)
	if err != nil {
		return err
	}

	p := ilysa.New(magnet)

	err = p.Map.SetActiveDifficulty(beatsaber.CharacteristicStandard, beatsaber.BeatmapDifficultyExpertPlus)
	if err != nil {
		return err
	}

	//sb := NewSandbox(p, 4)
	//sb.Play1()

	//leadIn := NewLeadInOut(p, 4)
	//leadIn.PlayIn()
	//
	//intro1 := NewIntro(p, 16)
	//intro1.Play1()
	//
	//verse1a := NewVerseA(p, 52)
	//verse1a.Play1()
	//
	//verse1b := NewVerseB(p, 84)
	//verse1b.Play1()
	//
	//chorus := NewChorus(p, 114)
	//chorus.Play1()
	//
	//breakdown := NewBreakdown(p, 149)
	//breakdown.Play()
	//
	//verse2a := NewVerseA(p, 164)
	//verse2a.Play2()
	//
	//verse2b := NewVerseB(p, 196)
	//verse2b.Play1()
	//
	//chorus2 := NewChorus(p, 226)
	//chorus2.Play1()
	//
	//guitarSolo := NewGuitarSolo(p, 260)
	//guitarSolo.Play()
	//
	//verse3 := NewVerseC(p, 292)
	//verse3.Play()
	//
	//chorus3 := NewChorus(p, 326)
	//chorus3.Play3()
	//
	//chorus4 := NewChorus(p, 359)
	//chorus4.Play4()
	//
	//outro := NewIntro(p, 391)
	//outro.Play2()
	//
	//leadout := NewLeadInOut(p, 423)
	//leadout.PlayOut()

	return p.Save()
}
