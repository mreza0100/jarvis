package os

import (
	"path/filepath"
	"runtime"
)

func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)
	path := filepath.Join(dir, "../..")

	return path
}

func FromRoot(path string) string {
	return filepath.Join(GetRootPath(), path)
}
