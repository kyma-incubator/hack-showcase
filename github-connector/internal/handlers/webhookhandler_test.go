package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers/mocks"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	t.Run("should respond with 403 status code", func(t *testing.T) {
		// given

		payload := toJSON{TestJSON: "test"}
		toSend, err := json.Marshal(payload)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		mockValidator := &mocks.Validator{}
		mockSender := &mocks.Sender{}

		mockValidator.On("GetToken").Return("test")
		mockValidator.On("ValidatePayload", req, []byte("test")).Return(nil, apperrors.AuthenticationFailed("fail"))

		// when
		wh := NewWebHookHandler(mockValidator, mockSender)

		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockValidator.AssertExpectations(t)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}

func TestWebhookHandler_TestWrongPayload(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {

		// given
		req := createRequest(t)
		rr := httptest.NewRecorder()

		mockValidator := &mocks.Validator{}
		mockSender := &mocks.Sender{}
		mockPayload, err := json.Marshal(toJSON{TestJSON: "test"})
		require.NoError(t, err)

		mockValidator.On("GetToken").Return("test")
		mockValidator.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockValidator.On("ParseWebHook", "", mockPayload).Return(nil, apperrors.WrongInput("fail"))

		wh := NewWebHookHandler(mockValidator, mockSender)

		// when
		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockValidator.AssertExpectations(t)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

}

func TestWebhookHandler_TestKnownEvent(t *testing.T) {
	t.Run("should respond with 200 status code", func(t *testing.T) {

		// given
		req := createRequest(t)
		rr := httptest.NewRecorder()

		mockValidator := &mocks.Validator{}
		mockSender := &mocks.Sender{}
		mockPayload, err := json.Marshal(toJSON{TestJSON: "test"})
		require.NoError(t, err)
		rawPayload := json.RawMessage(mockPayload)
		mockSender.On("SendToKyma", "issuesevent.opened", "v1", "", "github-connector-app", rawPayload).Return(nil)

		mockValidator.On("GetToken").Return("test")
		mockValidator.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		event := &github.IssuesEvent{}
		mockValidator.On("ParseWebHook", "", mockPayload).Return(event, nil)

		wh := NewWebHookHandler(mockValidator, mockSender)

		// when
		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockValidator.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestWebhookHandler_TestUnknownEvent(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {

		// given
		req := createRequest(t)
		rr := httptest.NewRecorder()

		mockValidator := &mocks.Validator{}
		mockSender := &mocks.Sender{}

		mockPayload, err := json.Marshal(toJSON{TestJSON: "test"})
		require.NoError(t, err)
		mockValidator.On("GetToken").Return("test")
		mockValidator.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockValidator.On("ParseWebHook", "", mockPayload).Return(1, nil)

		wh := NewWebHookHandler(mockValidator, mockSender)

		// when
		handler := http.HandlerFunc(wh.HandleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		mockValidator.AssertExpectations(t)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
