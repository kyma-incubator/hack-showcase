package k8scomponents

import (
	v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//BindingUsage describe bindingUsage struct
type BindingUsage interface {
	Create(body *v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error)
	Delete(name string, options *v1.DeleteOptions) error
	Prepare(name string, envPrefix string, lambdaName string) *v1alpha1.ServiceBindingUsage
}

//BindingUsageInterface describe constructors argument and containe ServiceBindingUsages method
type BindingUsageInterface interface {
	Create(*v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error)
	Delete(name string, options *v1.DeleteOptions) error
}

type bindingUsage struct {
	catalog   BindingUsageInterface
	namespace string
}

//NewBindingUsage create and return new bindingUsage
func NewBindingUsage(scatalog BindingUsageInterface, nspace string) BindingUsage {
	return &bindingUsage{catalog: scatalog, namespace: nspace}
}

func (s *bindingUsage) Create(body *v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error) {
	data, err := s.catalog.Create(body)
	if err != nil {
		return nil, errors.Wrap(err, "Can not create ServiceBindingUsage")
	}
	return data, nil
}

func (s *bindingUsage) Delete(name string, options *v1.DeleteOptions) error {
	return s.catalog.Delete(name, options)
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
		Spec: v1alpha1.ServiceBindingUsageSpec{
			ServiceBindingRef: v1alpha1.LocalReferenceByName{
				Name: name + "bind",
			},
			UsedBy: v1alpha1.LocalReferenceByKindAndName{
				Name: lambdaName,
				Kind: "function",
			},
			Parameters: &v1alpha1.Parameters{
				EnvPrefix: &v1alpha1.EnvPrefix{
					Name: envPrefix,
				},
			},
		},
	}
}
