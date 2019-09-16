package k8scomponents_test

import (
	"errors"
	"testing"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents/mocks"
	"github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1beta1svc "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

func bindingNiller() *v1beta1.ServiceBinding {
	return nil
}

func TestCreateBinding(t *testing.T) {
	t.Run("should create Binding, return new binding and nil", func(t *testing.T) {
		//given
		binding := &v1beta1.ServiceBinding{}
		mockClient := &mocks.BindingInterface{}
		mockClient.On("Create", binding).Return(binding, nil)

		//when
		bind, err := k8scomponents.NewBinding(mockClient, "default").Create(binding)

		//then
		assert.NoError(t, err)
		assert.Equal(t, binding, bind)
	})

	t.Run("should return nil and error when cannot create Binding", func(t *testing.T) {
		//given
		binding := &v1beta1.ServiceBinding{}
		mockClient := &mocks.BindingInterface{}
		mockClient.On("Create", binding).Return(nil, errors.New("error text"))

		//when
		bind, err := k8scomponents.NewBinding(mockClient, "default").Create(binding)

		//then
		assert.Error(t, err)
		assert.Equal(t, bindingNiller(), bind)
	})
}

func TestGetEventBodyBinding(t *testing.T) {
	t.Run("should return ServiceBinding", func(t *testing.T) {
		//given
		name := "github-repo"
		namespace := "namespace"
		body := &v1beta1.ServiceBinding{
			ObjectMeta: v1.ObjectMeta{
				Name:      name + "bind",
				Namespace: namespace,
				Labels: map[string]string{
					"Function": name[7:] + "-lambda",
				},
			},
			Spec: v1beta1svc.ServiceBindingSpec{
				InstanceRef: v1beta1svc.LocalObjectReference{
					Name: name + "inst",
				},
			},
		}
		mockClient := &mocks.BindingInterface{}

		//when
		binding := k8scomponents.NewBinding(mockClient, namespace).GetEventBody(name)

		//then
		assert.Equal(t, body, binding)

	})
}
