// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-incubator/github-slack-connectors/github-connector/internal/apperrors"

import json "encoding/json"
import mock "github.com/stretchr/testify/mock"

// Sender is an autogenerated mock type for the Sender type
type Sender struct {
	mock.Mock
}

// SendToKyma provides a mock function with given fields: eventType, eventTypeVersion, eventID, sourceID, data
func (_m *Sender) SendToKyma(eventType string, eventTypeVersion string, eventID string, sourceID string, data json.RawMessage) apperrors.AppError {
	ret := _m.Called(eventType, eventTypeVersion, eventID, sourceID, data)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, string, string, json.RawMessage) apperrors.AppError); ok {
		r0 = rf(eventType, eventTypeVersion, eventID, sourceID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
