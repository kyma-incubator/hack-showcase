package manager

import (
	"log"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

//Manager include important methods to deploy all k8s and kymas components to realize hack-showcase scenario
type Manager interface {
	CreateFunction(function k8scomponents.Function) apperrors.AppError
	CreateServiceBindings(binding k8scomponents.Binding) apperrors.AppError
	CreateSubscription(subscription k8scomponents.Subscription) apperrors.AppError
	CreateServiceBindingUsages(bindingUsage k8scomponents.BindingUsage) apperrors.AppError
	CreateServiceInstances(instance k8scomponents.ServiceInstance, serviceClassList *v1beta1.ServiceClassList) apperrors.AppError
}
type manager struct {
	githubRepo       string
	slackWorkspace   string
	azureServiceName string
	namespace        string
}

//NewManager create and return new manager struct
func NewManager(namespace string, githubRepo string, slackWorkspace string, azureServiceName string) Manager {
	return &manager{namespace: namespace, githubRepo: githubRepo, slackWorkspace: slackWorkspace, azureServiceName: azureServiceName}
}

func (s *manager) CreateSubscription(subscription k8scomponents.Subscription) apperrors.AppError {
	subscribe, err := subscription.Create(subscription.GetEventBody(s.githubRepo))
	if err != nil {
		return err
	}
	log.Printf("Subscription: %s", subscribe.Name)
	return nil
}

func (s *manager) CreateServiceBindingUsages(bindingUsage k8scomponents.BindingUsage) apperrors.AppError {
	usage1, err := bindingUsage.Create(bindingUsage.GetEventBody(s.githubRepo, "GITHUB_"))
	if err != nil {
		return err
	}
	log.Printf("SvcBindingUsage-1: %s\n", usage1.Name)

	usage2, err := bindingUsage.Create(bindingUsage.GetEventBody(s.slackWorkspace, ""))
	if err != nil {
		return err
	}
	log.Printf("SvcBindingUsage-2: %s\n", usage2.Name)

	usage3, err := bindingUsage.Create(bindingUsage.GetEventBody(s.azureServiceName, ""))
	if err != nil {
		return err
	}
	log.Printf("SvcBindingUsage-3: %s\n", usage3.Name)
	return nil
}

func (s *manager) CreateServiceBindings(binding k8scomponents.Binding) apperrors.AppError {
	bind1, err := binding.Create(binding.GetEventBody(s.githubRepo))
	log.Printf("SvcBinding-1: %s\n", bind1.Name)
	bind2, err := binding.Create(binding.GetEventBody(s.slackWorkspace))
	if err != nil {
		return err
	}
	log.Printf("SvcBinding-2: %s\n", bind2.Name)
	bind3, err := binding.Create(binding.GetEventBody(s.azureServiceName))
	if err != nil {
		return err
	}
	log.Printf("SvcBinding-3: %s\n", bind3.Name)
	return nil
}

func (s *manager) CreateFunction(function k8scomponents.Function) apperrors.AppError {
	funct, err := function.Create(function.GetEventBody())
	if err != nil {
		return err
	}
	log.Printf("Function: %s\n", funct.Name)
	return nil
}

func (s *manager) CreateServiceInstances(instance k8scomponents.ServiceInstance, serviceClassList *v1beta1.ServiceClassList) apperrors.AppError {
	for _, serv := range serviceClassList.Items {
		chars := []rune(serv.Spec.ExternalName)
		str := string(chars[0 : len(chars)-6])
		if str == s.githubRepo {
			svc, err := instance.Create(instance.GetEventBody(s.githubRepo, serv.Spec.ExternalName, "default", nil))
			if err != nil {
				return err
			}
			log.Printf("ServiceInstance-1: %s", svc.Name)
		}
		if str == s.slackWorkspace {
			svc, err := instance.Create(instance.GetEventBody(s.slackWorkspace, serv.Spec.ExternalName, "default", nil))
			if err != nil {
				return err
			}
			log.Printf("ServiceInstance-2: %s", svc.Name)
		}
		if string(chars) == s.azureServiceName {
			raw := runtime.RawExtension{}
			err := raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
			if err != nil {
				return apperrors.Internal("%s", err)
			}
			svc, err := instance.Create(instance.GetEventBody(s.azureServiceName, serv.Spec.ExternalName, "standard-s0", &raw))
			if err != nil {
				return apperrors.Internal("%s", err)
			}
			log.Printf("ServiceInstance-3: %s", svc.Name)
		}
	}
	return nil
}
