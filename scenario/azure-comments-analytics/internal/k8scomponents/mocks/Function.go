// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"

import mock "github.com/stretchr/testify/mock"
import v1beta1 "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"

// Function is an autogenerated mock type for the Function type
type Function struct {
	mock.Mock
}

// Create provides a mock function with given fields: body
func (_m *Function) Create(body *v1beta1.Function) (*v1beta1.Function, apperrors.AppError) {
	ret := _m.Called(body)

	var r0 *v1beta1.Function
	if rf, ok := ret.Get(0).(func(*v1beta1.Function) *v1beta1.Function); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Function)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(*v1beta1.Function) apperrors.AppError); ok {
		r1 = rf(body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetEventBody provides a mock function with given fields:
func (_m *Function) GetEventBody() *v1beta1.Function {
	ret := _m.Called()

	var r0 *v1beta1.Function
	if rf, ok := ret.Get(0).(func() *v1beta1.Function); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Function)
		}
	}

	return r0
}
