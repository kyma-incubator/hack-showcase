package k8scomponents

import (
	"github.com/pkg/errors"
	"github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1beta1svc "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Binding describe binding struct
type Binding interface {
	Create(body *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	Prepare(name string, lambdaName string) *v1beta1.ServiceBinding
}

//BindingInterface describe constructors argument and containe ServiceBindings method
type BindingInterface interface {
	Create(*v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
}

type binding struct {
	bindingInterface BindingInterface
	namespace        string
}

//NewBinding create and return new binding struct
func NewBinding(client BindingInterface, nspace string) Binding {
	return &binding{bindingInterface: client, namespace: nspace}
}

func (s *binding) Create(body *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error) {
	data, err := s.bindingInterface.Create(body)
	if err != nil {
		return nil, errors.Wrap(err, "Can not create ServiceBinding")
	}
	return data, nil
}

func (s *binding) Delete(name string, options *v1.DeleteOptions) error {
	return s.bindingInterface.Delete(name, options)
}

func (s *binding) Prepare(name string, lambdaName string) *v1beta1.ServiceBinding {
	return &v1beta1.ServiceBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      name + "bind",
			Namespace: s.namespace,
			Labels: map[string]string{
				"Function": lambdaName,
			},
		},
		Spec: v1beta1svc.ServiceBindingSpec{
			InstanceRef: v1beta1svc.LocalObjectReference{
				Name: name + "inst",
			},
		},
	}
}
