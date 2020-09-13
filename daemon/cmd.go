package daemon

import (
	goctx "context"
	"github.com/urfave/cli/v2"

	"github.com/zgfzgf/min-lotus/node"
)

var Cmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Action: func(context *cli.Context) error {
		ctx := goctx.Background()

		api, err := node.New(ctx)
		if err != nil {
			return err
		}

		return serveRPC(api)
	},
}
