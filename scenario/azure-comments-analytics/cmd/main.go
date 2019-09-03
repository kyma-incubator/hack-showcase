package main

import (
	"fmt"
	"log"

	"github.com/kyma-project/kyma/components/api-controller/pkg/apis/gateway.kyma-project.io/v1alpha2"
	gatewayClientset "github.com/kyma-project/kyma/components/api-controller/pkg/clients/gateway.kyma-project.io/clientset/versioned"
	"github.com/vrischmann/envconfig"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClientset "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Config holds application configuration
type Config struct {
	Kubeconfig string `envconfig:"optional"`
}

func main() {
	var cfg Config
	err := envconfig.InitWithPrefix(&cfg, "APP")
	fatalOnError(err)

	// general k8s config
	k8sConfig, err := newRestClientConfig(cfg.Kubeconfig)
	fatalOnError(err)

	// k8s Clientset
	k8sCli, err := k8sClientset.NewForConfig(k8sConfig)
	fatalOnError(err)

	// Kyma API kind Client
	kymaGatewayClient, err := gatewayClientset.NewForConfig(k8sConfig)
	fatalOnError(err)

	// usage
	svc, err := k8sCli.CoreV1().Services("default").Get("kubernetes", v1.GetOptions{})
	fatalOnError(err)
	fmt.Printf("Succesfuuly fetched Service: %s\n", svc.Name)

	createAPI, err := kymaGatewayClient.GatewayV1alpha2().Apis("default").Create(&v1alpha2.Api{
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-api",
			Namespace: "default",
		},
		Spec: v1alpha2.ApiSpec{
			Authentication: []v1alpha2.AuthenticationRule{},
			Hostname:       "http://please-change-me.com",
			Service: v1alpha2.Service{
				Name: "please-change-me",
				Port: 8443,
			},
		},
	})
	fatalOnError(err)

	fmt.Printf("Created API successfully: %\n", createAPI)
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
