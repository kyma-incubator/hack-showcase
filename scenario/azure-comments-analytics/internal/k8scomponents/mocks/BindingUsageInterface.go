// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"

// BindingUsageInterface is an autogenerated mock type for the BindingUsageInterface type
type BindingUsageInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *BindingUsageInterface) Create(_a0 *v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error) {
	ret := _m.Called(_a0)

	var r0 *v1alpha1.ServiceBindingUsage
	if rf, ok := ret.Get(0).(func(*v1alpha1.ServiceBindingUsage) *v1alpha1.ServiceBindingUsage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ServiceBindingUsage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.ServiceBindingUsage) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, options
func (_m *BindingUsageInterface) Delete(name string, options *v1.DeleteOptions) error {
	ret := _m.Called(name, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.DeleteOptions) error); ok {
		r0 = rf(name, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
