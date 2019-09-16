package manager_test

import (
	"testing"

	function "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	componentsMocks "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents/mocks"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/manager"
	subscriptions "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	serviceBindingUsages "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	bindings "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	serviceInstance "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/stretchr/testify/assert"
)

func TestCreateSubscription(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.Subscription{}
		subscriptionBody := &subscriptions.Subscription{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, nil)
		component.On("GetEventBody", "githubRepo").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateSubscription(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.Subscription{}
		subscriptionBody := &subscriptions.Subscription{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, apperrors.Internal("error"))
		component.On("GetEventBody", "githubRepo").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateSubscription(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateServiceBindingUsages(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.BindingUsage{}
		bindingUsageBody := &serviceBindingUsages.ServiceBindingUsage{}
		component.On("Create", bindingUsageBody).Return(bindingUsageBody, nil)
		component.On("GetEventBody", "githubRepo", "GITHUB_").Return(bindingUsageBody)
		component.On("GetEventBody", "slackWorkspace", "", "githubRepo").Return(bindingUsageBody)
		component.On("GetEventBody", "azureServiceName", "", "githubRepo").Return(bindingUsageBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindingUsages(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.BindingUsage{}
		bindingUsageBody := &serviceBindingUsages.ServiceBindingUsage{}
		component.On("Create", bindingUsageBody).Return(bindingUsageBody, apperrors.Internal("error"))
		component.On("GetEventBody", "githubRepo", "GITHUB_").Return(bindingUsageBody)
		component.On("GetEventBody", "slackWorkspace", "", "githubRepo").Return(bindingUsageBody)
		component.On("GetEventBody", "azureServiceName", "", "githubRepo").Return(bindingUsageBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindingUsages(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateServiceBindings(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.Binding{}
		bindingBody := &bindings.ServiceBinding{}
		component.On("Create", bindingBody).Return(bindingBody, nil)
		component.On("GetEventBody", "githubRepo").Return(bindingBody)
		component.On("GetEventBody", "slackWorkspace", "githubRepo").Return(bindingBody)
		component.On("GetEventBody", "azureServiceName", "githubRepo").Return(bindingBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindings(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.Binding{}
		bindingBody := &bindings.ServiceBinding{}
		component.On("Create", bindingBody).Return(bindingBody, apperrors.Internal("error"))
		component.On("GetEventBody", "githubRepo").Return(bindingBody)
		component.On("GetEventBody", "slackWorkspace", "githubRepo").Return(bindingBody)
		component.On("GetEventBody", "azureServiceName", "githubRepo").Return(bindingBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindings(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateFunction(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.Function{}
		subscriptionBody := &function.Function{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, nil)
		component.On("GetEventBody", "githubRepo").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateFunction(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.Function{}
		subscriptionBody := &function.Function{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, apperrors.Internal("error"))
		component.On("GetEventBody", "githubRepo").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateFunction(component)

		//then
		assert.Error(t, err)
	})
}

func RawExtensionNiller() *runtime.RawExtension {
	return nil
}

func TestCreateServiceInstances(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.ServiceInstance{}
		serviceInstanceBody := &serviceInstance.ServiceInstance{}
		raw := runtime.RawExtension{}
		unmarshalerr := raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
		component.On("Create", serviceInstanceBody).Return(serviceInstanceBody, nil)
		component.On("GetEventBody", "azureServiceName", "azureServiceName", "standard-s0", &raw).Return(serviceInstanceBody)
		component.On("GetEventBody", "githubRepo", "githubRepo-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		component.On("GetEventBody", "slackWorkspace", "slackWorkspace-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		serviceInstanceList := serviceInstance.ServiceClassList{
			Items: []serviceInstance.ServiceClass{
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "githubRepo-12345",
						},
					},
				}, serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "slackWorkspace-12345",
						},
					},
				},
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "azureServiceName",
						},
					},
				}},
		}
		//when
		err := testedManager.CreateServiceInstances(component, &serviceInstanceList)

		//then
		assert.NoError(t, err)
		assert.NoError(t, unmarshalerr)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.ServiceInstance{}
		serviceInstanceBody := &serviceInstance.ServiceInstance{}
		raw := runtime.RawExtension{}
		unmarshalerr := raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
		component.On("Create", serviceInstanceBody).Return(serviceInstanceBody, apperrors.Internal("error"))
		component.On("GetEventBody", "azureServiceName", "azureServiceName", "standard-s0", &raw).Return(serviceInstanceBody)
		component.On("GetEventBody", "githubRepo", "githubRepo-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		component.On("GetEventBody", "slackWorkspace", "slackWorkspace-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		serviceInstanceList := serviceInstance.ServiceClassList{
			Items: []serviceInstance.ServiceClass{
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "githubRepo-12345",
						},
					},
				}, serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "slackWorkspace-12345",
						},
					},
				},
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "azureServiceName",
						},
					},
				}},
		}
		//when
		err := testedManager.CreateServiceInstances(component, &serviceInstanceList)

		//then
		assert.Error(t, err)
		assert.NoError(t, unmarshalerr)
	})
}
