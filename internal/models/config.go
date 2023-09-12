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

type Configuration struct {
	RootDirName    string
	HistoryDirName string
	Token          string
	Mode           Mode

	ConfigFile *ConfigFile
}

type ConfigFile struct {
	PostgresConfig PostgresConfig `json:"postgres"`
}
