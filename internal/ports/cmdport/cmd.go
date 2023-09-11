package cmdport

import "github.com/urfave/cli/v2"

type CMD interface {
	OSController(*cli.Context) error
	PgsController(*cli.Context) error
}
