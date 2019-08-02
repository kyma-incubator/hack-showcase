package eventparser

import (
	"encoding/json"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

// EventParser adds proper structure to existing payload, so it can be consumed by Event-Service

type eventparser struct {
	EventParser
}

type EventParser interface {
	GetEventRequestPayload(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) (EventRequestPayload, apperrors.AppError)
	GetEventRequestAsJSON(EeventRequestPayload EventRequestPayload) ([]byte, apperrors.AppError)
}

func NewEventParser() eventparser {
	return eventparser{}
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
func (e eventparser) GetEventRequestPayload(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) (EventRequestPayload, apperrors.AppError) {

	if eventType == "" {
		return EventRequestPayload{}, apperrors.WrongInput("eventType should not be empty")
	}
	if eventTypeVersion == "" {
		return EventRequestPayload{}, apperrors.WrongInput("eventTypeVersion should not be empty")
	}
	if sourceID == "" {
		return EventRequestPayload{}, apperrors.WrongInput("sourceID should not be empty")
	}
	if len(data) == 0 {
		return EventRequestPayload{}, apperrors.WrongInput("data should not be empty")
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
func (e eventparser) GetEventRequestAsJSON(eventRequestPayload EventRequestPayload) ([]byte, apperrors.AppError) {
	r, err := json.MarshalIndent(eventRequestPayload, "", "  ")

	if err != nil {

		return []byte{}, apperrors.Internal("Can not marshall given struct: %s", err.Error())
	}
	return r, nil
}
