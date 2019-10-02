package k8scomponents_test

import (
	"errors"
	"testing"

	"github.com/kyma-incubator/github-slack-connectors/scenario/azure-comments-analytics/internal/k8scomponents"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/kyma-incubator/github-slack-connectors/scenario/azure-comments-analytics/internal/k8scomponents/mocks"
	v1beta1svc "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

func TestCreateServiceInstance(t *testing.T) {
	t.Run("should create Binding, return new binding and nil", func(t *testing.T) {
		//given
		instance := &v1beta1svc.ServiceInstance{}
		mockClient := &mocks.ServiceInstanceInterface{}
		mockClient.On("Create", instance).Return(instance, nil)

		//when
		data, err := k8scomponents.NewServiceInstance(mockClient, "default").Create(instance)

		//then
		assert.NoError(t, err)
		assert.Equal(t, instance, data)
	})

	t.Run("should return nil and error when cannot create Binding", func(t *testing.T) {
		//given
		instance := &v1beta1svc.ServiceInstance{}
		mockClient := &mocks.ServiceInstanceInterface{}
		mockClient.On("Create", instance).Return(nil, errors.New("error text"))

		//when
		data, err := k8scomponents.NewServiceInstance(mockClient, "default").Create(instance)

		//then
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestDeleteServiceInstance(t *testing.T) {
	t.Run("should return ServiceBinding", func(t *testing.T) {
		//given
		name := "name"
		namespace := "namespace"
		options := &v1.DeleteOptions{}
		mockClient := &mocks.ServiceInstanceInterface{}
		mockClient.On("Delete", name, options).Return(nil)
		//when
		err := k8scomponents.NewServiceInstance(mockClient, namespace).Delete(name, options)

		//then
		assert.NoError(t, err)

	})
}

func TestGetEventBodyServiceInstance(t *testing.T) {
	t.Run("should return ServiceBinding", func(t *testing.T) {
		//given
		name := "name"
		serviceClassExternalName := "serviceClassExternalName"
		plan := "plan"
		raw := runtime.RawExtension{}
		_ = raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
		namespace := "namespace"
		body := &v1beta1svc.ServiceInstance{
			ObjectMeta: v1.ObjectMeta{
				Name:      name + "inst",
				Namespace: namespace,
			},
			Spec: v1beta1svc.ServiceInstanceSpec{
				Parameters: &raw,
				PlanReference: v1beta1svc.PlanReference{
					ServiceClassExternalName: serviceClassExternalName,
					ServicePlanExternalName:  plan,
				},
			},
		}
		mockClient := &mocks.ServiceInstanceInterface{}

		//when
		binding := k8scomponents.NewServiceInstance(mockClient, namespace).Prepare(name, serviceClassExternalName, plan, &raw)

		//then
		assert.Equal(t, body, binding)

	})
}
