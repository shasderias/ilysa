package main

import (
	"context"
	"flag"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/shasderias/ilysa"
)

func makeRunCmd() *ffcli.Command {
	fs := flag.NewFlagSet("run", flag.ExitOnError)

	runCmd := &ffcli.Command{
		Name:       "run",
		ShortUsage: "ilysa run <directory>",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			if fs.NArg() == 0 || fs.NArg() >= 2 {
				return flag.ErrHelp
			}

			return ilysa.Invoke(fs.Arg(0))
		},
	}

	return runCmd
}
