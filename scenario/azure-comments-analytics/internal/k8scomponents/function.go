package k8scomponents

import (
	v1beta1kubeless "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	//"github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	autoscaling "k8s.io/api/autoscaling/v2beta1"
	core "k8s.io/api/core/v1"
	pts "k8s.io/api/core/v1"
	deplo "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ios "k8s.io/apimachinery/pkg/util/intstr"
)

//FunctionInterface describe constructors argument and include important methods of Functions
type FunctionInterface interface {
	Create(*v1beta1kubeless.Function) (*v1beta1kubeless.Function, error)
}

//Function describe function struct
type Function interface {
	Create(body *v1beta1kubeless.Function) (*v1beta1kubeless.Function, apperrors.AppError)
	GetEventBody(name string) *v1beta1kubeless.Function
}

type function struct {
	functionInterface FunctionInterface
	namespace         string
}

//NewFunction returns new function struck
func NewFunction(functionInterface FunctionInterface, nspace string) Function {
	return &function{functionInterface: functionInterface, namespace: nspace}
}

func (s *function) Create(body *v1beta1kubeless.Function) (*v1beta1kubeless.Function, apperrors.AppError) {
	data, err := s.functionInterface.Create(body)
	if err != nil {
		return nil, apperrors.WrongInput("Can not create Function: %s", err)
	}
	return data, nil
}

func (s *function) GetEventBody(name string) *v1beta1kubeless.Function {
	return &v1beta1kubeless.Function{
		ObjectMeta: v1.ObjectMeta{
			Name:      name[7:] + "-lambda",
			Namespace: s.namespace,
			Labels:    map[string]string{"app": name + "-app"},
		},
		Spec: v1beta1kubeless.FunctionSpec{
			Deps:                `{ "dependencies": { "axios": "^0.19.0", "slackify-markdown": "^1.1.1"}}`,
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
					"function":   name[7:] + "-lambda",
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
	}
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

async function createPayload(githubPayload) {
    let labels = getLabels(githubPayload.issue.labels)
    let sentiment = await checkSentiment(githubPayload.issue.body, githubPayload.issue.title)
    if (sentiment)
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

function getLabels(labelsArray) {
  let labels = []
  labelsArray.map(label => labels.push(label.name))
  return labels
}


async function checkSentiment(issueBody, issueTitle) {
  let result = await axios.post(process.env.textAnalyticsEndpoint + 'text/analytics/v2.1/sentiment',
  {documents: [{id: '1', text: issueBody}, {id: '2', text: issueTitle}]}, {headers: {...{'Ocp-Apim-Subscription-Key': process.env.textAnalyticsKey}}})
  return ((result.data.documents[0].score > 0.5) && (result.data.documents[1].score > 0.5))
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

async function setLabel(url, msg) {
    const config = {
        headers: {
            "Content-Type": "application/json;charset=UTF-8"
        }
    };
    let sendMsg = await axios.patch(url, msg, config);
    return sendMsg;
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
        text: 'User *'+payload.issue.user.login+'* created an issue that might need a review: <'+payload.issue.html_url+'|*#'+payload.issue.number+' '+payload.issue.title+'*>'
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
}`
