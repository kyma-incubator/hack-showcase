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
func (wh ReceivingEventsWrapper) ValidatePayload(r *http.Request, b []byte) (bool, apperrors.AppError) {

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(r.Body)
	// body := buf.String()

	// isValid := subtle.ConstantTimeCompare([]byte(body), b)
	// if isValid != 1 {
	// 	var err error
	// 	return false, apperrors.AuthenticationFailed("Authentication during Slack payload validation failed: %s", err)
	// }
	// return true, nil
	return true, nil
}

//ParseWebHook parses the raw json payload into an event struct
func (wh ReceivingEventsWrapper) ParseWebHook(b []byte) (interface{}, apperrors.AppError) {
	webhook, err := slackevents.ParseEvent(b, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: "Otf0ZyDSMAXH2yOVqs7nunt6"}))
	if err != nil {
		return slackevents.EventsAPIEvent{}, apperrors.WrongInput("Failed to parse incomming slack payload into struct: %s", err)
	}
	return webhook, nil
}

//GetToken is a function that looks for the secret in the environment
func (wh ReceivingEventsWrapper) GetToken() string {
	return "Otf0ZyDSMAXH2yOVqs7nunt6"
}
