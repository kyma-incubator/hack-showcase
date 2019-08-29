package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/httperrors"
	"github.com/nlopes/slack/slackevents"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"

	log "github.com/sirupsen/logrus"
)

//Validator is an interface used to allow mocking the github library methods
type Validator interface {
	ValidatePayload(*http.Request, []byte) ([]byte, apperrors.AppError)
	ParseWebHook([]byte) (interface{}, apperrors.AppError)
	GetToken() string
}

//Sender is an interface used to allow mocking sending events to Kyma's event bus
type Sender interface {
	SendToKyma(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) apperrors.AppError
}

//WebHookHandler is a struct used to allow mocking the github library methods
type WebHookHandler struct {
	validator Validator
	sender    Sender
}

//NewWebHookHandler creates a new webhook handler with the passed interface
func NewWebHookHandler(v Validator, s Sender) *WebHookHandler {
	return &WebHookHandler{validator: v, sender: s}
}

//HandleWebhook is a function that handles the /webhook endpoint.
func (wh *WebHookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	payload, apperr := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}

	event, apperr := wh.validator.ParseWebHook(payload)
	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}

	eventType := reflect.Indirect(reflect.ValueOf(event)).Type().Name()

	if eventType == "URLVerification" {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	sourceID := fmt.Sprintf("%s-app", os.Getenv("SLACK_CONNECTOR_NAME"))
	log.Info(eventType)
	apperr = wh.sender.SendToKyma(eventType, "v1", "", sourceID, payload)

	if apperr != nil {
		log.Info(apperrors.Internal("While handling the event: %s", apperr.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
