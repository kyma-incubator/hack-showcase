package main

import (
	"log"
	"os"
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

// Config holds application configuration
type Config struct {
	Kubeconfig string `envconfig:"optional"`
}

func main() {
	githubRepo := os.Getenv("GITHUB_REPO")
	slackWorkspace := os.Getenv("SLACK_WORKSPACE")
	namespace := os.Getenv("NAMESPACE")
	azure := "azure-text-analytics"

	log.Printf("Github url: %s\n", githubRepo)
	log.Printf("Slack workspace: %s\n", slackWorkspace)
	log.Printf("Workspace: %s", namespace)
	log.Printf("Azure: %s", azure)

	var cfg Config
	err := envconfig.InitWithPrefix(&cfg, "APP")
	fatalOnError(err)

	// general k8s config
	k8sConfig, err := newRestClientConfig(cfg.Kubeconfig)
	fatalOnError(err)

	//ServiceCatalog Client
	svcClient, err := svcCatalog.NewForConfig(k8sConfig)
	fatalOnError(err)
	svcList, err := svcClient.ServiceClasses(namespace).List(v1.ListOptions{})
	fatalOnError(err)

	//Create scenario Manager
	manager := manager.NewManager(namespace, githubRepo, slackWorkspace, azure)

	//ServiceInstance
	instance := wrappers.NewServiceCatalogClient(svcClient).Instance(namespace)
	err = manager.CreateServiceInstances(instance, svcList)
	fatalOnError(err)

	//Function
	kubeless, err := kubeless.NewForConfig(k8sConfig)
	fatalOnError(err)
	function := wrappers.NewKubelessWrapper(kubeless.Kubeless()).Function(namespace)
	err = manager.CreateFunction(function)
	fatalOnError(err)

	time.Sleep(5 * time.Second)
	//ServiceBindings
	bindingManager := wrappers.NewServiceCatalogClient(svcClient).Binding(namespace)
	err = manager.CreateServiceBindings(bindingManager)
	fatalOnError(err)

	//ServiceBindingUsages
	catalogClient, err := svcBind.NewForConfig(k8sConfig)
	fatalOnError(err)
	bindingUsage := wrappers.NewKymaServiceCatalogWrapper(catalogClient).BindingUsage(namespace)
	err = manager.CreateServiceBindingUsages(bindingUsage)
	fatalOnError(err)

	//To create subscription resources above must be ready. Wait for their creation.
	time.Sleep(5 * time.Second)

	//Subscription
	bus, err := eventbus.NewForConfig(k8sConfig)
	fatalOnError(err)
	subscription := wrappers.NewSubscriptionManager(bus.Eventing()).Subscription(namespace)
	err = manager.CreateSubscription(subscription)
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
