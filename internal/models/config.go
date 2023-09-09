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

// TODO: Have not put in use yet
type (
	SavedConfigs    struct{}
	ConstantConfigs struct {
		RootDirPath    string
		HistoryDirName string
	}
)

type Configuration struct {
	Mode            Mode
	SavedConfigs    *SavedConfigs
	ConstantConfigs *ConstantConfigs
}
