package models

type Response struct {
	MessageToUser     string         `json:"messageToUser"`
	WaitForUserPrompt bool           `json:"waitForUserPrompt"`
	ScriptRequest     *ScriptRequest `json:"Script"`
}

type ScriptRequest struct {
	Runtime       string `json:"runtime"`
	Script        string `json:"script"`
	ReturnResults bool   `json:"ReturnResults"`
}
