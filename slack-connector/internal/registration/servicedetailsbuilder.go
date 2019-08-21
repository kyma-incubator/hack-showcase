package registration

import (
	"io/ioutil"
	"os"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"
)

const (
	specificationURL          = "https://raw.githubusercontent.com/kyma-incubator/hack-showcase/slack-connector-boilerplate/slack-connector/internal/registration/configs/slackopenapi.json"
	applicationRegistryPrefix = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/"
	applicationRegistrySuffix = "-app/v1/metadata/services"
	applicationName           = "SLACK_CONNECTOR_NAME"
)

//OSCommunicator is an interface used to allow mocking file reading
type OSCommunicator interface {
	ReadFile(string) ([]byte, error)
	GetEnv(string) string
}

//ServiceDetailsBuilder is used for mocking building ServiceDetails struct
type serviceDetailsBuilder struct {
	builder        Builder
	osCommunicator OSCommunicator
}

//NewServiceDetailsBuilder creates a serviceDetailsBuilder instance
func NewServiceDetailsBuilder(fr OSCommunicator) serviceDetailsBuilder {
	return serviceDetailsBuilder{osCommunicator: fr}
}

//BuildServiceDetails creates a ServiceDetails structure with provided API specification URL
func (r serviceDetailsBuilder) BuildServiceDetails() (ServiceDetails, error) {

	var jsonBody = ServiceDetails{
		Provider:    "Kyma",
		Name:        r.osCommunicator.GetEnv(applicationName),
		Description: "Slack Connector, which is used for registering Slack API in Kyma",
		API: &API{
			TargetURL:         "https://slack.com/api",
			RequestParameters: &RequestParameters{Headers: &Headers{CustomHeader: []string{"Bearer xoxb-725497410967-712376428434-PDNTPXtxTYW9i6v0PB78AkfS"}}},
		},
	}
	file, err := r.osCommunicator.ReadFile("slackasyncapi.json")
	if err != nil {
		return ServiceDetails{}, apperrors.Internal("While reading 'slackopenapi.json' spec: %s", err)
	}
	jsonBody.Events = &Events{Spec: file}

	jsonBody.API.SpecificationURL = specificationURL
	return jsonBody, nil
}

//GetApplicationRegistryURL returns a URL used to POST json to Kyma's application registry
func (r serviceDetailsBuilder) GetApplicationRegistryURL() string {
	return applicationRegistryPrefix + r.osCommunicator.GetEnv(applicationName) + applicationRegistrySuffix
}

//ReadFile reads file specified with given path using ioutil library
func (r osCommunicator) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

type osCommunicator struct {
	osCommunicator OSCommunicator
}

//NewOSCommunicator creates new osCommunicator struct
func NewOSCommunicator() osCommunicator {
	return osCommunicator{}
}

//GetEnv returns environmental variable of a given name
func (r osCommunicator) GetEnv(name string) string {
	return os.Getenv(name)
}
