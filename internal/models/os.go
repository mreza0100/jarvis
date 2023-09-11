package models

type OSReply struct {
	MessageToUser     string           `json:"messageToUser"`
	WaitForUserPrompt bool             `json:"waitForUserPrompt"`
	ScriptRequest     *OSRunnerRequest `json:"Script"`
	TokensUsed        int              `json:"_"`
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
	StatusCode int    `json:"StatusCode"`
}

type OSScriptResult struct {
	RunnerOSResult *OSRunnerResult `json:"RunnerOSResult"`
}
