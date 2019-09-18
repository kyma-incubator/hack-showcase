package main

import (
	"log"
	"time"

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

// Config holds application configuration
type Config struct {
	Kubeconfig     string `envconfig:"APP,optional"`
	GithubURL      string `envconfig:"GITHUB_REPO"`
	SlackWorkspace string `envconfig:"SLACK_WORKSPACE"`
	Namespace      string `envconfig:"NAMESPACE"`
}

func main() {

	var cfg Config
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

	//ServiceInstance
	serviceCatalogClient := wrappers.NewServiceCatalogClient(svcClient)
	err = manager.CreateServiceInstances(serviceCatalogClient.Instance(cfg.Namespace), svcList)
	fatalOnError(err)

	//Function
	kubeless, err := kubeless.NewForConfig(k8sConfig)
	fatalOnError(err)
	kubelessClient := wrappers.NewKubelessClient(kubeless.Kubeless())
	err = manager.CreateFunction(kubelessClient.Function(cfg.Namespace))
	fatalOnError(err)

	//Other components have to wait for end of creating function
	time.Sleep(5 * time.Second)

	//ServiceBindings
	err = manager.CreateServiceBindings(serviceCatalogClient.Binding(cfg.Namespace))
	fatalOnError(err)

	//ServiceBindingUsages
	catalogClient, err := svcBind.NewForConfig(k8sConfig)
	fatalOnError(err)
	kymaServiceCatalogClient := wrappers.NewKymaServiceCatalogClient(catalogClient)
	err = manager.CreateServiceBindingUsages(kymaServiceCatalogClient.BindingUsage(cfg.Namespace))
	fatalOnError(err)

	//To create subscription resources above must be ready. Wait for their creation.
	time.Sleep(5 * time.Second)

	//Subscription
	bus, err := eventbus.NewForConfig(k8sConfig)
	fatalOnError(err)
	eventbusClient := wrappers.NewEventbusClient(bus.Eventing())
	err = manager.CreateSubscription(eventbusClient.Subscription(cfg.Namespace))
	fatalOnError(err)

}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newRestClientConfig(kubeConfigPath string) (*restclient.Config, error) {
	if kubeConfigPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	}

	return restclient.InClusterConfig()
}
