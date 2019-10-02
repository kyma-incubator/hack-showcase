package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/google/go-github/github"
	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/httperrors"

	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/apperrors"
	git "github.com/kyma-incubator/github-slack-connectors/github-connector/internal/github"
	log "github.com/sirupsen/logrus"
)

//Sender is an interface used to allow mocking sending events to Kyma's event bus
type Sender interface {
	SendToKyma(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) apperrors.AppError
}

//WebHookHandler is a struct used to allow mocking the github library methods
type WebHookHandler struct {
	validator git.Validator
	sender    Sender
}

//NewWebHookHandler creates a new webhook handler with the passed interface
func NewWebHookHandler(v git.Validator, s Sender) *WebHookHandler {
	return &WebHookHandler{validator: v, sender: s}
}

//HandleWebhook is a function that handles the /webhook endpoint.
func (wh *WebHookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload, apperr := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}

	event, apperr := wh.validator.ParseWebHook(github.WebHookType(r), payload)
	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}

	eventType := reflect.Indirect(reflect.ValueOf(event)).Type().Name()
	sourceID := fmt.Sprintf("%s-app", os.Getenv("GITHUB_CONNECTOR_NAME"))
	log.Info(fmt.Sprintf("Event type '%s' received.", eventType))
	apperr = wh.sender.SendToKyma(eventType, "v1", "", sourceID, payload)

	if apperr != nil {
		log.Info(apperrors.Internal("While handling the event: %s", apperr.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
