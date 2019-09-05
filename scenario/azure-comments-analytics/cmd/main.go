package main

import (
	"fmt"
	"log"
	"os"

	//kubeless "github.com/kubeless/kubeless/pkg/utils"
	v1beta1kubeless "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	kubeless "github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	autoscaling "k8s.io/api/autoscaling/v2beta1"
	core "k8s.io/api/core/v1"
	pts "k8s.io/api/core/v1"
	deplo "k8s.io/api/extensions/v1beta1"
	ios "k8s.io/apimachinery/pkg/util/intstr"

	svcCatalog "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"
	v1beta1svc "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
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

	fmt.Printf("Nazwa repo: %s\n", githubRepo)
	fmt.Printf("Nazwa workspace: %s\n", slackWorkspace)

	var cfg Config
	err := envconfig.InitWithPrefix(&cfg, "APP")
	fatalOnError(err)

	// general k8s config
	k8sConfig, err := newRestClientConfig(cfg.Kubeconfig)
	fatalOnError(err)

	app := make(map[string]string)
	app["connected-app"] = "github-colunira-podejmijtest"

	svcClient, err := svcCatalog.NewForConfig(k8sConfig)
	fatalOnError(err)

	svcList, err := svcClient.ServiceClasses("default").List(v1.ListOptions{})
	fatalOnError(err)

	//create Service Instance
	for _, s := range svcList.Items {
		fmt.Println(s.Name)
		chars := []rune(s.Spec.ExternalName)
		str := string(chars[0 : len(chars)-6])
		if str == githubRepo {
			fmt.Println("działa git")
			svc, err := svcClient.ServiceInstances("default").Create(&v1beta1svc.ServiceInstance{
				ObjectMeta: v1.ObjectMeta{
					Name:      githubRepo + "g",
					Namespace: "default",
				},
				Spec: v1beta1svc.ServiceInstanceSpec{
					PlanReference: v1beta1svc.PlanReference{
						ServiceClassExternalName: s.Spec.ExternalName,
						ServicePlanExternalName:  "default",
					},
				},
			})
			fatalOnError(err)
			fmt.Printf("Service Instance: %s, %s\n", svc.Name, svc.Status.ProvisionStatus)
		}
		if str == slackWorkspace {
			fmt.Println("działa slack")
			svc, err := svcClient.ServiceInstances("default").Create(&v1beta1svc.ServiceInstance{
				ObjectMeta: v1.ObjectMeta{
					Name:      slackWorkspace + "g",
					Namespace: "default",
				},
				Spec: v1beta1svc.ServiceInstanceSpec{
					PlanReference: v1beta1svc.PlanReference{
						ServiceClassExternalName: s.Spec.ExternalName,
						ServicePlanExternalName:  "default",
					},
				},
			})
			fatalOnError(err)
			fmt.Printf("Service Instance: %s, %s\n", svc.Name, svc.Status.ProvisionStatus)
		}
	}

	c, err := kubeless.NewForConfig(k8sConfig)
	c.Kubeless().Functions("default").Create(&v1beta1kubeless.Function{
		ObjectMeta: v1.ObjectMeta{
			Name:      "julia-the-lambda",
			Namespace: "default",
			Labels:    map[string]string{"app": "no-jakas-apka"},
		},
		Spec: v1beta1kubeless.FunctionSpec{
			Deps: `{
				"name": "example-1",
				"version": "0.0.1",
				"dependencies": {
				  "request": "^2.85.0"
				}
			}`,
			Function: `module.exports = { main: function (event, context) {
				console.log("Issue opened")
			} }`,
			FunctionContentType: "text",
			Handler:             "handler.main",
			Timeout:             "",
			HorizontalPodAutoscaler: autoscaling.HorizontalPodAutoscaler{
				Spec: autoscaling.HorizontalPodAutoscalerSpec{
					MaxReplicas: 0,
				},
			},
			Runtime: "nodejs8",
			ServiceSpec: core.ServiceSpec{
				Ports: []core.ServicePort{core.ServicePort{
					Name:       "http-function-port",
					Port:       8080,
					Protocol:   "TCP",
					TargetPort: ios.FromInt(8080),
				}},
				Selector: map[string]string{
					"created-by": "kubeless",
					"function":   "julia-the-lambda",
				},
			},
			Deployment: deplo.Deployment{
				Spec: deplo.DeploymentSpec{
					Template: pts.PodTemplateSpec{
						Spec: pts.PodSpec{
							Containers: []pts.Container{pts.Container{
								Name: "",
								//Resources: {},
							}},
						},
					},
				},
			},
		},
	})

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
