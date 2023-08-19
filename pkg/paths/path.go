package paths

import "strings"

func GetFormat(path string) string {
	splt := strings.Split(path, ".")
	return splt[len(splt)-1]
}
