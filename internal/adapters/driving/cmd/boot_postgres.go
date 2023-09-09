package cmd

// func readPostgresConfig(path string) (*runnerport.PostgresConfig, error) {
// 	rawContent, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	configs := &runnerport.PostgresConfig{}
// 	return configs, json.Unmarshal(rawContent, configs)
// }

// func (c *cmd) bootPostgres(ctx *cli.Context) error {
// 	configFilePath := ctx.Args().Get(1)
// 	configs, err := readPostgresConfig(configFilePath)
// 	if err != nil {
// 		return err
// 	}

// 	cfgProvider := config.NewConfigProvider()
// 	history := history.NewHistory(cfgProvider)
// 	runner := postgres.NewPostgresRunner(&postgres.PostgresRunnerReq{
// 		Configs: configs,
// 	})

// 	chat := chat.NewChat(&chat.NewChatReq{
// 		Clinet: openai.NewClient("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"),
// 	})
// 	interactor := interactor.NewInteractor(interactor.InteractorArg{
// 		CfgProvider: cfgProvider,

// 		Stdin:  os.Stdin,
// 		Stdout: os.Stdout,
// 		Stderr: os.Stderr,
// 	})
// 	bootSrv := services.NewBootSrv(&srvport.ServicesReq{
// 		Chat:       chat,
// 		Runner:     runner,
// 		Interactor: interactor,
// 		History:    history,
// 	})

// 	template, err := c.getTemplate("postgres.gpt")
// 	if err != nil {
// 		return err
// 	}

// 	return bootSrv.Start(template)
// }
