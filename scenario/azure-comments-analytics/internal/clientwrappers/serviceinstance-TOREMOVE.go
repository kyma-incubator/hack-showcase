package wrappers

// import (
// 	svcCatalog "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"
// 	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
// )

// type ServiceCatalogClient interface {
// 	Instance(namespace string) k8scomponents.ServiceInstance
// }

// type ServiceCatalogClientInterface interface {
// 	ServiceInstances(string) svcCatalog.ServiceInstanceInterface
// }

// type serviceCatalogClient struct {
// 	client ServiceCatalogClientInterface
// }

// func NewServiceCatalogClient(client ServiceCatalogClientInterface) ServiceCatalogClient {
// 	return serviceCatalogClient{client: client}
// }

// func (s serviceCatalogClient) Instance(namespace string) k8scomponents.ServiceInstance {
// 	return k8scomponents.NewServiceInstance(s.client.ServiceInstances(namespace), namespace)
// }
