// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// Create provides a mock function with given fields: body
func (_m *Subscription) Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, apperrors.AppError) {
	ret := _m.Called(body)

	var r0 *v1alpha1.Subscription
	if rf, ok := ret.Get(0).(func(*v1alpha1.Subscription) *v1alpha1.Subscription); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Subscription)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(*v1alpha1.Subscription) apperrors.AppError); ok {
		r1 = rf(body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Prepare provides a mock function with given fields: id, lambdaName
func (_m *Subscription) Prepare(id string, lambdaName string) *v1alpha1.Subscription {
	ret := _m.Called(id, lambdaName)

	var r0 *v1alpha1.Subscription
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.Subscription); ok {
		r0 = rf(id, lambdaName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Subscription)
		}
	}

	return r0
}
