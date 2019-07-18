package eventparser

import (
	"time"
)

// GitHubPayload mock
type GitHubPayload struct {
}

// EventRequestPayload represents a request which is sent to Event-Service"
type EventRequestPayload struct {
	EventType        string `json:"event-type"`
	EventTypeVersion string `json:"event-type-version"`
	EventID          string `json:"event-id"` //uuid should be generated automatically if send empty
	EventTime        string `json:"event-time"`
	Data             []byte `json:"data"`
}

// GetEventRequestPayload is getter
func GetEventRequestPayload(eventType, eventTypeVersion, eventID string, data []byte) EventRequestPayload {
	return EventRequestPayload{
		eventType,
		eventTypeVersion,
		eventID,
		time.Now().Format(time.RFC3339),
		data}

}
