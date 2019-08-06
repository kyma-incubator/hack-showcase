// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
import events "github.com/kyma-incubator/hack-showcase/github-connector/internal/events"
import mock "github.com/stretchr/testify/mock"

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: payload
func (_m *Validator) Validate(payload events.EventRequestPayload) apperrors.AppError {
	ret := _m.Called(payload)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(events.EventRequestPayload) apperrors.AppError); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
