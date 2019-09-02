package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/httperrors"
	"github.com/nlopes/slack/slackevents"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/slack"
	log "github.com/sirupsen/logrus"
)

//Sender is an interface used to allow mocking sending events to Kyma's event bus
type Sender interface {
	SendToKyma(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) apperrors.AppError
}

//WebHookHandler is a struct used to allow mocking the github library methods
type WebHookHandler struct {
	validator slack.Validator
	sender    Sender
}

//NewWebHookHandler creates a new webhook handler with the passed interface
func NewWebHookHandler(v slack.Validator, s Sender) *WebHookHandler {
	return &WebHookHandler{validator: v, sender: s}
}

//HandleWebhook is a function that handles the /webhook endpoint.
func (wh *WebHookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	payload := buf.String()

	_, apperr := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}
	event, apperr := wh.validator.ParseWebHook([]byte(payload))

	log.Info(event)
	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}
	var replacer = strings.NewReplacer("_", ".")
	eventType := event.(slackevents.EventsAPIEvent).InnerEvent.Type //e.g.: "member_joined_channel"
	withDots := replacer.Replace(eventType)
	log.Info(withDots)
	// member_joined_channel -> (	slack.events.member.joined.channel) => sendToKyma()
	eventTypeToKyma := fmt.Sprintf("slack.events.%s", withDots)
	log.Info(eventTypeToKyma)

	log.Info(eventType)
	if event.(slackevents.EventsAPIEvent).Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		log.Info([]byte(payload))
		err := json.Unmarshal([]byte(payload), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(r.Challenge))
	}

	sourceID := fmt.Sprintf("%s-app", os.Getenv("SLACK_CONNECTOR_NAME"))
	log.Info(eventType)
	apperr = wh.sender.SendToKyma(withDots, "v1", "", sourceID, []byte(payload))

	if apperr != nil {
		log.Info(apperrors.Internal("While handling the event: %s", apperr.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
