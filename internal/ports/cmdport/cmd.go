package cmdport

import "github.com/urfave/cli/v2"

type CMD interface {
	BootstrapConfig(ctx *cli.Context) error

	OSController(*cli.Context) error
	PgsController(*cli.Context) error
}
