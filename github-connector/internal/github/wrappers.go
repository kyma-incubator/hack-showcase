package github

import (
	"net/http"

	"github.com/google/go-github/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

//ReceivingEventsWrapper that bundles the github library functions into one struct with a Validator interface
type ReceivingEventsWrapper struct {
	secret string
}

func NewReceivingEventsWrapper(s string) ReceivingEventsWrapper {
	return ReceivingEventsWrapper{secret: s}
}

//ValidatePayload is a function used for checking whether the secret provided in the request is correct
func (wh ReceivingEventsWrapper) ValidatePayload(r *http.Request, b []byte) ([]byte, apperrors.AppError) {
	payload, err := github.ValidatePayload(r, b)
	if err != nil {
		return nil, apperrors.AuthenticationFailed("Authentication during GitHub payload validation failed: %s", err)
	}
	return payload, nil
}

//ParseWebHook parses the raw json payload into an event struct
func (wh ReceivingEventsWrapper) ParseWebHook(s string, b []byte) (interface{}, apperrors.AppError) {
	webhook, err := github.ParseWebHook(s, b)
	if err != nil {
		return nil, apperrors.WrongInput("Failed to parse incomming github payload into struct: %s", err)
	}
	return webhook, nil
}

//GetToken is a function that looks for the secret in the environment
func (wh ReceivingEventsWrapper) GetToken() string {
	return wh.secret
}
