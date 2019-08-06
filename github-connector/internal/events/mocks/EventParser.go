// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	json "encoding/json"

	apperrors "github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/events"

	mock "github.com/stretchr/testify/mock"
)

// EventParser is an autogenerated mock type for the EventParser type
type EventParser struct {
	mock.Mock
}

// GetEventRequestAsJSON provides a mock function with given fields: EeventRequestPayload
func (_m *EventParser) GetEventRequestAsJSON(EeventRequestPayload events.EventRequestPayload) ([]byte, apperrors.AppError) {
	ret := _m.Called(EeventRequestPayload)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(events.EventRequestPayload) []byte); ok {
		r0 = rf(EeventRequestPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(events.EventRequestPayload) apperrors.AppError); ok {
		r1 = rf(EeventRequestPayload)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetEventRequestPayload provides a mock function with given fields: eventType, eventTypeVersion, eventID, sourceID, data
func (_m *EventParser) GetEventRequestPayload(eventType string, eventTypeVersion string, eventID string, sourceID string, data json.RawMessage) (events.EventRequestPayload, apperrors.AppError) {
	ret := _m.Called(eventType, eventTypeVersion, eventID, sourceID, data)

	var r0 events.EventRequestPayload
	if rf, ok := ret.Get(0).(func(string, string, string, string, json.RawMessage) events.EventRequestPayload); ok {
		r0 = rf(eventType, eventTypeVersion, eventID, sourceID, data)
	} else {
		r0 = ret.Get(0).(events.EventRequestPayload)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string, string, string, json.RawMessage) apperrors.AppError); ok {
		r1 = rf(eventType, eventTypeVersion, eventID, sourceID, data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
