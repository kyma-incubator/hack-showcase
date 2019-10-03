package manager

import (
	"log"
	"strings"

	kubelessbeta1 "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	eventingalpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	kymaservicecatalogaplha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	servicecatalogbeta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"

	kubeless "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	"github.com/kyma-incubator/github-slack-connectors/scenario/github-issue-sentiment-analysis/internal/k8scomponents"
	eventing "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	kymaservicecatalog "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	servicecatalog "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const azureConfiguration = `{"location": "westeurope","resourceGroup": "github-slack-scenario"}`

//Manager include important methods to deploy all k8s and kymas components to realize github-slack-connectors scenario
type Manager interface {
	DeleteAllComponents(installedComponents *InstalledComponents, clientWrappers *Wrappers)
	CreateFunction(function k8scomponents.Function) ([]kubeless.Function, error)
	CreateServiceBindings(binding k8scomponents.Binding) ([]servicecatalog.ServiceBinding, error)
	CreateSubscription(subscription k8scomponents.Subscription) ([]eventing.Subscription, error)
	CreateServiceBindingUsages(bindingUsage k8scomponents.BindingUsage) ([]kymaservicecatalog.ServiceBindingUsage, error)
	CreateServiceInstances(instance k8scomponents.ServiceInstance, serviceClassList *servicecatalog.ServiceClassList) ([]servicecatalog.ServiceInstance, error)
}
type manager struct {
	githubRepo       string
	slackWorkspace   string
	azureServiceName string
	namespace        string
	lambdaName       string
	channelName      string
}

// InstalledComponents allow you to store informations about installed components
type InstalledComponents struct {
	Subscriptions        []eventingalpha1.Subscription
	ServiceBindingUsages []kymaservicecatalogaplha1.ServiceBindingUsage
	Functions            []kubelessbeta1.Function
	ServiceInstances     []servicecatalogbeta1.ServiceInstance
	ServiceBindings      []servicecatalogbeta1.ServiceBinding
}

// Wrappers store all client wrappers
type Wrappers struct {
	Subscription    k8scomponents.Subscription
	BindingUsage    k8scomponents.BindingUsage
	Binding         k8scomponents.Binding
	Function        k8scomponents.Function
	ServiceInstance k8scomponents.ServiceInstance
}

//NewManager create and return new manager struct
func NewManager(namespace string, githubRepo string, slackWorkspace string, azureServiceName string, channelName string) Manager {
	return &manager{
		namespace:        namespace,
		githubRepo:       githubRepo,
		slackWorkspace:   slackWorkspace,
		azureServiceName: azureServiceName,
		lambdaName:       githubRepo[7:] + "-lambda", //Due to Kyma's requirements lambda's name has to be short - it's trimmed here
		channelName:      channelName,
	}
}

func (s *manager) CreateSubscription(subscription k8scomponents.Subscription) ([]eventing.Subscription, error) {
	var subscriptions []eventing.Subscription
	subscribe, err := subscription.Create(subscription.Prepare(s.githubRepo, s.lambdaName))
	if err != nil {
		return subscriptions, err
	}
	log.Printf("Subscription: %s", subscribe.Name)
	subscriptions = append(subscriptions, *subscribe)
	return subscriptions, nil
}

func (s *manager) CreateServiceBindingUsages(bindingUsage k8scomponents.BindingUsage) ([]kymaservicecatalog.ServiceBindingUsage, error) {
	var serviceBindingUsages []kymaservicecatalog.ServiceBindingUsage
	usage1, err := bindingUsage.Create(bindingUsage.Prepare(s.githubRepo, "GITHUB_", s.lambdaName))
	if err != nil {
		return serviceBindingUsages, err
	}
	log.Printf("SvcBindingUsage-1: %s\n", usage1.Name)
	serviceBindingUsages = append(serviceBindingUsages, *usage1)

	usage2, err := bindingUsage.Create(bindingUsage.Prepare(s.slackWorkspace, "", s.lambdaName))
	if err != nil {
		return serviceBindingUsages, err
	}
	log.Printf("SvcBindingUsage-2: %s\n", usage2.Name)
	serviceBindingUsages = append(serviceBindingUsages, *usage2)

	usage3, err := bindingUsage.Create(bindingUsage.Prepare(s.azureServiceName, "", s.lambdaName))
	if err != nil {
		return serviceBindingUsages, err
	}
	log.Printf("SvcBindingUsage-3: %s\n", usage3.Name)
	serviceBindingUsages = append(serviceBindingUsages, *usage3)
	return serviceBindingUsages, nil
}

func (s *manager) CreateServiceBindings(binding k8scomponents.Binding) ([]servicecatalog.ServiceBinding, error) {
	var serviceBindings []servicecatalog.ServiceBinding
	bind1, err := binding.Create(binding.Prepare(s.githubRepo, s.lambdaName))
	if err != nil {
		return serviceBindings, err
	}
	log.Printf("SvcBinding-1: %s\n", bind1.Name)
	serviceBindings = append(serviceBindings, *bind1)

	bind2, err := binding.Create(binding.Prepare(s.slackWorkspace, s.lambdaName))
	if err != nil {
		return serviceBindings, err
	}
	log.Printf("SvcBinding-2: %s\n", bind2.Name)
	serviceBindings = append(serviceBindings, *bind2)

	bind3, err := binding.Create(binding.Prepare(s.azureServiceName, s.lambdaName))
	if err != nil {
		return serviceBindings, err
	}
	log.Printf("SvcBinding-3: %s\n", bind3.Name)
	serviceBindings = append(serviceBindings, *bind3)

	return serviceBindings, nil
}

func (s *manager) CreateFunction(function k8scomponents.Function) ([]kubeless.Function, error) {
	var functions []kubeless.Function
	funct, err := function.Create(function.Prepare(s.githubRepo, s.lambdaName, s.channelName))
	if err != nil {
		return functions, err
	}
	log.Printf("Function: %s\n", funct.Name)
	functions = append(functions, *funct)
	return functions, nil
}

func (s *manager) CreateServiceInstances(instance k8scomponents.ServiceInstance, serviceClassList *servicecatalog.ServiceClassList) ([]servicecatalog.ServiceInstance, error) {
	//ServiceClass ExternalName suffix is generated randomly, but its prefix is based on name provided by user.
	//Looking for ServiceClass with matching prefix on which basis ServiceInstance should be created.
	var serviceInstances []servicecatalog.ServiceInstance
	for _, serv := range serviceClassList.Items {
		if strings.HasPrefix(serv.Spec.ExternalName, s.githubRepo) {
			svc, err := instance.Create(instance.Prepare(s.githubRepo, serv.Spec.ExternalName, "default", nil))
			if err != nil {
				return serviceInstances, err
			}
			log.Printf("ServiceInstance-1: %s", svc.Name)
			serviceInstances = append(serviceInstances, *svc)
		}
		if strings.HasPrefix(serv.Spec.ExternalName, s.slackWorkspace) {
			svc, err := instance.Create(instance.Prepare(s.slackWorkspace, serv.Spec.ExternalName, "default", nil))
			if err != nil {
				return serviceInstances, err
			}
			log.Printf("ServiceInstance-2: %s", svc.Name)
			serviceInstances = append(serviceInstances, *svc)
		}
		if serv.Spec.ExternalName == s.azureServiceName {
			raw := runtime.RawExtension{}
			err := raw.UnmarshalJSON([]byte(azureConfiguration))
			if err != nil {
				return serviceInstances, err
			}
			svc, err := instance.Create(instance.Prepare(s.azureServiceName, serv.Spec.ExternalName, "standard-s0", &raw))
			if err != nil {
				return serviceInstances, err
			}
			log.Printf("ServiceInstance-3: %s", svc.Name)
			serviceInstances = append(serviceInstances, *svc)
		}
	}
	return serviceInstances, nil
}

func (s *manager) DeleteAllComponents(installedComponents *InstalledComponents, clientWrappers *Wrappers) {
	deleteOptions := &v1.DeleteOptions{}

	for _, element := range installedComponents.Subscriptions {
		err := clientWrappers.Subscription.Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.ServiceBindingUsages {
		err := clientWrappers.BindingUsage.Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.ServiceBindings {
		err := clientWrappers.Binding.Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.Functions {
		err := clientWrappers.Function.Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.ServiceInstances {
		err := clientWrappers.ServiceInstance.Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}
}
