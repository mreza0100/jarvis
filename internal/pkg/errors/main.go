package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

type Code int32

type Error struct {
	Code   Code
	Msg    string
	Params []any
	Data   map[string]any
}

var (
	As   = errors.As
	Wrap = errors.Wrap
)

func (e Error) Error() string {
	return fmt.Sprintf(e.Msg, e.Params...)
}
