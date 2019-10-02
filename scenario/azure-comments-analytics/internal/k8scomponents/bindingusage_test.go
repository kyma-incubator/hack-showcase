package k8scomponents_test

import (
	"errors"
	"testing"

	"github.com/kyma-incubator/github-slack-connectors/scenario/azure-comments-analytics/internal/k8scomponents"
	"github.com/kyma-incubator/github-slack-connectors/scenario/azure-comments-analytics/internal/k8scomponents/mocks"
	v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	v1alpha1svc "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateBindingUsage(t *testing.T) {
	t.Run("should create Binding, return new bindingUsage and nil", func(t *testing.T) {
		//given
		bindingUsage := &v1alpha1.ServiceBindingUsage{}
		mockClient := &mocks.BindingUsageInterface{}
		mockClient.On("Create", bindingUsage).Return(bindingUsage, nil)

		//when
		bind, err := k8scomponents.NewBindingUsage(mockClient, "default").Create(bindingUsage)

		//then
		assert.NoError(t, err)
		assert.Equal(t, bindingUsage, bind)
	})

	t.Run("should return nil and error when cannot create BindingUsage", func(t *testing.T) {
		//given
		bindingUsage := &v1alpha1.ServiceBindingUsage{}
		mockClient := &mocks.BindingUsageInterface{}
		mockClient.On("Create", bindingUsage).Return(nil, errors.New("error text"))

		//when
		bind, err := k8scomponents.NewBindingUsage(mockClient, "default").Create(bindingUsage)

		//then
		assert.Error(t, err)
		assert.Nil(t, bind)
	})
}

func TestDeleteBindingUsage(t *testing.T) {
	t.Run("should return ServiceBinding", func(t *testing.T) {
		//given
		name := "name"
		namespace := "namespace"
		options := &v1.DeleteOptions{}
		mockClient := &mocks.BindingUsageInterface{}
		mockClient.On("Delete", name, options).Return(nil)
		//when
		err := k8scomponents.NewBindingUsage(mockClient, namespace).Delete(name, options)

		//then
		assert.NoError(t, err)

	})
}

func TestGetEventBodyBindingUsage(t *testing.T) {
	t.Run("should return ServiceBindingUsage", func(t *testing.T) {
		//given
		name := "github-repo"
		namespace := "namespace"
		envPrefix := "prefix"
		lambdaName := "lambdaName"
		body := &v1alpha1.ServiceBindingUsage{
			TypeMeta: v1.TypeMeta{
				Kind:       "ServiceBindingUsage",
				APIVersion: "servicecatalog.kyma-project.io/v1alpha1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      name + "bu",
				Namespace: namespace,
				Labels: map[string]string{
					"Function":       lambdaName,
					"ServiceBinding": name + "bind",
				},
			},
			Spec: v1alpha1svc.ServiceBindingUsageSpec{
				ServiceBindingRef: v1alpha1svc.LocalReferenceByName{
					Name: name + "bind",
				},
				UsedBy: v1alpha1svc.LocalReferenceByKindAndName{
					Name: lambdaName,
					Kind: "function",
				},
				Parameters: &v1alpha1svc.Parameters{
					EnvPrefix: &v1alpha1svc.EnvPrefix{
						Name: envPrefix,
					},
				},
			},
		}
		mockClient := &mocks.BindingUsageInterface{}

		//when
		binding := k8scomponents.NewBindingUsage(mockClient, namespace).Prepare(name, envPrefix, lambdaName)

		//then
		assert.Equal(t, body, binding)

	})
}
