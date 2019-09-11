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

	runtime "k8s.io/apimachinery/pkg/runtime"

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
	azure := "azure-text-analytics"
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
		log.Printf("%s%s%s", string(chars), " ->- ", str)
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
	fatalOnError(err)
	log.Printf("SvcBindingUsage3: %s\n", svcBindingUsage3.Name)

	eb, err := eventbus.NewForConfig(k8sConfig)
	fatalOnError(err)
	sub, err := eb.Eventing().Subscriptions(namespace).Create(&subscription.Subscription{
		ObjectMeta: v1.ObjectMeta{
			Name:      "lambda-julia-lambda-issuesevent-v1",
			Namespace: namespace,
			Labels:    map[string]string{"Function": "julia-lambda"},
		},
		SubscriptionSpec: subscription.SubscriptionSpec{
			Endpoint:                      fmt.Sprintf("%s%s%s", "http://julia-lambda.", namespace, ":8080/"),
			EventType:                     "IssuesEvent",
			EventTypeVersion:              "v1",
			IncludeSubscriptionNameHeader: true,
			SourceID:                      "github-flying-seal-app",
		},
	})
	fatalOnError(err)
	log.Printf("Subscription: %s", sub.Name)
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

const funcCode = `const axios = require("axios");
const md = require("slackify-markdown");
const slackURL = process.env.GATEWAY_URL || "https://slack.com/api";
const githubURL = process.env.GITHUB_GATEWAY_URL 
const channelID = process.env.channelID || "node-best";

module.exports = {
    main: async function (event, context) {
        const githubPayload = event.data;
        if (githubPayload.action == "opened" || githubPayload.action == "edited") {

            let payload = await createPayload(githubPayload);

                try {
                    let issueURL = githubURL + '/repos/'+githubPayload.repository.full_name+'/issues/'+ githubPayload.issue.number 
                    console.log(issueURL)
                    let result = await setLabel(issueURL, payload);
                    console.log(result)
                } catch (error) {
                    console.error(error);
                }
            
        }
    }
};

async function checkIfBad(issueBody, issueTitle) {
    let result = await axios.post(process.env.textAnalyticsEndpoint + 'text/analytics/v2.1/sentiment',
    {documents: [{id: '1', text: issueBody}, {id: '2', text: issueTitle}]}, {headers: {...{'Ocp-Apim-Subscription-Key': process.env.textAnalyticsKey}}})
    return ((result.data.documents[0].score < 0.5) || (result.data.documents[1].score < 0.5))
}

function getLabels(labelsArray) {
    let labels = []
    labelsArray.map(label => labels.push(label.name))
    return labels
}

function createMessage(payload) {
  const blocks = [
    {
      type: "section",
      text: {
        type: "mrkdwn",
        text: "Hello @here!"
      }
    },
    {
      type: "section",
      text: {
        type: "mrkdwn",
        text: 'User *'+payload.issue.user.login+'* created an issue that might need a review: <$'+payload.issue.html_url+'|*#'+payload.issue.number+ payload.issue.title+'*>'
      }
    },
    {
      type: "section",
      text: {
        type: "mrkdwn",
        text: '*Issue* \n' + md(payload.issue.body)
      }
    }
  ];
  return blocks;
}

async function sendToSlack(payload){
    let msg = createMessage(payload)
     const config = {
    headers: {
      "Content-Type": "application/json;charset=UTF-8"
    }
  };
  const data = {
    channel: channelID,
    text: "New issue needs a review.",
    blocks: msg,
    link_names: true
  };
  let sendMsg = await axios.post(slackURL + "/chat.postMessage", data, config);
  return sendMsg;
}

async function createPayload(githubPayload) {
    let labels = getLabels(githubPayload.issue.labels)
    let sentiment = await checkIfBad(githubPayload.issue.body, githubPayload.issue.title)
    if (!sentiment)
    {
    labels = labels.filter(word => word != ':thinking: Review needed')    }
    else
    {
        labels.push(":thinking: Review needed")
        await sendToSlack(githubPayload)
        
    }
    const pld = {
        labels: labels
    }
    return pld;
}

async function setLabel(url, msg) {
    const config = {
        headers: {
            "Content-Type": "application/json;charset=UTF-8"
        }
    };
    let sendMsg = await axios.patch(url, msg, config);
    return sendMsg;
}`
