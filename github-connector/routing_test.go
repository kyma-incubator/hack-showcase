package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type payload struct {
	cokolwiek string `json:"hejka"`
}

func TestWebhookHandler_TestNoSecret(t *testing.T) {
	t.Run("should respond with 401 status code", func(t *testing.T) {
		// given
		pld := payload{cokolwiek: "filip"}
		toSend, err := json.Marshal(pld)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Hub-Signature", "sha1=dcdf499d859b235063659adceea6ef474cb23a51")

		rr := httptest.NewRecorder()

		// when
		handler := http.HandlerFunc(handleWebhook)

		handler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}

func TestWebhookHandler_TestWrongPayload(t *testing.T) {
	t.Run("should respond with 400 status code", func(t *testing.T) {
		// given
		pld := payload{cokolwiek: "filip"}
		toSend, err := json.Marshal(pld)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(toSend))
		req.Header.Set("Content-Type", "application/json")

		h := hmac.New(sha1.New, []byte("my-secret-key"))

		req.Header.Set("X-Hub-Signature", ("sha1=" + base64.StdEncoding.EncodeToString(h.Sum(nil))))
		//	req.SetBasicAuth("X-Hub-Signature", "sha1="+hex.EncodeToString(bs))
		rr := httptest.NewRecorder()

		// when
		handler := http.HandlerFunc(handleWebhook)

		handler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

	})
}
