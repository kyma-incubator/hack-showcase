package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KarolJaksik/hack-showcase/github-connector/mocks"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type payload struct {
	Anything string `json:"anything"`
}

func TestWebhookHandler_TestNoSecret(t *testing.T) {
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
		wh := NewWebhookHandler(WebhookStructHelper{})

		handler := http.HandlerFunc(wh.handleWebhook)
		handler.ServeHTTP(rr, req)
		// then
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}

func TestWebhookHandler_TestWrongPayload(t *testing.T) {
	t.Run("should respond with 401 status code", func(t *testing.T) {
		// given
		pld := payload{Anything: "test"}
		toSend, err := json.Marshal(pld)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Hub-Signature", "test")

		rr := httptest.NewRecorder()

		mockHandler := &mocks.Validator{}
		mockPayload, err := json.Marshal(payload{Anything: "test"})
		//require.NoError(t, err)
		mockHandler.On("GetToken").Return("test")
		mockHandler.On("ValidatePayload", req, []byte("test")).Return(mockPayload, nil)
		mockHandler.On("ParseWebHook", "", mockPayload).Return(errors.New("failed"))

		wh := NewWebhookHandler(mockHandler)

		// when
		handler := http.HandlerFunc(wh.handleWebhook)
		handler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})

}
