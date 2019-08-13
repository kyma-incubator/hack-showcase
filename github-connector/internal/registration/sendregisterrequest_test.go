package registration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func exampleServiceID(w http.ResponseWriter, r *http.Request) {
	id := RegisterResponse{ID: "123-456-789"}
	res, err := json.Marshal(id)
	if err != nil {
	}
	w.Write(res)
}

func TestDo(t *testing.T) {
	t.Run("should return an error when server is not responding", func(t *testing.T) {
		//given
		jsonBody := ServiceDetails{}
		sender := NewRegisterRequestSender()

		//when
		res, err := sender.Do(jsonBody, "example.com")

		//then
		assert.Error(t, err)
		assert.Equal(t, "", res)
	})

	t.Run("should return service ID", func(t *testing.T) {
		//given
		jsonBody := ServiceDetails{}
		sender := NewRegisterRequestSender()
		handler := http.HandlerFunc(exampleServiceID)
		server := httptest.NewServer(handler)
		defer server.Close()

		//when
		res, err := sender.Do(jsonBody, server.URL)

		//then
		assert.NoError(t, err)
		assert.Equal(t, "123-456-789", res)
	})

	t.Run("should return an error when server response with status code other than 200", func(t *testing.T) {
		//given
		jsonBody := ServiceDetails{}
		sender := NewRegisterRequestSender()
		handler := http.HandlerFunc(exampleServiceID)
		server := httptest.NewServer(handler)
		defer server.Close()

		//when

		//then

	})
}

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
