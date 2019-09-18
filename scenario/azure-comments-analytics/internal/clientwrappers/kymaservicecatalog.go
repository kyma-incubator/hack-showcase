package wrappers

import (
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	svcBind "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"
)

//KymaServiceCatalogWrapper is a wrapper dedicated for kyma's servicecatalog ClientSet
type KymaServiceCatalogWrapper interface {
	BindingUsage(namespace string) k8scomponents.BindingUsage
}

//KymaServiceCatalogClient describe constructors argument
type KymaServiceCatalogClient interface {
	ServiceBindingUsages(string) svcBind.ServiceBindingUsageInterface
}

type kymaServiceCatalogWrapper struct {
	client KymaServiceCatalogClient
}

//NewKymaServiceCatalogClient create and return kymaServiceCatalogWrapper
func NewKymaServiceCatalogClient(client KymaServiceCatalogClient) KymaServiceCatalogWrapper {
	return &kymaServiceCatalogWrapper{client: client}
}

func (s *kymaServiceCatalogWrapper) BindingUsage(namespace string) k8scomponents.BindingUsage {
	return k8scomponents.NewBindingUsage(s.client.ServiceBindingUsages(namespace), namespace)
}
