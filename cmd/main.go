package cmd

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

func Main() {
	cmdHandlers := getCMDHandlers()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "boot",
				Aliases:     []string{},
				Description: "",
				Flags:       []cli.Flag{},
				Action:      cmdHandlers.Boot,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
