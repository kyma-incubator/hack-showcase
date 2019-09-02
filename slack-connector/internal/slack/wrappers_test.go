package slack

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePayload(t *testing.T) {
	t.Run("Should respond with payload body when signature is proper", func(t *testing.T) {
		//given
		toSend := "example payload"
		body := bytes.NewReader([]byte(toSend))
		request, err := http.NewRequest(http.MethodPost, "http://slack.com", body)
		if err != nil {

		}
		request.Header.Add("X-Slack-Request-Timestamp", "123")
		request.Header.Add("X-Slack-Signature", "v0=e86821bcdfd89b7f1bbaf2ed912095cb340c6dd8d21ca6a4a40c7d4a08d39620")
		secret := "secret"
		wh := NewReceivingEventsWrapper(secret)
		//when
		res, err := wh.ValidatePayload(request, []byte(secret))

		//then
		assert.Equal(t, []byte(toSend), res)
		assert.NoError(t, err)
	})
	t.Run("Should respond with an error body when signature is wrong", func(t *testing.T) {
		//given
		toSend := "example payload"
		body := bytes.NewReader([]byte(toSend))
		request, err := http.NewRequest(http.MethodPost, "http://slack.com", body)
		if err != nil {

		}
		request.Header.Add("X-Slack-Request-Timestamp", "123")
		request.Header.Add("X-Slack-Signature", "v0=e86821bcdfd89b7f1bbaf2ed912095cb340c6dd8d21ca6a4a40c7d4a08d39620")
		secret := "wrong-secret"
		wh := NewReceivingEventsWrapper(secret)
		//when
		res, err := wh.ValidatePayload(request, []byte(secret))

		//then
		assert.Equal(t, []byte{}, res)
		assert.Error(t, err)
	})
}
