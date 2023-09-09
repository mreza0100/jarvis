package models

import "encoding/json"

type HistoryRecord struct {
	Prompt interface{}
	Reply  interface{}
}

func (h *HistoryRecord) Marshal() (string, error) {
	enc, err := json.Marshal(h)
	return string(enc), err
}

func (h *HistoryRecord) Unmarshal(rawHistory string) error {
	return json.Unmarshal([]byte(rawHistory), h)
}
