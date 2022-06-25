package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
)

const (
	mapDirectory   = ""
	characteristic = beatsaber.CharacteristicStandard
	difficulty     = beatsaber.BeatmapDifficultyExpertPlus
)

var (
	buildTargets = []ilysa.BuildTarget{
		{beatsaber.CharacteristicStandard, beatsaber.BeatmapDifficultyExpertPlus},
		{beatsaber.CharacteristicStandard, beatsaber.BeatmapDifficultyExpert},
	}
)

func IlysaMain(ctx context.Context) error {
	return nil
}
