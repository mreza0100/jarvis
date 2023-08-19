package configs

type Configs struct {
	App *App `yaml:"app"`
}

type App struct {
	AppName string `yaml:"app_name"`
}
