package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/zgfzgf/min-lotus/build"
	lcli "github.com/zgfzgf/min-lotus/cli"
	"github.com/zgfzgf/min-lotus/daemon"
)

func main() {
	local := []*cli.Command{
		daemon.Cmd,
	}

	app := &cli.App{
		Name:    "lotus",
		Usage:   "Filecoin decentralized storage network client",
		Version: build.Version,

		Commands: append(local, lcli.Commands...),
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		return
	}
}
