package templatestore

import "embed"

//go:embed *.gpt
var ModelsFS embed.FS
