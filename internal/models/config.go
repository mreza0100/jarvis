package models

type Mode uint8

func (m Mode) IsDev() bool {
	return m == dev
}

const (
	dev Mode = iota
	prod
)

var Modes = struct {
	Dev  Mode
	Prod Mode
}{
	Dev:  dev,
	Prod: prod,
}

type SavedConfigs struct{}

type Configuration struct {
	Mode         Mode
	SavedConfigs *SavedConfigs
}
