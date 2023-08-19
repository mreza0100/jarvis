package cmdport

import "github.com/urfave/cli/v2"

type CMD interface {
	Boot(ctx *cli.Context) error
}
