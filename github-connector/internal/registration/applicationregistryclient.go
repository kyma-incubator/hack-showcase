package registration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	log "github.com/sirupsen/logrus"
)

//PayloadBuilder is an interface containing all necessary functions required to build an ServiceDetails structure
type PayloadBuilder interface {
	Build() (ServiceDetails, error)
	GetApplicationRegistryURL() string
}

//ApplicationRegistryClient is an interface containing all necessary functions required to register service in Kyma's application registry
type ApplicationRegistryClient interface {
	RegisterService() (string, apperrors.AppError)
}

type applicationRegistryClient struct {
	envName      string
	builder      PayloadBuilder
	register     ApplicationRegistryClient
	retryDelay   int
	retriesCount int
}

//NewApplicationRegistryClient creates a applicationRegistryClient instance with the passed in interface
func NewApplicationRegistryClient(deploymentEnvName string, b PayloadBuilder, retryTime int, retries int) applicationRegistryClient {

	return applicationRegistryClient{
		envName:      deploymentEnvName,
		builder:      b,
		retryDelay:   retryTime * int(time.Second),
		retriesCount: retries,
	}
}

//RegisterService - register service in Kyma and get a response
func (r applicationRegistryClient) RegisterService() (string, apperrors.AppError) {

	jsonBody, err := r.builder.Build()
	if err != nil {
		return "", apperrors.Internal("While building service details json: %s", err)
	}
	id, err := jsonBody.requestWithRetries(r.envName, r.builder.GetApplicationRegistryURL(), r.retryDelay, r.retriesCount)
	if err != nil {
		return "", apperrors.Internal("While trying to register service: %s", err.Error())
	}
	return id, nil
}

func (jsonBody *ServiceDetails) requestWithRetries(appName string, url string, retryDelay int, retriesCount int) (string, error) {
	var id string
	var err error

	for i := 0; i < retriesCount; i++ {
		time.Sleep(time.Duration(retryDelay))
		id, err = sendRequest(*jsonBody, url)
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
		return "", apperrors.Internal("Failed while unmarshalling JSON response from application registry: %s", err)
	}
	return jsonResponse.ID, nil
}
