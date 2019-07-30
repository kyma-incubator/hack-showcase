package githubwrappers

import (
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

//ReceivingEventsWrapper that bundles the github library functions into one struct with a Validator interface
type ReceivingEventsWrapper struct {
}

//ValidatePayload is a function used for checking whether the secret provided in the request is correct
func (wh ReceivingEventsWrapper) ValidatePayload(r *http.Request, b []byte) ([]byte, apperrors.AppError) {
	payload, err := github.ValidatePayload(r, b)
	return payload, apperrors.AuthenticationFailed("Authentication during GitHub payload validation failed: %s", err)
}

//ParseWebHook parses the raw json payload into an event struct
func (wh ReceivingEventsWrapper) ParseWebHook(s string, b []byte) (interface{}, apperrors.AppError) {
	webhook, err := github.ParseWebHook(s, b)
	return webhook, apperrors.Internal("Failed to parse incomming github payload into object: %s", err)
}

//GetToken is a function that looks for the secret in the environment
func (wh ReceivingEventsWrapper) GetToken() string {
	return os.Getenv("GITHUB-SECRET")
}
