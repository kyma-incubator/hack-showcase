package registration

import (
	"io/ioutil"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

const (
	specificationURL          = "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.json"
	applicationRegistryPrefix = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/"
	applicationRegistrySuffix = "-app/v1/metadata/services"
	applicationName           = "GITHUB_CONNECTOR_NAME"
)

//OSCommunicator is an interface used to allow mocking file reading
type OSCommunicator interface {
	ReadFile(string) ([]byte, error)
	GetEnv(string) string
}

type payloadBuilder struct {
	builder        PayloadBuilder
	osCommunicator OSCommunicator
}

//NewPayloadBuilder creates a serviceDetailsPayloadBuilder instance
func NewPayloadBuilder(fr OSCommunicator) payloadBuilder {
	return payloadBuilder{osCommunicator: fr}
}

//Build creates a ServiceDetails structure with provided API specification URL
func (r payloadBuilder) Build() (ServiceDetails, error) {

	var jsonBody = ServiceDetails{
		Provider:    "Kyma",
		Name:        r.osCommunicator.GetEnv(applicationName),
		Description: "GitHub Connector, which can be used for communication and handling events from GitHub",
		API: &API{
			TargetURL: "https://api.github.com",
		},
	}
	file, err := r.osCommunicator.ReadFile("githubasyncapi.json")
	if err != nil {
		return ServiceDetails{}, apperrors.Internal("While reading githubasyncapi.json: %s", err)
	}
	jsonBody.Events = &Events{Spec: file}

	jsonBody.API.SpecificationURL = specificationURL
	return jsonBody, nil
}

//GetApplicationRegistryURL returnes a URL used to POST json to Kyma's application registry
func (r payloadBuilder) GetApplicationRegistryURL() string {
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
