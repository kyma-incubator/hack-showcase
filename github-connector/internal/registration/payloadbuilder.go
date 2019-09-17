package registration

import (
	"fmt"
	"io/ioutil"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

const (
	specificationURL          = "https://raw.githubusercontent.com/kyma-incubator/hack-showcase/master/github-connector/internal/registration/configs/githubopenAPI.json"
	applicationRegistryPrefix = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/"
	applicationRegistrySuffix = "-app/v1/metadata/services"
	applicationRegistryFormat = "%s%s%s"
)

//FileReader is an interface used to allow mocking file reading
type FileReader interface {
	Read(string) ([]byte, error)
}

type payloadBuilder struct {
	builder         PayloadBuilder
	fileReader      FileReader
	applicationName string
	githubToken     string
	receiveEvents   bool
	sendEvents      bool
}

//NewPayloadBuilder creates a serviceDetailsPayloadBuilder instance
func NewPayloadBuilder(fr FileReader, appName string, token string, receive bool, send bool) payloadBuilder {
	return payloadBuilder{
		fileReader:      fr,
		applicationName: appName,
		githubToken:     token,
		receiveEvents:   receive,
		sendEvents:      send,
	}
}

//Build creates a ServiceDetails structure with provided API specification URL
func (r payloadBuilder) Build() (ServiceDetails, error) {

	jsonBody := ServiceDetails{
		Provider:    "Kyma",
		Name:        r.applicationName,
		Description: "GitHub Connector, which can be used for communication and handling events from GitHub",
		API: &API{
			TargetURL:         "https://api.github.com",
			RequestParameters: &RequestParameters{Headers: &Headers{CustomHeader: []string{}}},
		},
	}
	if r.receiveEvents {
		file, err := r.fileReader.Read("githubasyncapi.json")
		if err != nil {
			return ServiceDetails{}, apperrors.Internal("While reading githubasyncapi.json: %s", err)
		}
		jsonBody.Events = &Events{Spec: file}
	}

	if r.sendEvents {
		jsonBody.API.SpecificationURL = specificationURL
		jsonBody.API.RequestParameters.Headers.CustomHeader = []string{"token " + r.githubToken}
	}
	return jsonBody, nil
}

//GetApplicationRegistryURL returns a URL used to POST json to Kyma's application registry
func (r payloadBuilder) GetApplicationRegistryURL() string {
	return fmt.Sprintf(applicationRegistryFormat, applicationRegistryPrefix, r.applicationName, applicationRegistrySuffix)
}

//Read reads file specified with given path using ioutil library
func (r fileReader) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

type fileReader struct {
	fileReader FileReader
}

//NewFileReader creates new fileReader struct
func NewFileReader() fileReader {
	return fileReader{}
}
