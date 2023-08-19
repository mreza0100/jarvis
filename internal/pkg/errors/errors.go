package errors

const (
	Code_TmplFormatNotFound = iota + 1
	Code_ConfGenFileExists
)

func TmplFormatNotFound(format string) error {
	return error{
		msg:    `format "%s" not found`,
		params: []any{format},
		code:   Code_TmplFormatNotFound,
	}
}

func ConfGenFileExists(path string) error {
	// return fmt.Errorf("file with path %s already exists, use -o flag", path)
	return error{
		msg:    `file with path "%s" already exists, use -o flag`,
		params: []any{path},
		code:   Code_ConfGenFileExists,
	}
}
