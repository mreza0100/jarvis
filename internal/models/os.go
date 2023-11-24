package models

type OSReply struct {
	ReplyToUser       string           `json:"ReplyToUser"`
	WaitForUserPrompt bool             `json:"WaitForUserPrompt"`
	ScriptRequest     *OSRunnerRequest `json:"Script"`
}

type OSRunnerRequest struct {
	Runtime string `json:"Runtime"`
	Script  string `json:"Script"`
}

type OSPrompt struct {
	ClientPrompt     *string         `json:"ClientPrompt"`
	UserPrompt       *string         `json:"UserPrompt"`
	Screen           *Screen         `json:"Screen"`
	LastScriptResult *OSScriptResult `json:"LastScriptResult"`
}

type OSRunnerResult struct {
	Stdout     string `json:"Stdout"`
	Stderr     string `json:"Stderr"`
	StatusCode int    `json:"StatusCode"`
}

type OSScriptResult struct {
	RunnerOSResult *OSRunnerResult `json:"RunnerOSResult"`
}

type OSConfig struct {
	Config *ChatConfig `json:"config"`
}
