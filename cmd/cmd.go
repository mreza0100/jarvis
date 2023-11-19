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

	interactive := &cli.Command{
		Name:        "interactive",
		Aliases:     []string{"i"},
		Description: "jarvis interactive terminal interactive models",
		Subcommands: []*cli.Command{
			pgsCliConf, osCliConf,
		},
	}

	bootstrap := &cli.Command{
		Name:        "bootstrap",
		Aliases:     []string{},
		Description: "initiate & create jarvis configuration on host",
		Action:      handlers.BootstrapConfig,
		Flags: []cli.Flag{&cli.BoolFlag{
			Name:  "trunk",
			Value: false,
		}},
	}

	configuration := &cli.Command{
		Name:        "config",
		Aliases:     []string{"c"},
		Description: "jarvis config controller root command entry",
		Subcommands: []*cli.Command{bootstrap},
	}

	app := &cli.App{
		Commands: []*cli.Command{
			interactive,
			configuration,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
