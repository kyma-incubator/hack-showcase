package main

import (
	"log"
	"os"
	"time"

	//kubeless "github.com/kubeless/kubeless/pkg/utils"

	kubeless "github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	eventbus "github.com/kyma-project/kyma/components/event-bus/generated/push/clientset/versioned"
	svcBind "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	svcCatalog "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/typed/servicecatalog/v1beta1"
	wrappers "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/clientwrappers"
	"github.com/vrischmann/envconfig"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
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
	azure := "azure-text-analytics"

	fake.NewSimpleClientset()

	log.Printf("Nazwa repo: %s\n", githubRepo)
	log.Printf("Nazwa workspace: %s\n", slackWorkspace)
	log.Printf("Workspace: %s", os.Getenv("NAMESPACE"))

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

	log.Println("No i wylistowalem :D")

	//create Service Instance
	instanceManager := wrappers.NewServiceCatalogClient(svcClient).Instance(namespace)
	for _, s := range svcList.Items {
		chars := []rune(s.Spec.ExternalName)
		str := string(chars[0 : len(chars)-6])
		if str == githubRepo {
			svc, err := instanceManager.Create(instanceManager.GetEventBody(githubRepo, s.Spec.ExternalName, "default", nil))
			fatalOnError(err)
			log.Printf("ServiceInstance-1: %s", svc.Name)
		}
		if str == slackWorkspace {
			svc, err := instanceManager.Create(instanceManager.GetEventBody(slackWorkspace, s.Spec.ExternalName, "default", nil))
			fatalOnError(err)
			log.Printf("ServiceInstance-2: %s", svc.Name)
		}
		if string(chars) == azure {
			raw := runtime.RawExtension{}
			err := raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
			fatalOnError(err)
			svc, err := instanceManager.Create(instanceManager.GetEventBody(azure, s.Spec.ExternalName, "standard-s0", &raw))
			fatalOnError(err)
			log.Printf("ServiceInstance-3: %s", svc.Name)
		}
	}

	//==================== DONE
	kubeless, err := kubeless.NewForConfig(k8sConfig)
	fatalOnError(err)
	functionManager := wrappers.NewKubelessWrapper(kubeless.Kubeless()).Function(namespace)
	funct, err := functionManager.Create(functionManager.GetEventBody())
	fatalOnError(err)
	log.Printf("Function: %s\n", funct.Name)

	time.Sleep(5 * time.Second)
	//==================== DONE
	bindingManager := wrappers.NewServiceCatalogClient(svcClient).Binding(namespace)
	bind1, err := bindingManager.Create(bindingManager.GetEventBody(githubRepo))
	log.Printf("SvcBinding-1: %s\n", bind1.Name)
	bind2, err := bindingManager.Create(bindingManager.GetEventBody(slackWorkspace))
	fatalOnError(err)
	log.Printf("SvcBinding-2: %s\n", bind2.Name)
	bind3, err := bindingManager.Create(bindingManager.GetEventBody(azure))
	fatalOnError(err)
	log.Printf("SvcBinding-3: %s\n", bind3.Name)
	//==================== DONE
	catalogClient, err := svcBind.NewForConfig(k8sConfig)
	fatalOnError(err)
	bindingUsageManager := wrappers.NewKymaServiceCatalogWrapper(catalogClient)
	bindingUsage := bindingUsageManager.BindingUsage(namespace)
	usage1, err := bindingUsage.Create(bindingUsage.GetEventBody(githubRepo, "GITHUB_"))
	fatalOnError(err)
	log.Printf("SvcBindingUsage-1: %s\n", usage1.Name)

	usage2, err := bindingUsage.Create(bindingUsage.GetEventBody(slackWorkspace, ""))
	fatalOnError(err)
	log.Printf("SvcBindingUsage-2: %s\n", usage2.Name)

	usage3, err := bindingUsage.Create(bindingUsage.GetEventBody(azure, ""))
	fatalOnError(err)
	log.Printf("SvcBindingUsage-3: %s\n", usage3.Name)

	//===================== DONE
	bus, err := eventbus.NewForConfig(k8sConfig)
	fatalOnError(err)
	subscriptionManager := wrappers.NewSubscriptionManager(bus.Eventing()).Subscription(namespace)
	subscribe, err := subscriptionManager.Create(subscriptionManager.GetSubscription(githubRepo))
	fatalOnError(err)
	log.Printf("Subscription: %s", subscribe.Name)
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

/*
	//ServiceBindingUsage Client
	svcBindClient, err := svcBind.NewForConfig(k8sConfig)
	fatalOnError(err)

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
	log.Printf("SvcBindingUsage: %s\n", svcBindingUsage.Name)

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
	log.Printf("SvcBindingUsage2: %s\n", svcBindingUsage2.Name)

	svcBindingUsage3, err := svcBindClient.ServiceBindingUsages(namespace).Create(&v1alpha1.ServiceBindingUsage{
		TypeMeta: v1.TypeMeta{
			Kind:       "ServiceBindingUsage",
			APIVersion: "servicecatalog.kyma-project.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      azure + "bu",
			Namespace: namespace,
			Labels: map[string]string{
				"Function":       "julia-lambda",
				"ServiceBinding": azure + "bind",
			},
		},
		Spec: v1alpha1svc.ServiceBindingUsageSpec{
			ServiceBindingRef: v1alpha1svc.LocalReferenceByName{
				Name: azure + "bind",
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
*/

/*
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
	log.Printf("SvcBinding: %s\n", svcBinding.Name)

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
				Name: slackWorkspace + "inst",
			},
		},
	})
	fatalOnError(err)
	log.Printf("SvcBinding2: %s\n", svcBinding2.Name)

	svcBinding3, err := svcClient.ServiceBindings(namespace).Create(&v1beta1.ServiceBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      azure + "bind",
			Namespace: namespace,
			Labels: map[string]string{
				"Function": "julia-lambda",
			},
		},
		Spec: v1beta1svc.ServiceBindingSpec{
			InstanceRef: v1beta1svc.LocalObjectReference{
				Name: azure + "inst",
			},
		},
	})
	fatalOnError(err)
	log.Printf("SvcBinding3: %s\n", svcBinding3.Name)
*/

/*
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
				"dependencies": {
			  "axios": "^0.19.0",
			  "slackify-markdown": "^1.1.1"
			}
		  }`,
			Function:            funcCode,
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
	log.Printf("Function: %s", fun.Name)
*/

/*
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
						ServicePlanExternalName:  "default",
					},
				},
			})
			fatalOnError(err)
			log.Printf("Service Instance: %s, %s\n -----< Plan Name: %s    \n --- %v", svc.Name, svc.Status.ProvisionStatus, s.Spec.ExternalName, svc)
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
						ServicePlanExternalName:  "default",
					},
				},
			})
			fatalOnError(err)
			log.Printf("Service Instance: %s, %s\n -----< Plan Name: %s    \n --- %v", svc.Name, svc.Status.ProvisionStatus, s.Spec.ExternalName, svc)
		}
		if string(chars) == azure {
			raw := runtime.RawExtension{}
			raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
			fmt.Println("Dziala Azure")
			svc, err := svcClient.ServiceInstances(namespace).Create(&v1beta1svc.ServiceInstance{
				ObjectMeta: v1.ObjectMeta{
					Name:      azure + "inst",
					Namespace: namespace,
				},
				Spec: v1beta1svc.ServiceInstanceSpec{
					Parameters: &raw,
					PlanReference: v1beta1svc.PlanReference{
						ServiceClassExternalName: s.Spec.ExternalName,
						ServicePlanExternalName:  "standard-s0",
					},
				},
			})
			fatalOnError(err)
			log.Printf("Service Instance: %s, %s\n -----< Plan Name: %s    \n --- %v", svc.Name, svc.Status.ProvisionStatus, s.Spec.ExternalName, svc)
		}
*/
