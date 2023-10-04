package models

import (
	"os"
)

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
	Mode           Mode

	ConfigFile *ConfigFile
}

type ChatConfig struct {
	Token struct {
		EnvName string `json:"env_name"`
		Value   string `json:"value"`
	} `json:"token"`
	Model       string  `json:"model"`
	Temperature float32 `json:"temperature"`
}

func (c *ChatConfig) GetToken() string {
	if c.Token.Value != "" {
		return c.Token.Value
	}

	return os.Getenv(c.Token.EnvName)
}

type ConfigFile struct {
	Postgres *PostgresConfig `json:"postgres"`
	OS       *OSConfig       `json:"os"`
}
