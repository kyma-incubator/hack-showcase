package registerapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/model"
	"github.com/stretchr/testify/assert"
)

const (
	exampleID = "123-456789-abcdefghi"
)

type TestServiceDetails struct {
	Name string
}

func TestCreateJSONRequest(t *testing.T) {
	t.Run("should respond with the same json properties (body, url, method)", func(t *testing.T) {
		//given
		JSONBody := TestServiceDetails{
			Name: "kyma",
		}
		requestByte, err := json.Marshal(JSONBody)
		if err != nil {
			panic(err.Error)
		}
		requestReader := bytes.NewReader(requestByte)
		config := RequestConfig{
			Type: "POST",
			URL:  "http://www.hello-test.com",
			Body: requestReader,
		}

		//when
		req, err := CreateJSONRequest(config)
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		s := buf.String()

		//then
		assert.NoError(t, err)
		assert.Equal(t, s, string(requestByte))
		assert.Equal(t, req.URL.String(), config.URL)
		assert.Equal(t, req.Method, config.Type)
	})
}

func StatusOKResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(RegisterResponse{
		ID: exampleID,
	})
}
func TestSendJSONRequest_TestDataOK(t *testing.T) {
	t.Run("should response with StatusOK code", func(t *testing.T) {
		// given

		handler := http.HandlerFunc(StatusOKResponse)
		server := httptest.NewServer(handler)
		defer server.Close()
		req, errNewRequest := http.NewRequest("POST", server.URL, nil)
		client := server.Client()
		config := RegisterConfig{
			HTTPClient:  client,
			HTTPRequest: req,
		}
		//when
		res, errSendJSON := SendJSONRequest(config)
		//then
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.NoError(t, errSendJSON)
		assert.NoError(t, errNewRequest)
	})
}

func TestRegisterApp(t *testing.T) {
	t.Run("should response exampleID", func(t *testing.T) {
		//given
		JSONBody := model.ServiceDetails{
			Provider:    "kyma",
			Name:        "github-connector",
			Description: "Boilerplate for GitHub connector",
			API: &model.API{
				TargetURL: "https://console.35.195.62.81.xip.io/github-api",
			},
		}
		handler := http.HandlerFunc(StatusOKResponse)
		server := httptest.NewServer(handler)
		defer server.Close()

		//when
		res, err := RegisterApp(JSONBody, server.URL)
		fmt.Println(res)
		//then

		assert.NoError(t, err)
		assert.Equal(t, exampleID, res)

	})
}
