package k8scomponents

import (
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	"github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1beta1svc "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Binding describe binding struct
type Binding interface {
	Create(body *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, apperrors.AppError)
	GetEventBody(name string) *v1beta1.ServiceBinding
}

//BindingInterface describe constructors argument and containe ServiceBindings method
type BindingInterface interface {
	Create(*v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
}

type binding struct {
	bindingInterface BindingInterface
	namespace        string
}

//NewBinding create and return new binding struct
func NewBinding(client BindingInterface, nspace string) Binding {
	return &binding{bindingInterface: client, namespace: nspace}
}

func (s *binding) Create(body *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, apperrors.AppError) {
	data, err := s.bindingInterface.Create(body)
	if err != nil {
		return nil, apperrors.Internal("Can not create ServiceBinding: %s", err)
	}
	return data, nil
}

func (s *binding) GetEventBody(name string) *v1beta1.ServiceBinding {
	return &v1beta1.ServiceBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      name + "bind",
			Namespace: s.namespace,
			Labels: map[string]string{
				"Function": "julia-lambda",
			},
		},
		Spec: v1beta1svc.ServiceBindingSpec{
			InstanceRef: v1beta1svc.LocalObjectReference{
				Name: name + "inst",
			},
		},
	}
}
