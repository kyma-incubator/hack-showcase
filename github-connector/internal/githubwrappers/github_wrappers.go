package githubwrappers

import (
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

//ReceivingEventsWrapper that bundles the github library functions into one struct with a Validator interface
type ReceivingEventsWrapper struct {
}

//ValidatePayload is a function used for checking whether the secret provided in the request is correct
func (wh ReceivingEventsWrapper) ValidatePayload(r *http.Request, b []byte) ([]byte, error) {
	return github.ValidatePayload(r, b)
}

//ParseWebHook parses the raw json payload into an event struct
func (wh ReceivingEventsWrapper) ParseWebHook(s string, b []byte) (interface{}, error) {
	return github.ParseWebHook(s, b)
}

//GetToken is a function that looks for the secret in the environment
func (wh ReceivingEventsWrapper) GetToken() string {
	return os.Getenv("GITHUB-SECRET")
}
