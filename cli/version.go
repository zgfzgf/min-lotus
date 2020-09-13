package cli

import (
	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(context *cli.Context) error {
		// TODO: print more useful things

		cli.VersionPrinter(context)
		return nil
	},
}
