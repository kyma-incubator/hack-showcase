package registration

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"
)

const (
	specificationURL          = "https://raw.githubusercontent.com/kyma-incubator/hack-showcase/slack-connector-boilerplate/slack-connector/internal/registration/configs/slackopenapi.json" // to edit after merge
	applicationRegistryPrefix = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/"
	applicationRegistrySuffix = "-app/v1/metadata/services"
	applicationRegistryFormat = "%s%s%s"
)

//FileReader is an interface used to allow mocking file reading
type FileReader interface {
	Read(string) ([]byte, error)
}

//NewPayloadBuilder creates a serviceDetailsPayloadBuilder instance
type payloadBuilder struct {
	builder         PayloadBuilder
	fileReader      FileReader
	applicationName string
}

//NewPayloadBuilder creates a serviceDetailsPayloadBuilder instance
func NewPayloadBuilder(fr FileReader, appName string) payloadBuilder {
	return payloadBuilder{fileReader: fr, applicationName: appName}
}

//Build creates a ServiceDetails structure with provided API specification URL
func (r payloadBuilder) Build() (ServiceDetails, error) {

	var jsonBody = ServiceDetails{
		Provider:    "Kyma",
		Name:        r.applicationName,
		Description: "Slack Connector, which is used for registering Slack API in Kyma",
		API: &API{
			TargetURL:         "https://slack.com/api",
			RequestParameters: &RequestParameters{Headers: &Headers{CustomHeader: []string{os.Getenv("SLACK_BOT_TOKEN")}}},
		},
	}
	file, err := r.fileReader.Read("slackasyncapi.json")
	if err != nil {
		return ServiceDetails{}, apperrors.Internal("While reading 'slackopenapi.json' spec: %s", err)
	}
	jsonBody.Events = &Events{Spec: file}

	jsonBody.API.SpecificationURL = specificationURL
	return jsonBody, nil
}

//GetApplicationRegistryURL returns a URL used to POST json to Kyma's application registry
func (r payloadBuilder) GetApplicationRegistryURL() string {
	return fmt.Sprintf(applicationRegistryFormat, applicationRegistryPrefix, r.applicationName, applicationRegistrySuffix)
}

//ReadFile reads file specified with given path using ioutil library
func (r fileReader) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

type fileReader struct {
	fileReader FileReader
}

//NewFileReader creates new osCommunicator struct
func NewFileReader() fileReader {
	return fileReader{}
}
