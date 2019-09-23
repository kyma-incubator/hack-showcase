package main

import (
	"errors"
	"log"
	"time"

	kubelessbeta1 "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	eventingalpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	kymaservicecatalogaplha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	servicecatalogbeta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"

	kubeless "github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	eventbus "github.com/kyma-project/kyma/components/event-bus/generated/push/clientset/versioned"
	svcBind "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"

	svcCatalog "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"
	wrappers "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/clientwrappers"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/manager"
	"github.com/vrischmann/envconfig"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const azureClassName = "azure-text-analytics"

var installedComponents InstalledComponents
var clientWrappers Wrappers
var cfg Config

// Config holds application configuration
type Config struct {
	Kubeconfig     string `envconfig:"APP,optional"`
	GithubURL      string `envconfig:"GITHUB_REPO"`
	SlackWorkspace string `envconfig:"SLACK_WORKSPACE"`
	Namespace      string `envconfig:"NAMESPACE"`
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
	Namespace          string
	Eventbus           wrappers.EventbusWrapper
	Kubeless           wrappers.KubelessWrapper
	ServiceCatalog     wrappers.ServiceCatalogWrapper
	KymaServiceCatalog wrappers.KymaServiceCatalogWrapper
}

func main() {
	err := envconfig.Init(&cfg)
	fatalOnError(err)

	log.Printf("Kubeconfig: %s", cfg.Kubeconfig)
	log.Printf("Github url: %s\n", cfg.GithubURL)
	log.Printf("Slack workspace: %s\n", cfg.SlackWorkspace)
	log.Printf("Workspace: %s", cfg.Namespace)
	log.Printf("Azure: %s", azureClassName)

	// general k8s config
	k8sConfig, err := newRestClientConfig(cfg.Kubeconfig)
	fatalOnError(err)

	//ServiceCatalog Client
	svcClient, err := svcCatalog.NewForConfig(k8sConfig)
	fatalOnError(err)
	svcList, err := svcClient.ServiceClasses(cfg.Namespace).List(v1.ListOptions{})
	fatalOnError(err)

	//Create scenario Manager
	manager := manager.NewManager(cfg.Namespace, cfg.GithubURL, cfg.SlackWorkspace, azureClassName)

	//Contain all created components

	//ServiceInstance
	clientWrappers.ServiceCatalog = wrappers.NewServiceCatalogClient(svcClient)
	installedComponents.ServiceInstances, err = manager.CreateServiceInstances(clientWrappers.ServiceCatalog.Instance(cfg.Namespace), svcList)
	fatalOnError(err)

	//Function
	kubeless, err := kubeless.NewForConfig(k8sConfig)
	fatalOnError(err)
	clientWrappers.Kubeless = wrappers.NewKubelessClient(kubeless.Kubeless())
	installedComponents.Functions, err = manager.CreateFunction(clientWrappers.Kubeless.Function(cfg.Namespace))
	fatalOnError(err)

	//Other components have to wait for end of creating function
	time.Sleep(5 * time.Second)

	//ServiceBindings
	installedComponents.ServiceBindings, err = manager.CreateServiceBindings(clientWrappers.ServiceCatalog.Binding(cfg.Namespace))
	fatalOnError(err)

	//ServiceBindingUsages
	catalogClient, err := svcBind.NewForConfig(k8sConfig)
	fatalOnError(err)
	clientWrappers.KymaServiceCatalog = wrappers.NewKymaServiceCatalogClient(catalogClient)
	installedComponents.ServiceBindingUsages, err = manager.CreateServiceBindingUsages(clientWrappers.KymaServiceCatalog.BindingUsage(cfg.Namespace))
	fatalOnError(err)

	//To create subscription resources above must be ready. Wait for their creation.
	time.Sleep(5 * time.Second)

	//Subscription
	bus, err := eventbus.NewForConfig(k8sConfig)
	fatalOnError(err)
	clientWrappers.Eventbus = wrappers.NewEventbusClient(bus.Eventing())
	installedComponents.Subscriptions, err = manager.CreateSubscription(clientWrappers.Eventbus.Subscription(cfg.Namespace))
	fatalOnError(err)

	fatalOnError(errors.New("Now Uninstalling"))
}

func fatalOnError(err error) {
	if err != nil {
		uninstallComponents()
		log.Fatal(err)
	}
}

func newRestClientConfig(kubeConfigPath string) (*restclient.Config, error) {
	if kubeConfigPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	}

	return restclient.InClusterConfig()
}

var deleteOptions *v1.DeleteOptions

func uninstallComponents() {
	deleteOptions = &v1.DeleteOptions{}

	for _, element := range installedComponents.Subscriptions {
		err := clientWrappers.Eventbus.Subscription(cfg.Namespace).Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.ServiceBindingUsages {
		err := clientWrappers.KymaServiceCatalog.BindingUsage(cfg.Namespace).Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.ServiceBindings {
		err := clientWrappers.ServiceCatalog.Binding(cfg.Namespace).Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.Functions {
		err := clientWrappers.Kubeless.Function(cfg.Namespace).Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}

	for _, element := range installedComponents.ServiceInstances {
		err := clientWrappers.ServiceCatalog.Instance(cfg.Namespace).Delete(element.ObjectMeta.Name, deleteOptions)
		if err != nil {
			log.Printf("%s can't be removed. Please, remove it manually: %s", element.ObjectMeta.Name, err.Error())
		} else {
			log.Printf("%s removed", element.ObjectMeta.Name)
		}
	}
}
