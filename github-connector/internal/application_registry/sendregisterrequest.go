package registerservice

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/model"
)

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

//CreateJSONRequest - create http request, add headers and return client
func CreateJSONRequest(config RequestConfig) (*http.Request, apperrors.AppError) {
	if config.Type != "POST" {
		return nil, apperrors.Internal("Wrong http request method")
	}

	req, err := http.NewRequest(config.Type, config.URL, config.Body)

	if err != nil {
		return nil, apperrors.Internal("Failed to create JSON request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

//SendJSONRequest - create json struct and try post it into application-register's url
func SendJSONRequest(config RegisterConfig) (*http.Response, apperrors.AppError) {

	resp, err := config.HTTPClient.Do(config.HTTPRequest)

	if err != nil {
		return nil, apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", config.HTTPRequest.URL, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, apperrors.UpstreamServerCallFailed("Incorrect response code '%d' while sending JSON request from %s", resp.StatusCode, config.HTTPRequest.URL)
	}
	return resp, nil
}

//SendRegisterRequest - create request and send it to kyma application registry
func SendRegisterRequest(JSONBody model.ServiceDetails, url string) (string, error) {

	// parse json to io.Reader
	requestByte, err := json.Marshal(JSONBody)
	if err != nil {
		return "", apperrors.Internal("Failed to parse application registry request JSON body: %s", err.Error())
	}

	requestReader := bytes.NewReader(requestByte)

	// create POST request
	requestConfig := RequestConfig{
		Type: "POST",
		URL:  url,
		Body: requestReader,
	}

	httpRequest, apperr := CreateJSONRequest(requestConfig)
	if apperr != nil {
		return "", apperr.Append("While preparing application registry JSON request")
	}

	// create register config
	config := RegisterConfig{
		HTTPClient:  &http.Client{},
		HTTPRequest: httpRequest,
	}

	// happy JSONRequestSend-ing!
	httpResponse, apperr := SendJSONRequest(config)

	if apperr != nil {
		return "", apperr.Append("While sending application registry JSON request")
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
