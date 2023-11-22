package errs

import (
	"time"
)

const (
	Code_TEMPLATE_FORMAT_NOT_FOUND = iota + 1
	CODE_CONFIG_FILE_EXISTS
	API_AUTH
	API_RATE_LIMIT
	API_INTERNAL_ERROR
)

func TmplFormatNotFound(format string) Error {
	return Error{
		Msg:    `format "%s" not found`,
		Params: []any{format},
		Code:   Code_TEMPLATE_FORMAT_NOT_FOUND,
	}
}

func ConfGenFileExists(path string) Error {
	// return fmt.Errorf("file with path %s already exists, use -o flag", path)
	return Error{
		Msg:    `file with path "%s" already exists, use -o flag`,
		Params: []any{path},
		Code:   CODE_CONFIG_FILE_EXISTS,
	}
}

func OpenAPIInvalidAuth(message string) Error {
	return Error{
		Msg:    `openai API responded with 401 which means your authentication data is wrong, check your token in "~/.jarvis/.jarvisrc.json", reason: %s`,
		Params: []any{message},
		Code:   API_AUTH,
	}
}

func OpenAPIRateLimit(message string, waitTime time.Duration) Error {
	return Error{
		Msg:    `rate limit error reason: %s`,
		Params: []any{message},
		Data: map[string]any{
			"wait": waitTime,
		},
		Code: API_RATE_LIMIT,
	}
}

func OpenAPIInternalError(message string) Error {
	return Error{
		Msg:    `openai API internal error, retrying in 3 seconds, reason: %s`,
		Params: []any{message},
		Code:   API_INTERNAL_ERROR,
	}
}
