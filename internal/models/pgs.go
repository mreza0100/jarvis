package models

type PgsReply struct {
	ReplyToUser       string            `json:"ReplyToUser"`
	WaitForUserPrompt bool              `json:"WaitForUserPrompt"`
	QueryRequest      *PgsRunnerRequest `json:"QueryRequest"`
	TokensUsed        int               `json:"_"`
}

type PgsRunnerRequest struct {
	Query string `json:"Query"`
}

type PgsPrompt struct {
	ClientPrompt    *string          `json:"ClientPrompt"`
	UserPrompt      *string          `json:"UserPrompt"`
	Screen          *Screen          `json:"Screen"`
	LastQueryResult *PgsScriptResult `json:"LastQueryResult"`
}

type PgsScriptResult struct {
	RunnerPgsResult *PgsRunnerResponse `json:"RunnerPgsResult"`
}

type QueryResult struct {
	ColumnValues []any         `json:"ColumnValues"`
	Columns      []string      `json:"Columns"`
	ColumnsType  []*ColumnType `json:"ColumnType"`
	Err          error         `json:"Err"`
}

type ColumnType struct {
	Name             string `json:"Name"`
	Length           *int64 `json:"Length"`
	Nullable         *bool  `json:"Nullable"`
	DatabaseTypeName string `json:"DatabaseTypeName"`
}

type PgsRunnerResponse struct {
	Err            error          `json:"Err"`
	QueryResponses []*QueryResult `json:"PgsRunnerResponse"`
}

type PostgresConnConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type PostgresConfig struct {
	Config *ChatConfig `json:"config"`

	PostgresConnConfig *PostgresConnConfig `json:"connection"`
}
