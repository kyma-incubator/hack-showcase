package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/httperrors"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
)

//Validator is an interface used to allow mocking the github library methods
type Validator interface {
	ValidatePayload(*http.Request, []byte) ([]byte, apperrors.AppError)
	ParseWebHook(string, []byte) (interface{}, apperrors.AppError)
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

	switch e := event.(type) {
	case *github.IssuesEvent:
		if e.GetAction() == "opened" {
			apperr = wh.sender.SendToKyma("issuesevent.opened", "v1", "", os.Getenv("GITHUB_CONNECTOR_NAME")+"-app", payload)
		}
	case *github.PullRequestReviewEvent:
		if e.GetAction() == "submitted" {
			apperr = wh.sender.SendToKyma("pullrequestreviewevent.submitted", "v1", "", os.Getenv("GITHUB_CONNECTOR_NAME")+"-app", payload)
		}
	case *github.PushEvent:
		log.Infof("Push")
	case *github.WatchEvent:
		log.Infof("%s is watching repo '%s'.",
			e.GetSender().GetLogin(), e.GetRepo().GetFullName())
	case *github.StarEvent:
		if e.GetAction() == "created" {
			log.Infof("Repository starred.")
		} else if e.GetAction() == "deleted" {
			log.Infof("Repository unstarred.")
		}
	case *github.PingEvent:

	default:
		apperr := apperrors.NotFound("Unknown event type: '%s'", github.WebHookType(r))

		log.Warnf(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}
	if apperr != nil {
		log.Info(apperrors.Internal("While handling the event: %s", apperr.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
