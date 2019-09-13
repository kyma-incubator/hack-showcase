package wrappers

import (
	svcCatalog "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
)

type ServiceCatalogWrapper interface {
	Binding(namespace string) k8scomponents.Binding
	Instance(namespace string) k8scomponents.ServiceInstance
}

type ServiceCatalogClient interface {
	ServiceBindings(string) svcCatalog.ServiceBindingInterface
	ServiceInstances(string) svcCatalog.ServiceInstanceInterface
}

type serviceCatalogWrapper struct {
	client ServiceCatalogClient
}

func NewServiceCatalogClient(client ServiceCatalogClient) ServiceCatalogWrapper {
	return &serviceCatalogWrapper{client: client}
}

func (s *serviceCatalogWrapper) Binding(namespace string) k8scomponents.Binding {
	return k8scomponents.NewBinding(s.client.ServiceBindings(namespace), namespace)
}

func (s *serviceCatalogWrapper) Instance(namespace string) k8scomponents.ServiceInstance {
	return k8scomponents.NewServiceInstance(s.client.ServiceInstances(namespace), namespace)
}
