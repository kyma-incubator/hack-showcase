package main

import (
	"fmt"
	"log"
	"os"
	"time"

	//kubeless "github.com/kubeless/kubeless/pkg/utils"
	v1beta1kubeless "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	kubeless "github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	eventbus "github.com/kyma-project/kyma/components/event-bus/generated/push/clientset/versioned"
	autoscaling "k8s.io/api/autoscaling/v2beta1"
	core "k8s.io/api/core/v1"
	pts "k8s.io/api/core/v1"
	deplo "k8s.io/api/extensions/v1beta1"
	ios "k8s.io/apimachinery/pkg/util/intstr"

	svcCatalog "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"
	subscription "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	"github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	v1alpha1svc "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	svcBind "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"
	"github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
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
	namespace := "default"

	fmt.Printf("Nazwa repo: %s\n", githubRepo)
	fmt.Printf("Nazwa workspace: %s\n", slackWorkspace)

	var cfg Config
	err := envconfig.InitWithPrefix(&cfg, "APP")
	fatalOnError(err)

	// general k8s config
	k8sConfig, err := newRestClientConfig(cfg.Kubeconfig)
	fatalOnError(err)

	//ServiceBindingUsage Client
	svcBindClient, err := svcBind.NewForConfig(k8sConfig)
	fatalOnError(err)

	//ServiceCatalog Client
	svcClient, err := svcCatalog.NewForConfig(k8sConfig)
	fatalOnError(err)

	svcList, err := svcClient.ServiceClasses(namespace).List(v1.ListOptions{})
	fatalOnError(err)

	//create Service Instance
	for _, s := range svcList.Items {
		fmt.Println(s.Name)
		chars := []rune(s.Spec.ExternalName)
		str := string(chars[0 : len(chars)-6])
		if str == githubRepo {
			fmt.Println("działa git")
			svc, err := svcClient.ServiceInstances(namespace).Create(&v1beta1svc.ServiceInstance{
				ObjectMeta: v1.ObjectMeta{
					Name:      githubRepo + "inst",
					Namespace: namespace,
				},
				Spec: v1beta1svc.ServiceInstanceSpec{
					PlanReference: v1beta1svc.PlanReference{
						ServiceClassExternalName: s.Spec.ExternalName,
						ServicePlanExternalName:  namespace,
					},
				},
			})
			fatalOnError(err)
			fmt.Printf("Service Instance: %s, %s\n", svc.Name, svc.Status.ProvisionStatus)
		}
		if str == slackWorkspace {
			fmt.Println("działa slack")
			svc, err := svcClient.ServiceInstances(namespace).Create(&v1beta1svc.ServiceInstance{
				ObjectMeta: v1.ObjectMeta{
					Name:      slackWorkspace + "inst",
					Namespace: namespace,
				},
				Spec: v1beta1svc.ServiceInstanceSpec{
					PlanReference: v1beta1svc.PlanReference{
						ServiceClassExternalName: s.Spec.ExternalName,
						ServicePlanExternalName:  namespace,
					},
				},
			})
			fatalOnError(err)
			fmt.Printf("Service Instance: %s, %s\n", svc.Name, svc.Status.ProvisionStatus)
		}
	}

	c, err := kubeless.NewForConfig(k8sConfig)
	fatalOnError(err)
	fun, err := c.Kubeless().Functions(namespace).Create(&v1beta1kubeless.Function{
		ObjectMeta: v1.ObjectMeta{
			Name:      "julia-lambda",
			Namespace: namespace,
			Labels:    map[string]string{"app": "julia"},
		},
		Spec: v1beta1kubeless.FunctionSpec{
			Deps: `{
				"name": "example-1",
				"version": "0.0.1",
				"dependencies": {
				  "request": "^2.85.0"
				}
			}`,
			Function: `mmodule.exports = { 
				main: function (event, context) {
					console.log("Issue");
				} 
			};`,
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
					"function":   "julia-lambda",
				},
			},
			Deployment: deplo.Deployment{
				Spec: deplo.DeploymentSpec{
					Template: pts.PodTemplateSpec{
						Spec: pts.PodSpec{
							Containers: []pts.Container{pts.Container{
								Name:      "",
								Resources: pts.ResourceRequirements{},
							}},
						},
					},
				},
			},
		},
	})
	fatalOnError(err)
	fmt.Printf("Function: %s", fun.Name)

	time.Sleep(5 * time.Second)

	//ServiceBinding

	fmt.Println("Building svcBinding...")
	svcBinding, err := svcClient.ServiceBindings(namespace).Create(&v1beta1.ServiceBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      githubRepo + "bind",
			Namespace: namespace,
			Labels: map[string]string{
				"Function": "julia-lambda",
			},
		},
		Spec: v1beta1svc.ServiceBindingSpec{
			InstanceRef: v1beta1svc.LocalObjectReference{
				Name: githubRepo + "inst",
			},
		},
	})
	fatalOnError(err)
	fmt.Printf("SvcBinding: %s\n", svcBinding.Name)

	svcBinding2, err := svcClient.ServiceBindings(namespace).Create(&v1beta1.ServiceBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      slackWorkspace + "bind",
			Namespace: namespace,
			Labels: map[string]string{
				"Function": "julia-lambda",
			},
		},
		Spec: v1beta1svc.ServiceBindingSpec{
			InstanceRef: v1beta1svc.LocalObjectReference{
				Name: githubRepo + "inst",
			},
		},
	})
	fatalOnError(err)
	fmt.Printf("SvcBinding2: %s\n", svcBinding2.Name)

	//Service Binding Usage
	fmt.Println("Building svcBindingUsage...")
	svcBindingUsage, err := svcBindClient.ServiceBindingUsages(namespace).Create(&v1alpha1.ServiceBindingUsage{
		TypeMeta: v1.TypeMeta{
			Kind:       "ServiceBindingUsage",
			APIVersion: "servicecatalog.kyma-project.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      githubRepo + "bu",
			Namespace: namespace,
			Labels: map[string]string{
				"Function":       "julia-lambda",
				"ServiceBinding": githubRepo + "bind",
			},
		},
		Spec: v1alpha1svc.ServiceBindingUsageSpec{
			ServiceBindingRef: v1alpha1svc.LocalReferenceByName{
				Name: githubRepo + "bind",
			},
			UsedBy: v1alpha1svc.LocalReferenceByKindAndName{
				Kind: "function",
				Name: "julia-lambda",
			},
			Parameters: &v1alpha1svc.Parameters{
				EnvPrefix: &v1alpha1svc.EnvPrefix{
					Name: "GITHUB_",
				},
			},
		},
	})
	fatalOnError(err)
	fmt.Printf("SvcBindingUsage: %s\n", svcBindingUsage.Name)

	svcBindingUsage2, err := svcBindClient.ServiceBindingUsages(namespace).Create(&v1alpha1.ServiceBindingUsage{
		TypeMeta: v1.TypeMeta{
			Kind:       "ServiceBindingUsage",
			APIVersion: "servicecatalog.kyma-project.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      slackWorkspace + "bu",
			Namespace: namespace,
			Labels: map[string]string{
				"Function":       "julia-lambda",
				"ServiceBinding": slackWorkspace + "bind",
			},
		},
		Spec: v1alpha1svc.ServiceBindingUsageSpec{
			ServiceBindingRef: v1alpha1svc.LocalReferenceByName{
				Name: slackWorkspace + "bind",
			},
			UsedBy: v1alpha1svc.LocalReferenceByKindAndName{
				Name: "julia-lambda",
				Kind: "function",
			},
			Parameters: &v1alpha1svc.Parameters{
				EnvPrefix: &v1alpha1svc.EnvPrefix{
					Name: "",
				},
			},
		},
	})
	fatalOnError(err)
	fmt.Printf("SvcBindingUsage2: %s\n", svcBindingUsage2.Name)

	eb, err := eventbus.NewForConfig(k8sConfig)
	fatalOnError(err)
	sub, err := eb.Eventing().Subscriptions(namespace).Create(&subscription.Subscription{
		ObjectMeta: v1.ObjectMeta{
			Name:      "lambda-julia-lambda-issuesevent-v1",
			Namespace: namespace,
			Labels:    map[string]string{"Function": "julia-lambda"},
		},
		SubscriptionSpec: subscription.SubscriptionSpec{
			Endpoint:                      fmt.Sprintf("%s%s%s", "http://julia-lambda.", namespace, "8080/"),
			EventType:                     "IssuesEvent",
			EventTypeVersion:              "v1",
			IncludeSubscriptionNameHeader: true,
			SourceID:                      "julia-lambda-app",
		},
	})
	fatalOnError(err)
	fmt.Printf("Subscription: %s", sub.Name)
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
