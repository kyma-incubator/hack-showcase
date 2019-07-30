package eventparser

import (
	"encoding/json"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

// EventParser adds proper structure to existing payload, so it can be consumed by Event-Service
type EventParser interface {
	GetEventRequestPayload(string) (EventRequestPayload, error)
	GetEventRequestAsJSON(EventRequestPayload) []byte
}

// EventRequestPayload represents a POST request's body which is sent to Event-Service
type EventRequestPayload struct {
	EventType        string          `json:"event-type"`
	EventTypeVersion string          `json:"event-type-version"`
	EventID          string          `json:"event-id,omitempty"` //uuid should be generated automatically if send empty
	EventTime        string          `json:"event-time"`
	Data             json.RawMessage `json:"data"` //github webhook json payload
}

// GetEventRequestPayload generates structure which is mapped to JSON required by Event-Service request body
func GetEventRequestPayload(eventType, eventTypeVersion, eventID string, data json.RawMessage) (EventRequestPayload, apperrors.AppError) {

	if eventType == "" {
		return EventRequestPayload{}, apperrors.WrongInput("eventType should not be empty")
	}
	if eventTypeVersion == "" {
		return EventRequestPayload{}, apperrors.WrongInput("eventTypeVersion should not be empty")
	}
	if len(data) == 0 {
		return EventRequestPayload{}, apperrors.WrongInput("data should not be empty")
	}

	res := EventRequestPayload{
		eventType,
		eventTypeVersion,
		eventID,
		time.Now().Format(time.RFC3339),
		data}

	return res, nil

}

// GetEventRequestAsJSON returns ready-to-sent JSON request body
func GetEventRequestAsJSON(eventRequestPayload EventRequestPayload) ([]byte, apperrors.AppError) {
	r, err := json.MarshalIndent(eventRequestPayload, "", "  ")

	if err != nil {

		return []byte{}, apperrors.Internal("Can not marshall given struct: %s", err.Error())
	}
	return r, nil
}
