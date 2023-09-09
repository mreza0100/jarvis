package services

// type pgsInteractiveSrv struct {
// 	clinet             *openai.Client
// 	scriptCrashedTimes int

// 	chat       chatport.Chat
// 	interactor interactorport.Interactor
// 	history    historyport.History
// }

// func NewPgsSrv(req *srvport.ServicesReq) srvport.BootService {
// 	return &pgsInteractiveSrv{
// 		clinet:             openai.NewClient("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"),
// 		scriptCrashedTimes: 0,

// 		chat:       req.Chat,
// 		interactor: req.Interactor,
// 		history:    req.History,
// 	}
// }

// func (b *pgsInteractiveSrv) initLLMRole(prePrompt string) (*models.OSReply, error) {
// 	prompt := &models.Prompt{
// 		ClientPrompt: &prePrompt,
// 	}
// 	b.history.SavePrompt(prompt)
// 	reply, err := b.chat.Prompt(prompt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	b.interactor.Reply(reply)

// 	return reply, nil
// }

// func (b *pgsInteractiveSrv) runScript(reply *models.OSReply) (*models.ScriptResult, error) {
// 	b.interactor.Script(reply.ScriptRequest)
// 	defer func() { b.scriptCrashedTimes = 0 }()

// 	result, err := executeScript(reply.ScriptRequest)
// 	if err != nil {
// 		b.scriptCrashedTimes++
// 		crashPrompt := "last executed command crashed. recovering... try again"
// 		reply, err = b.chat.Prompt(&models.Prompt{
// 			ClientPrompt: &crashPrompt,
// 			UserPrompt:   nil,
// 			LastScriptResult: &models.ScriptResult{
// 				RunnerOSResult: &models.RunnerOSResponse{
// 					Stdout:     result.Stdout,
// 					StatusCode: result.StatusCode,
// 				},
// 			},
// 		})
// 		if err != nil {
// 			if b.scriptCrashedTimes > 5 {
// 				return nil, err
// 			}
// 			if reply != nil && reply.ScriptRequest != nil {
// 				return b.runScript(reply)
// 			}
// 			return nil, err
// 		}
// 	}

// 	scriptResults := &models.ScriptResult{
// 		RunnerOSResult: &models.RunnerOSResponse{
// 			Stdout:     result.Stdout,
// 			StatusCode: result.StatusCode,
// 		},
// 	}
// 	b.interactor.ScriptResults(scriptResults)

// 	return scriptResults, nil
// }

// func (b *pgsInteractiveSrv) Start(prePrompt string) (err error) {
// 	defer func() { fmt.Println("Start Defer, err:", err) }()

// 	reply, err := b.initLLMRole(prePrompt)
// 	if err != nil {
// 		return err
// 	}

// 	for {
// 		b.history.SaveResponse(reply)

// 		prompt := &models.Prompt{}

// 		if reply.MessageToUser != "" {
// 			b.interactor.Message(reply.MessageToUser, reply.ToeknsUsed)
// 		}

// 		if reply.ScriptRequest != nil {
// 			fmt.Println("response.ScriptRequest", reply.ScriptRequest)
// 			scriptResult, err := b.runScript(reply)
// 			if err != nil {
// 				return err
// 			}
// 			prompt.LastScriptResult = scriptResult
// 		}
// 		if reply.WaitForUserPrompt {
// 			userPrompt, err := b.interactor.GetUserInput()
// 			prompt.UserPrompt = &userPrompt
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		b.history.SavePrompt(prompt)
// 		reply, err = b.chat.Prompt(prompt)
// 		if err != nil {
// 			return err
// 		}
// 		b.interactor.Reply(reply)
// 	}
// }

// func executeScript(req *models.RunnerOSRequest) (*models.RunnerOSResponse, error) {
// 	cmd := exec.Command(req.Runtime)
// 	cmd.Stdin = strings.NewReader(req.Script)

// 	// output = stdout + stderr
// 	rawOutput, err := cmd.CombinedOutput()
// 	output := string(rawOutput)
// 	if err != nil {
// 		var exitErr *exec.ExitError
// 		if errors.As(err, &exitErr) {
// 			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
// 				return &models.RunnerOSResponse{
// 					StatusCode: status.ExitStatus(),
// 					Stdout:     output,
// 				}, err
// 			}
// 		}
// 		return &models.RunnerOSResponse{
// 			StatusCode: 0,
// 			Stdout:     output,
// 		}, err
// 	}

// 	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
// 		return &models.RunnerOSResponse{
// 			StatusCode: status.ExitStatus(),
// 			Stdout:     output,
// 		}, nil
// 	}

// 	return &models.RunnerOSResponse{
// 		StatusCode: 0,
// 		Stdout:     output,
// 	}, nil
// }
