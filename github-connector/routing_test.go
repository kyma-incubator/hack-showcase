package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KarolJaksik/hack-showcase/github-connector/mocks"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type payload struct {
	Anything string `json:"anything"`
}

//createRequest creates an HTTP request for test purposes
func createRequest(t *testing.T) (*httptest.ResponseRecorder, *http.Request, *mocks.Validator, []byte) {

	pld := payload{Anything: "test"}
	toSend, err := json.Marshal(pld)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hub-Signature", "test")

	rr := httptest.NewRecorder()

	mockHandler := &mocks.Validator{}
	mockPayload, err := json.Marshal(payload{Anything: "test"})
	require.NoError(t, err)
	return rr, req, mockHandler, mockPayload
}

func TestWebhookHandler_TestBadSecret(t *testing.T) {
	t.Run("should respond with 401 status code", func(t *testing.T) {
		// given

		pld := payload{Anything: "test"}
		toSend, err := json.Marshal(pld)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Hub-Signature", "sha1=dcdf499d859b235063659adceea6ef474cb23a51")

		rr := httptest.NewRecorder()

		// when
		wh := NewWebHookHandler(WebHookStruct{})

		handler := http.HandlerFunc(wh.handleWebhook)
		handler.ServeHTTP(rr, req)
		// then
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}

func TestWebhookHandler_TestWrongPayload(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {
		// given
		rr, req, mockHandler, mockPayload := createRequest(t)

		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockHandler.On("ParseWebHook", "", mockPayload).Return(nil, errors.New("failed"))

		wh := NewWebHookHandler(mockHandler)
		// when
		handler := http.HandlerFunc(wh.handleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestWebhookHandler_TestKnownEvent(t *testing.T) {
	t.Run("should respond with 200 status code", func(t *testing.T) {
		// given

		rr, req, mockHandler, mockPayload := createRequest(t)

		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		event := &github.StarEvent{}
		mockHandler.On("ParseWebHook", "", mockPayload).Return(event, nil)

		wh := NewWebHookHandler(mockHandler)
		// when
		handler := http.HandlerFunc(wh.handleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusOK, rr.Code)

	})

}

func TestWebhookHandler_TestUnknownEvent(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {
		// given
		rr, req, mockHandler, mockPayload := createRequest(t)

		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockHandler.On("ParseWebHook", "", mockPayload).Return(1, nil)

		wh := NewWebHookHandler(mockHandler)
		// when
		handler := http.HandlerFunc(wh.handleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
