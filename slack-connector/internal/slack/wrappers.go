package slack

import (
	"net/http"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"
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
func (wh receivingEventsWrapper) ValidatePayload(r *http.Request, b []byte) ([]byte, apperrors.AppError) {
	//TODO: https://api.slack.com/docs/verifying-requests-from-slack#about

	return []byte{}, nil
}

//ParseWebHook parses the raw json payload into an event struct
func (wh receivingEventsWrapper) ParseWebHook(b []byte) (interface{}, apperrors.AppError) {
	webhook, err := slackevents.ParseEvent(b, slackevents.OptionNoVerifyToken())
	if err != nil {
		return slackevents.EventsAPIEvent{}, apperrors.WrongInput("Failed to parse incomming slack payload into struct: %s", err)
	}
	return webhook, nil
}

//GetToken is a function that looks for the secret in the environment
func (wh receivingEventsWrapper) GetToken() string {
	return wh.secret
}
