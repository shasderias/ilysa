package main

import (
	"fmt"
	"os"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
)

const (
	cmdRun = iota
	cmdBuild
)

func main() {
	cmd := doRun

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "run":
			cmd = doRun
		case "build":
			cmd = doBuild
		}
	}

	err := cmd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func doRun() error {
	m, err := beatsaber.Open(mapDirectory)
	if err != nil {
		return err
	}
	diff, err := m.OpenDifficulty(characteristic, difficulty)
	if err != nil {
		return err
	}

	ctx := ilysa.New(diff)

	err = IlysaMain(ctx)
	if err != nil {
		return err
	}

	return ctx.Save()
}

func doBuild() error {
	m, err := beatsaber.Open(mapDirectory)
	if err != nil {
		return err
	}
	diff, err := m.OpenDifficulty(characteristic, difficulty)
	if err != nil {
		return err
	}

	ctx := ilysa.New(diff)

	err = IlysaMain(ctx)
	if err != nil {
		return err
	}

	for _, target := range buildTargets {
		d, err := m.OpenDifficulty(target.Characteristic, target.BeatmapDifficulty)
		if err != nil {
			return err
		}
		dctx := ilysa.New(d)
		*dctx.Events() = *ctx.Events()
		if err := dctx.Save(); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "saved %d events to characterstic %s difficulty %s\n",
			len(*ctx.Events()), target.Characteristic, target.BeatmapDifficulty)
	}

	return nil
}
