package slack

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kyma-incubator/github-slack-connectors/slack-connector/internal/apperrors"
	"github.com/nlopes/slack/slackevents"
)

//ReceivingEventsWrapper that bundles the github library functions into one struct with a Validator interface
type receivingEventsWrapper struct {
	secret string
}

//NewReceivingEventsWrapper return receivingEventsWrapper struct
func NewReceivingEventsWrapper(s string) Validator {
	return &receivingEventsWrapper{secret: s}
}

//Validator is an interface providing wrapper methods for external library
type Validator interface {
	ValidatePayload(*http.Request, []byte) ([]byte, apperrors.AppError)
	ParseWebHook([]byte) (interface{}, apperrors.AppError)
	GetToken() string
}

//ValidatePayload is a function used for checking whether the secret provided in the request is correct
func (wh receivingEventsWrapper) ValidatePayload(r *http.Request, secret []byte) ([]byte, apperrors.AppError) {
	timestamp := r.Header.Get("X-Slack-Request-Timestamp")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, apperrors.Internal("While reading Slack request body: %s", err)
	}
	basestring := fmt.Sprintf("v0:%s:%s", timestamp, body)

	h := hmac.New(sha256.New, []byte(secret))

	_, err = h.Write([]byte(basestring))
	if err != nil {
		return []byte{}, apperrors.Internal("Failed to validate payload: %s", err)
	}

	sha := "v0=" + hex.EncodeToString(h.Sum(nil))
	slackSign := r.Header.Get("X-Slack-Signature")

	if slackSign != sha {
		return []byte{}, apperrors.AuthenticationFailed("Failed to validate signature: Signature is different than expected")
	}

	return body, nil
}

//ParseWebHook parses the raw json payload into an event struct
func (wh receivingEventsWrapper) ParseWebHook(b []byte) (interface{}, apperrors.AppError) {
	webhook, err := slackevents.ParseEvent(b, slackevents.OptionNoVerifyToken())
	if err != nil {
		return slackevents.EventsAPIEvent{}, apperrors.WrongInput("Failed to parse incoming slack payload into struct: %s", err)
	}
	return webhook, nil
}

//GetToken is a function that looks for the secret in the environment
func (wh receivingEventsWrapper) GetToken() string {
	return wh.secret
}
