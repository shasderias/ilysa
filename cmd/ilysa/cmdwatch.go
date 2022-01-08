package main

import (
	"context"
	"flag"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/shasderias/ilysa/internal/refresh"
)

func makeWatchCmd() *ffcli.Command {
	fs := flag.NewFlagSet("watch", flag.ExitOnError)

	watchCmd := &ffcli.Command{
		Name:       "watch",
		ShortUsage: "ilysa watch <directory>",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			if fs.NArg() == 0 || fs.NArg() >= 2 {
				return flag.ErrHelp
			}

			config := &refresh.Configuration{
				Debug:      false,
				ProjectDir: fs.Arg(0),
			}

			return refresh.New(config).Start()
		},
	}

	return watchCmd
}
