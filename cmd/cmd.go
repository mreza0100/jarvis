package cmd

import (
	"os"

	"github.com/mreza0100/jarvis/internal/adapters/driving/cmd"
	cli "github.com/urfave/cli/v2"
)

func Main() {
	handlers := cmd.NewCMDHandler(&cmd.CmdHandlerParams{})

	pgsCliConf := &cli.Command{
		Name:        "postgres",
		Aliases:     []string{"pgs"},
		Description: "jarvis interactive Postgres",
		Action:      handlers.PgsController,
		Flags:       []cli.Flag{},
	}

	osCliConf := &cli.Command{
		Name:        "os",
		Aliases:     []string{},
		Description: "jarvis interactive terminal",
		Action:      handlers.OSController,
		Flags:       []cli.Flag{},
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "interactive",
				Aliases:     []string{"i"},
				Description: "jarvis interactive terminal interactive models",
				Subcommands: []*cli.Command{
					pgsCliConf, osCliConf,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
