package k8scomponents

import (
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	v1beta1svc "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

//ServiceInstanceInterface describe constructors argument and include important methods of ServiceInstances
type ServiceInstanceInterface interface {
	Create(*v1beta1svc.ServiceInstance) (*v1beta1svc.ServiceInstance, error)
}

//ServiceInstance describe serviceInstance struct
type ServiceInstance interface {
	GetEventBody(name string, serviceClassExternalName string, plan string, parameters *runtime.RawExtension) *v1beta1svc.ServiceInstance
	Create(body *v1beta1svc.ServiceInstance) (*v1beta1svc.ServiceInstance, apperrors.AppError)
}

type serviceInstance struct {
	instance  ServiceInstanceInterface
	namespace string
}

//NewServiceInstance returns new serviceInstance struct
func NewServiceInstance(instance ServiceInstanceInterface, namespace string) ServiceInstance {
	return serviceInstance{instance: instance, namespace: namespace}
}

func (s serviceInstance) Create(body *v1beta1svc.ServiceInstance) (*v1beta1svc.ServiceInstance, apperrors.AppError) {
	data, err := s.instance.Create(body)
	if err != nil {
		return nil, apperrors.WrongInput("Can not create Function: %s", err)
	}
	return data, nil
}

func (s serviceInstance) GetEventBody(name string, serviceClassExternalName string, plan string, parameters *runtime.RawExtension) *v1beta1svc.ServiceInstance {
	return &v1beta1svc.ServiceInstance{
		ObjectMeta: v1.ObjectMeta{
			Name:      name + "inst",
			Namespace: s.namespace,
		},
		Spec: v1beta1svc.ServiceInstanceSpec{
			Parameters: parameters,
			PlanReference: v1beta1svc.PlanReference{
				ServiceClassExternalName: serviceClassExternalName,
				ServicePlanExternalName:  plan,
			},
		},
	}
}
