package rest

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/apex/log"
)

type LogContent struct {
	Timestamp int64  `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"content"`
	FleetId   string `json:"fleet.id"`
	TaxiId    string `json:"taxi.id"`
}

// Handler implementation
type Handler struct {
	mu     sync.Mutex
	Client DTClient
}

// New handler with DT Client
func New(dtc DTClient) *Handler {
	return &Handler{
		Client: dtc,
	}
}

// HandleLog implements log.Handler
func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	var logContent LogContent
	// build content
	if e.Fields.Get("taxi.id") != "" {
		logContent = LogContent{
			Timestamp: e.Timestamp.UTC().UnixMilli(),
			Level:     e.Level.String(),
			Message:   e.Message,
			TaxiId:    fmt.Sprintf("%v", e.Fields.Get("taxi.id")),
			FleetId:   fmt.Sprintf("%v", e.Fields.Get("fleet.id")),
		}
	} else {
		logContent = LogContent{
			Timestamp: e.Timestamp.UTC().UnixMilli(),
			Level:     e.Level.String(),
			Message:   e.Message,
			FleetId:   fmt.Sprintf("%v", e.Fields.Get("fleet.id")),
		}
	}
	content, _ := json.Marshal(logContent)

	// send to client
	h.Client.PostLogEvent(content)

	return nil
}
