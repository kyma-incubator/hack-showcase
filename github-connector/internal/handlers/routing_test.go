package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KarolJaksik/hack-showcase/github-connector/internal/handlers/mocks"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type toJSON struct {
	TestJSON string `json:TestJSON"`
}

//createRequest creates an HTTP request for test purposes
func createRequest(t *testing.T) *http.Request {

	payload := toJSON{TestJSON: "test"}
	toSend, err := json.Marshal(payload)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hub-Signature", "test")

	return req
}

func TestWebhookHandler_TestBadSecret(t *testing.T) {
	t.Run("should respond with 401 status code", func(t *testing.T) {
		// given

		payload := toJSON{TestJSON: "test"}
		toSend, err := json.Marshal(payload)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		mockHandler := &mocks.Validator{}

		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(nil, errors.New("failed"))

		// when
		wh := NewWebHookHandler(mockHandler)

		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)
		// then
		mockHandler.AssertExpectations(t)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}

func TestWebhookHandler_TestWrongPayload(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {
		// given
		req := createRequest(t)
		rr := httptest.NewRecorder()

		mockHandler := &mocks.Validator{}
		mockPayload, err := json.Marshal(toJSON{TestJSON: "test"})
		require.NoError(t, err)

		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockHandler.On("ParseWebHook", "", mockPayload).Return(nil, errors.New("failed"))

		wh := NewWebHookHandler(mockHandler)
		// when
		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockHandler.AssertExpectations(t)
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}

func TestWebhookHandler_TestKnownEvent(t *testing.T) {
	t.Run("should respond with 200 status code", func(t *testing.T) {
		// given

		req := createRequest(t)
		rr := httptest.NewRecorder()

		mockHandler := &mocks.Validator{}
		mockPayload, err := json.Marshal(toJSON{TestJSON: "test"})
		require.NoError(t, err)

		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		event := &github.StarEvent{}
		mockHandler.On("ParseWebHook", "", mockPayload).Return(event, nil)

		wh := NewWebHookHandler(mockHandler)
		// when
		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockHandler.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, rr.Code)

	})

}

func TestWebhookHandler_TestUnknownEvent(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {
		// given
		req := createRequest(t)
		rr := httptest.NewRecorder()

		mockHandler := &mocks.Validator{}
		mockPayload, err := json.Marshal(toJSON{TestJSON: "test"})
		require.NoError(t, err)
		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockHandler.On("ParseWebHook", "", mockPayload).Return(1, nil)

		wh := NewWebHookHandler(mockHandler)
		// when
		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockHandler.AssertExpectations(t)
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
