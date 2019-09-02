package slack

import (
	"net/http"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"
	"github.com/nlopes/slack/slackevents"
)

//ReceivingEventsWrapper that bundles the github library functions into one struct with a Validator interface
type ReceivingEventsWrapper struct {
}

//ValidatePayload is a function used for checking whether the secret provided in the request is correct
func (wh ReceivingEventsWrapper) ValidatePayload(r *http.Request, b []byte) ([]byte, apperrors.AppError) {
	payload, err := slack.ValidatePayload(r, b)
	if err != nil {
		return nil, apperrors.AuthenticationFailed("Authentication during GitHub payload validation failed: %s", err)
	}
	return payload, nil
}

//ParseWebHook parses the raw json payload into an event struct
func (wh ReceivingEventsWrapper) ParseWebHook(b []byte) (interface{}, apperrors.AppError) {
	webhook, err := slackevents.ParseEvent(b)
	if err != nil {
		return slackevents.EventsAPIEvent{}, apperrors.WrongInput("Failed to parse incomming slack payload into struct: %s", err)
	}
	return webhook, nil
}

//GetToken is a function that looks for the secret in the environment
func (wh ReceivingEventsWrapper) GetToken() string {
	return os.Getenv("SLACK_SECRET")
}
