package manager

import (
	"log"
	"strings"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

const azureConfiguration = `{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`

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
	lambdaName       string
}

//NewManager create and return new manager struct
func NewManager(namespace string, githubRepo string, slackWorkspace string, azureServiceName string) Manager {
	return &manager{
		namespace:        namespace,
		githubRepo:       githubRepo,
		slackWorkspace:   slackWorkspace,
		azureServiceName: azureServiceName,
		lambdaName:       githubRepo[7:] + "-lambda", //Due to Kyma's requirements lambda's name has to be short - it's trimmed here
	}
}

func (s *manager) CreateSubscription(subscription k8scomponents.Subscription) apperrors.AppError {
	subscribe, err := subscription.Create(subscription.Prepare(s.githubRepo, s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("Subscription: %s", subscribe.Name)
	return nil
}

func (s *manager) CreateServiceBindingUsages(bindingUsage k8scomponents.BindingUsage) apperrors.AppError {
	usage1, err := bindingUsage.Create(bindingUsage.Prepare(s.githubRepo, "GITHUB_", s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("SvcBindingUsage-1: %s\n", usage1.Name)

	usage2, err := bindingUsage.Create(bindingUsage.Prepare(s.slackWorkspace, "", s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("SvcBindingUsage-2: %s\n", usage2.Name)

	usage3, err := bindingUsage.Create(bindingUsage.Prepare(s.azureServiceName, "", s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("SvcBindingUsage-3: %s\n", usage3.Name)
	return nil
}

func (s *manager) CreateServiceBindings(binding k8scomponents.Binding) apperrors.AppError {
	bind1, err := binding.Create(binding.Prepare(s.githubRepo, s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("SvcBinding-1: %s\n", bind1.Name)
	bind2, err := binding.Create(binding.Prepare(s.slackWorkspace, s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("SvcBinding-2: %s\n", bind2.Name)
	bind3, err := binding.Create(binding.Prepare(s.azureServiceName, s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("SvcBinding-3: %s\n", bind3.Name)
	return nil
}

func (s *manager) CreateFunction(function k8scomponents.Function) apperrors.AppError {
	funct, err := function.Create(function.Prepare(s.githubRepo, s.lambdaName))
	if err != nil {
		return err
	}
	log.Printf("Function: %s\n", funct.Name)
	return nil
}

func (s *manager) CreateServiceInstances(instance k8scomponents.ServiceInstance, serviceClassList *v1beta1.ServiceClassList) apperrors.AppError {
	//ServiceClass ExternalName suffix is generated randomly, but its prefix is based on name provided by user.
	//Looking for ServiceClass with matching prefix on which basis ServiceInstance should be created.
	for _, serv := range serviceClassList.Items {
		if strings.HasPrefix(serv.Spec.ExternalName, s.githubRepo) {
			svc, err := instance.Create(instance.Prepare(s.githubRepo, serv.Spec.ExternalName, "default", nil))
			if err != nil {
				return err
			}
			log.Printf("ServiceInstance-1: %s", svc.Name)
		}
		if strings.HasPrefix(serv.Spec.ExternalName, s.slackWorkspace) {
			svc, err := instance.Create(instance.Prepare(s.slackWorkspace, serv.Spec.ExternalName, "default", nil))
			if err != nil {
				return err
			}
			log.Printf("ServiceInstance-2: %s", svc.Name)
		}
		if serv.Spec.ExternalName == s.azureServiceName {
			raw := runtime.RawExtension{}
			err := raw.UnmarshalJSON([]byte(azureConfiguration))
			if err != nil {
				return apperrors.Internal("%s", err)
			}
			svc, err := instance.Create(instance.Prepare(s.azureServiceName, serv.Spec.ExternalName, "standard-s0", &raw))
			if err != nil {
				return apperrors.Internal("%s", err)
			}
			log.Printf("ServiceInstance-3: %s", svc.Name)
		}
	}
	return nil
}
