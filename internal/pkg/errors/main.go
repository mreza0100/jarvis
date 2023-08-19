package errors

import "fmt"

type Code int32

type error struct {
	code   Code
	msg    string
	params []any
}

func (e error) Error() string {
	return fmt.Sprintf(e.msg, e.params...)
}
