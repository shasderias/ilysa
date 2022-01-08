package main

import (
	"fmt"
	"os"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func do() error {
	m, err := beatsaber.Open(mapDirectory)
	if err != nil {
		return err
	}
	diff, err := m.OpenDifficulty(characteristic, difficulty)
	if err != nil {
		return err
	}

	ctx := ilysa.New(diff)

	err = Ilysa(ctx)
	if err != nil {
		return err
	}

	return ctx.Save()
}
