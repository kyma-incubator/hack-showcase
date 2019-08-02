package eventparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// EventParser adds proper structure to existing payload, so it can be consumed by Event-Service
type EventParser interface {
	GetEventRequestPayload(string) (EventRequestPayload, error)
	GetEventRequestAsJSON(EventRequestPayload) []byte
}

type eventparser struct {
}

// EventRequestPayload represents a POST request's body which is sent to Event-Service
type EventRequestPayload struct {
	EventType        string          `json:"event-type"`
	EventTypeVersion string          `json:"event-type-version"`
	EventID          string          `json:"event-id,omitempty"` //uuid should be generated automatically if send empty
	EventTime        string          `json:"event-time"`
	SourceID         string          `json:"source-id"` //put your application name here
	Data             json.RawMessage `json:"data"`      //github webhook json payload
}

// GetEventRequestPayload generates structure which is mapped to JSON required by Event-Service request body
func (e eventparser) GetEventRequestPayload(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) (EventRequestPayload, error) {

	if eventType == "" {
		return EventRequestPayload{}, errors.New("eventType should not be empty")
	}
	if eventTypeVersion == "" {
		return EventRequestPayload{}, errors.New("eventTypeVersion should not be empty")
	}
	if sourceID == "" {
		return EventRequestPayload{}, errors.New("sourceID should not be empty")
	}
	if len(data) == 0 {
		return EventRequestPayload{}, errors.New("data should not be empty")
	}

	res := EventRequestPayload{
		eventType,
		eventTypeVersion,
		eventID,
		time.Now().Format(time.RFC3339),
		sourceID,
		data}

	return res, nil

}

// GetEventRequestAsJSON returns ready-to-sent JSON request body
func (e eventparser) GetEventRequestAsJSON(eventRequestPayload EventRequestPayload) ([]byte, error) {
	r, err := json.MarshalIndent(eventRequestPayload, "", "  ")
	if err != nil {
		fmt.Printf("error: %s", err)

		return []byte{}, err
	}
	return r, nil
}
