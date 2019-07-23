package registerservice

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

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
func CreateJSONRequest(config RequestConfig) (*http.Request, error) {
	req, err := http.NewRequest(config.Type, config.URL, config.Body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

//SendJSONRequest - create json struct and try post it into application-register's url
func SendJSONRequest(config RegisterConfig) (*http.Response, error) {

	resp, err := config.HTTPClient.Do(config.HTTPRequest)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

//SendRegisterRequest - create request and send it to kyma application registry
func SendRegisterRequest(JSONBody model.ServiceDetails, url string) (string, error) {

	// parse json to io.Reader
	requestByte, err := json.Marshal(JSONBody)
	if err != nil {
		return "", err
	}

	requestReader := bytes.NewReader(requestByte)

	// create POST request
	requestConfig := RequestConfig{
		Type: "POST",
		URL:  url,
		Body: requestReader,
	}

	httpRequest, err := CreateJSONRequest(requestConfig)
	if err != nil {
		return "", err
	}

	// create register config
	config := RegisterConfig{
		HTTPClient:  &http.Client{},
		HTTPRequest: httpRequest,
	}

	// happy JSONRequestSend-ing!
	httpResponse, err := SendJSONRequest(config)

	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)

	var jsonResponse RegisterResponse
	json.Unmarshal(bodyBytes, &jsonResponse)
	return jsonResponse.ID, nil
}
