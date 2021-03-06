// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"

// KymaServiceCatalogClient is an autogenerated mock type for the KymaServiceCatalogClient type
type KymaServiceCatalogClient struct {
	mock.Mock
}

// ServiceBindingUsages provides a mock function with given fields: _a0
func (_m *KymaServiceCatalogClient) ServiceBindingUsages(_a0 string) v1alpha1.ServiceBindingUsageInterface {
	ret := _m.Called(_a0)

	var r0 v1alpha1.ServiceBindingUsageInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ServiceBindingUsageInterface); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ServiceBindingUsageInterface)
		}
	}

	return r0
}
