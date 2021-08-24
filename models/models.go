package models

import (
	"encoding/json"
	"time"
)

type PlainHeaders map[string]string

type ProxyRequest struct {
	Method  string       `json:"method" binding:"alpha"`
	URL     string       `json:"url" binding:"url"`
	Headers PlainHeaders `json:"headers"`
}

type ProxyResponse struct {
	ID      string       `json:"id,omitempty"`
	Status  int          `json:"status"`
	Length  int64        `json:"length"`
	Headers PlainHeaders `json:"headers"`
}

type HistoryEntry struct {
	AddedAt  time.Time     `json:"-"`
	Request  ProxyRequest  `json:"request"`
	Response ProxyResponse `json:"response"`
}

func (e *HistoryEntry) MarshalJSON() ([]byte, error) {
	type TimeSerializable HistoryEntry
	return json.Marshal(&struct {
		*TimeSerializable
		AddedAt string `json:"added_at"`
	}{
		TimeSerializable: (*TimeSerializable)(e),
		AddedAt:          e.AddedAt.Format(time.RFC1123),
	})
}
