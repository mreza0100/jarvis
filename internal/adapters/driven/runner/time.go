package runner

import (
	"strings"
	"time"
)

func getTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("05-15-04-02_01_2006")
	return makePathSafe(formattedTime)
}

func makePathSafe(s string) string {
	// Replace characters that are not suitable for filenames
	s = strings.ReplaceAll(s, ":", "-")
	s = strings.ReplaceAll(s, "/", "-")
	s = strings.ReplaceAll(s, " ", "_")
	return s
}
