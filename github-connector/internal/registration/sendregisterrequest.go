package registration

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

//RequestSender is an interface used to allow mocking sending service registration request
type RequestSender interface {
	Do(JSONBody ServiceDetails, url string) (string, error)
}

//requestSender is an struct used to allow mocking the service registration functions
type requestSender struct {
	sender RequestSender
}

//RegisterResponse contain structure of response json
type RegisterResponse struct {
	ID string
}

//RegisterConfig contain configs
type RegisterConfig struct {
	HTTPClient  *http.Client
	HTTPRequest *http.Request
}

//RequestConfig contain configs to create http requests
type RequestConfig struct {
	Type string
	URL  string
	Body io.Reader
}

//NewRegisterRequestSender creates a registerRequestSender instance with the passed in interface
func NewRegisterRequestSender() requestSender {
	return requestSender{}
}

//Do - create request and send it to kyma application registry
func (r requestSender) Do(JSONBody ServiceDetails, url string) (string, error) {

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
