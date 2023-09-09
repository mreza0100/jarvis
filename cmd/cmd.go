package cmd

import (
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driving/cmd"
	cli "github.com/urfave/cli/v2"
)

func Main() {
	handlers := cmd.NewCMDHandler(&cmd.CmdHandlerParams{})

	// TODO: Get this cli configs from the adaper itself
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "interactive",
				Aliases:     []string{"i"},
				Description: "jarvis interactive terminal",
				Flags:       []cli.Flag{},
				Action:      handlers.Interactive,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
