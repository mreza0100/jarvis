package os

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
)

type mode int8

const (
	AppendMode mode = iota + 1
	TruncMode
	CreateMode
	ReadMode
	ReadCreateMode
)

var modeSets = map[mode]int{
	AppendMode:     os.O_APPEND | os.O_WRONLY | os.O_CREATE,
	TruncMode:      os.O_APPEND | os.O_WRONLY | os.O_TRUNC,
	CreateMode:     os.O_APPEND | os.O_WRONLY | os.O_CREATE,
	ReadCreateMode: os.O_RDONLY | os.O_CREATE,
	ReadMode:       os.O_RDONLY,
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

func ReadFile(p string, m mode) ([]byte, error) {
	f, err := OpenFile(p, m)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}

func OpenFile(p string, m mode) (*os.File, error) {
	dirPath := path.Join(p, "..")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}

	return os.OpenFile(filepath.Clean(p), modeSets[m], 0o600)
}
