package registration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"

	"github.com/stretchr/testify/assert"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration/mocks"
)

const expectedID = "123-456-789"

func exampleServiceID(w http.ResponseWriter, r *http.Request) {
	id := registration.RegisterResponse{ID: expectedID}
	res, err := json.Marshal(id)
	if err != nil {
	}
	w.Write(res)
}

func TestRegisterService(t *testing.T) {
	t.Run("should return service ID", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleServiceID)
		server := httptest.NewServer(handler)
		defer server.Close()

		mockBuilder := &mocks.Builder{}
		mockBuilder.On("BuildServiceDetails").Return(registration.ServiceDetails{}, nil)
		mockBuilder.On("GetApplicationRegistryURL").Return(server.URL)

		service := registration.NewServiceRegister("deploymentEnvName", mockBuilder)

		//when
		id, err := service.RegisterService()

		//then
		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)

	})

	t.Run("should return an error", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleServiceID)
		server := httptest.NewServer(handler)
		defer server.Close()

		mockBuilder := &mocks.Builder{}
		mockBuilder.On("BuildServiceDetails").Return(registration.ServiceDetails{}, apperrors.Internal("error"))
		mockBuilder.On("GetApplicationRegistryURL").Return(server.URL)

		service := registration.NewServiceRegister("deploymentEnvName", mockBuilder)

		//when
		id, err := service.RegisterService()

		//then
		assert.Error(t, err)
		assert.Equal(t, "", id)

	})
}

// func Test(t *testing.T) {
// 	t.Run("should return an error when server is not responding", func(t *testing.T) {
// 		//given
// 		jsonBody := ServiceDetails{}
// 		sender := NewRequestSender()

// 		//when
// 		res, err := sender.Do(jsonBody, "example.com")

// 		//then
// 		assert.Error(t, err)
// 		assert.Equal(t, "", res)
// 	})

// 	t.Run("should return service ID", func(t *testing.T) {
// 		//given
// 		jsonBody := ServiceDetails{}
// 		sender := NewRequestSender()
// 		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			id := RegisterResponse{ID: "123-456-789"}
// 			res, err := json.Marshal(id)
// 			if err != nil {
// 			}
// 			w.Write(res)
// 		})
// 		server := httptest.NewServer(handler)
// 		defer server.Close()

// 		//when
// 		res, err := sender.Do(jsonBody, server.URL)

// 		//then
// 		assert.NoError(t, err)
// 		assert.Equal(t, "123-456-789", res)
// 	})

// 	t.Run("should return an error when server response with status code other than 200", func(t *testing.T) {
// 		//given
// 		jsonBody := ServiceDetails{}
// 		sender := NewRequestSender()
// 		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.WriteHeader(404)
// 		})
// 		server := httptest.NewServer(handler)
// 		defer server.Close()

// 		//when
// 		res, err := sender.Do(jsonBody, server.URL)

// 		//then
// 		assert.Error(t, err)
// 		assert.Equal(t, "", res)
// 	})
// }

// func TestNewRegisterRequestSender(t *testing.T) {
// 	t.Run("should return requestSender struct", func(t *testing.T) {
// 		//when
// 		sender := NewRequestSender()

// 		//then
// 		assert.Equal(t, requestSender{}, sender)
// 	})
// }

// const (
// 	exampleID = "123-456789-abcdefghi"
// )

// type TestServiceDetails struct {
// 	Name string
// }

// func TestCreateJSONRequest(t *testing.T) {
// 	t.Run("should respond with the same json properties (body, url, method)", func(t *testing.T) {
// 		given
// 		JSONBody := TestServiceDetails{
// 			Name: "kyma",
// 		}
// 		requestByte, err := json.Marshal(JSONBody)
// 		if err != nil {
// 			panic(err.Error)
// 		}
// 		requestReader := bytes.NewReader(requestByte)
// 		config := RequestConfig{
// 			Type: "POST",
// 			URL:  "http://www.hello-test.com",
// 			Body: requestReader,
// 		}

// 		when
// 		req, err := CreateJSONRequest(config)
// 		buf := new(bytes.Buffer)
// 		buf.ReadFrom(req.Body)
// 		s := buf.String()

// 		then
// 		assert.NoError(t, err)
// 		assert.Equal(t, s, string(requestByte))
// 		assert.Equal(t, req.URL.String(), config.URL)
// 		assert.Equal(t, req.Method, config.Type)
// 	})
// 	t.Run("should return an error when creating a header fails", func(t *testing.T) {
// 		given
// 		config := RequestConfig{URL: ":foo"}

// 		when
// 		resp, err := CreateJSONRequest(config)

// 		then
// 		assert.Error(t, err)
// 		assert.Nil(t, resp)
// 	})
// }

// func StatusBadRequestResponse(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusBadRequest)

// 	json.NewEncoder(w).Encode(RegisterResponse{
// 		ID: exampleID,
// 	})
// }

// func StatusOKResponse(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)

// 	json.NewEncoder(w).Encode(RegisterResponse{
// 		ID: exampleID,
// 	})
// }
// func TestSendJSONRequest_TestDataOK(t *testing.T) {
// 	t.Run("should response with StatusOK code", func(t *testing.T) {
// 		given
// 		handler := http.HandlerFunc(StatusOKResponse)
// 		server := httptest.NewServer(handler)
// 		defer server.Close()
// 		req, errNewRequest := http.NewRequest("POST", server.URL, nil)
// 		client := server.Client()
// 		config := RegisterConfig{
// 			HTTPClient:  client,
// 			HTTPRequest: req,
// 		}

// 		when
// 		res, errSendJSON := SendJSONRequest(config)

// 		then
// 		assert.Equal(t, res.StatusCode, http.StatusOK)
// 		assert.NoError(t, errSendJSON)
// 		assert.NoError(t, errNewRequest)
// 	})
// 	t.Run("should return an error when server responses with code other than 200", func(t *testing.T) {
// 		given
// 		handler := http.HandlerFunc(StatusBadRequestResponse)
// 		server := httptest.NewServer(handler)
// 		defer server.Close()
// 		req, errNewRequest := http.NewRequest("POST", server.URL, nil)
// 		client := server.Client()
// 		config := RegisterConfig{
// 			HTTPClient:  client,
// 			HTTPRequest: req,
// 		}

// 		when
// 		res, err := SendJSONRequest(config)

// 		then
// 		assert.Error(t, err)
// 		assert.NoError(t, errNewRequest)
// 		assert.Nil(t, res)
// 	})
// }

// func TestRegisterApp(t *testing.T) {
// 	t.Run("should response exampleID", func(t *testing.T) {
// 		given
// 		JSONBody := ServiceDetails{
// 			Provider:    "kyma",
// 			Name:        "github-connector",
// 			Description: "Boilerplate for GitHub connector",
// 			API: &API{
// 				TargetURL: "https://console.35.195.62.81.xip.io/github-api",
// 			},
// 		}
// 		handler := http.HandlerFunc(StatusOKResponse)
// 		server := httptest.NewServer(handler)
// 		defer server.Close()

// 		when
// 		res, err := SendRegisterRequest(JSONBody, server.URL)
// 		fmt.Println(res)
// 		then

// 		assert.NoError(t, err)
// 		assert.Equal(t, exampleID, res)
// 	})
// }
