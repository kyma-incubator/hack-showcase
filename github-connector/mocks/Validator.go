// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// ParseWebHook provides a mock function with given fields:
func (_m *Validator) ParseWebHook() {
	_m.Called()
}

// ValidatePayload provides a mock function with given fields: _a0, _a1
func (_m *Validator) ValidatePayload(_a0 *http.Request, _a1 []byte) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*http.Request, []byte) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, []byte) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
