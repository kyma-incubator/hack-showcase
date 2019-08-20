package registration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	log "github.com/sirupsen/logrus"
)

const (
	retryDelay                = 5 * time.Second
	retriesCount              = 10
	applicationRegistryPrefix = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/"
	applicationRegistrySuffix = "-app/v1/metadata/services"
)

//ServiceRegister is an interface containing all necessary functions required to register a service in Kyma Application Registry
type ServiceRegister interface {
	RegisterService() (string, error)
}

type serviceRegister struct {
	envName string
	builder Builder
}

//NewServiceRegister creates a serviceRegister instance with the passed in interface
func NewServiceRegister(deploymentEnvName string, b Builder) serviceRegister {
	return serviceRegister{envName: deploymentEnvName, builder: b}
}

//RegisterService - register service in Kyma and get a response
func (r serviceRegister) RegisterService() (string, apperrors.AppError) {

	jsonBody, err := r.builder.BuildServiceDetails()
	if err != nil {
		return "", apperrors.Internal("While building service details json: %s", err)
	}

	id, err := jsonBody.requestWithRetries(r.envName)
	if err != nil {
		return "", apperrors.Internal("While trying to register service: %s", err.Error())
	}
	return id, nil
}

func (jsonBody *ServiceDetails) requestWithRetries(appName string) (string, error) {
	var id string
	var err error
	var applicationRegistryURL = applicationRegistryPrefix + os.Getenv(appName) + applicationRegistrySuffix
	for i := 0; i < retriesCount; i++ {
		time.Sleep(retryDelay)
		id, err = sendRequest(*jsonBody, applicationRegistryURL)
		if err == nil {
			break
		}
		log.Warn(err.Error())
	}
	if err != nil {
		return "", apperrors.UpstreamServerCallFailed("While sending service registration request: %s", err)
	}
	return id, nil
}

//RegisterResponse contain structure of response json
type RegisterResponse struct {
	ID string
}

//Do - create request and send it to kyma application registry
func sendRequest(JSONBody ServiceDetails, url string) (string, error) {

	// parse json to io.Reader
	requestByte, err := json.Marshal(JSONBody)
	if err != nil {
		return "", apperrors.Internal("Failed to parse application registry request JSON body: %s", err.Error())
	}

	requestReader := bytes.NewReader(requestByte)

	httpRequest, err := http.NewRequest(http.MethodPost, url, requestReader)

	if err != nil {
		return "", apperrors.Internal("Failed to create JSON request: %s", err.Error())
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return "", apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", url, err.Error())
	}

	if httpResponse.StatusCode != http.StatusOK {
		return "", apperrors.UpstreamServerCallFailed("Incorrect response code '%d' while sending JSON request from %s", httpResponse.StatusCode, url)
	}

	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		return "", apperrors.UpstreamServerCallFailed("Failed to read service ID from application registry JSON response: %s", err)
	}

	var jsonResponse RegisterResponse
	err = json.Unmarshal(bodyBytes, &jsonResponse)
	if err != nil {
		return "", apperrors.Internal("Failed while unmarshaling JSON response from application registry: %s", err)
	}
	return jsonResponse.ID, nil
}
