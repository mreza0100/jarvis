package cmdport

import "github.com/urfave/cli/v2"

type CMD interface {
	Interactive(ctx *cli.Context) error
}
