package k8scomponents

import (
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	v1alpha1svc "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//BindingUsage describe bindingUsage struct
type BindingUsage interface {
	Create(body *v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, apperrors.AppError)
	Prepare(name string, envPrefix string, lambdaName string) *v1alpha1.ServiceBindingUsage
}

//BindingUsageInterface describe constructors argument and containe ServiceBindingUsages method
type BindingUsageInterface interface {
	Create(*v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error)
}

type bindingUsage struct {
	catalog   BindingUsageInterface
	namespace string
}

//NewBindingUsage create and return new bindingUsage
func NewBindingUsage(scatalog BindingUsageInterface, nspace string) BindingUsage {
	return &bindingUsage{catalog: scatalog, namespace: nspace}
}

func (s *bindingUsage) Create(body *v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, apperrors.AppError) {
	data, err := s.catalog.Create(body)
	if err != nil {
		return nil, apperrors.Internal("Can not create ServiceBindingUsage: %s", err)
	}
	return data, nil
}

func (s *bindingUsage) Prepare(name string, envPrefix string, lambdaName string) *v1alpha1.ServiceBindingUsage {
	return &v1alpha1.ServiceBindingUsage{
		TypeMeta: v1.TypeMeta{
			Kind:       "ServiceBindingUsage",
			APIVersion: "servicecatalog.kyma-project.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      name + "bu",
			Namespace: s.namespace,
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
}
