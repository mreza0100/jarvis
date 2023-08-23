package models

type Prompt struct {
	ClientPrompt     *string       `json:"ClientPrompt"`
	UserPrompt       *string       `json:"UserPrompt"`
	Screen           Screen        `json:"Screen"`
	LastScriptResult *ScriptResult `json:"LastScriptResult"`
}

type Screen struct {
	Width  int `json:"Width"`
	Height int `json:"Height"`
}

type ScriptResult struct {
	Stdout         string `json:"Stdout"`
	TerminalStatus uint8  `json:"terminalStatus"`
}
