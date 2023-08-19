package os

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

type mode int8

const (
	AppendMode = iota + 1
	TruncMode
	CreateMode
	ReadMode
)

var modeSets = map[mode]int{
	AppendMode: os.O_APPEND | os.O_WRONLY | os.O_CREATE,
	TruncMode:  os.O_APPEND | os.O_WRONLY | os.O_TRUNC,
	CreateMode: os.O_APPEND | os.O_WRONLY | os.O_CREATE,
	ReadMode:   os.O_RDONLY,
}

var (
	Stdout = os.Stdout
	Stderr = os.Stderr
)

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

func OpenFile(p string, m mode) (*os.File, error) {
	if !path.IsAbs(p) {
		p = FromRoot(p)
	}

	if err := os.MkdirAll(path.Join(p, ".."), os.ModePerm); err != nil {
		return nil, err
	}

	return os.OpenFile(filepath.Clean(p), modeSets[m], 0o600)
}

func RemoveDir(p string) error {
	return os.RemoveAll(p)
}
