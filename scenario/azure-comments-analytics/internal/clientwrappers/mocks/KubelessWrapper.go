// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import k8scomponents "github.com/kyma-incubator/github-slack-connectors/scenario/github-issue-sentiment-analysis/internal/k8scomponents"
import mock "github.com/stretchr/testify/mock"

// KubelessWrapper is an autogenerated mock type for the KubelessWrapper type
type KubelessWrapper struct {
	mock.Mock
}

// Function provides a mock function with given fields: namespace
func (_m *KubelessWrapper) Function(namespace string) k8scomponents.Function {
	ret := _m.Called(namespace)

	var r0 k8scomponents.Function
	if rf, ok := ret.Get(0).(func(string) k8scomponents.Function); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(k8scomponents.Function)
		}
	}

	return r0
}
