package models

type OSReply struct {
	MessageToUser     string           `json:"messageToUser"`
	WaitForUserPrompt bool             `json:"waitForUserPrompt"`
	ScriptRequest     *RunnerOSRequest `json:"Script"`
	TokensUsed        int              `json:"_"`
}

type RunnerOSRequest struct {
	Runtime string `json:"Runtime"`
	Script  string `json:"Script"`
}

type RunnerOSResponse struct {
	Stdout     string `json:"Stdout"`
	StatusCode int    `json:"StatusCode"`
}

type Prompt struct {
	ClientPrompt     *string       `json:"ClientPrompt"`
	UserPrompt       *string       `json:"UserPrompt"`
	Screen           *Screen       `json:"Screen"`
	LastScriptResult *ScriptResult `json:"LastScriptResult"`
}

type Screen struct {
	Width  int `json:"Width"`
	Height int `json:"Height"`
}

type ScriptResult struct {
	RunnerOSResult *RunnerOSResponse `json:"RunnerOSResult"`
}
