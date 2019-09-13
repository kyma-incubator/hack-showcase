package wrappers

import (
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	svcBind "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"
)

type KymaServiceCatalogWrapper interface {
	BindingUsage(namespace string) k8scomponents.BindingUsage
}

type KymaServiceCatalogClient interface {
	ServiceBindingUsages(string) svcBind.ServiceBindingUsageInterface
}

type kymaServiceCatalogWrapper struct {
	client KymaServiceCatalogClient
}

func NewKymaServiceCatalogWrapper(client KymaServiceCatalogClient) KymaServiceCatalogWrapper {
	return &kymaServiceCatalogWrapper{client: client}
}

func (s *kymaServiceCatalogWrapper) BindingUsage(namespace string) k8scomponents.BindingUsage {
	return k8scomponents.NewBindingUsage(s.client.ServiceBindingUsages(namespace), namespace)
}
