package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/httperrors"

	"github.com/google/go-github/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	log "github.com/sirupsen/logrus"
)

//Validator is an interface used to allow mocking the github library methods
type Validator interface {
	ValidatePayload(*http.Request, []byte) ([]byte, apperrors.AppError)
	ParseWebHook(string, []byte) (interface{}, apperrors.AppError)
	GetToken() string
}

//WebHookHandler is a struct used to allow mocking the github library methods
type WebHookHandler struct {
	validator Validator
}

//NewWebHookHandler creates a new webhook handler with the passed interface
func NewWebHookHandler(v Validator) *WebHookHandler {
	return &WebHookHandler{validator: v}
}

//HandleWebhook is a function that handles the /webhook endpoint.
func (wh *WebHookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {

	payload, apperr := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if apperr != nil {
		apperr.Append("While handling '/webhook' endpoint")
		log.Warn(apperr.Error())
		httpcode, resp := httperrors.AppErrorToResponse(apperr, false)
		w.WriteHeader(httpcode)
		body, err := json.Marshal(resp)
		if err != nil {
			errjson := apperrors.Internal("Failed to marshal json response: %s", err)
			log.Warn(errjson)
			return
		}
		w.Write(body)
		return
	}
	defer r.Body.Close()

	event, apperr := wh.validator.ParseWebHook(github.WebHookType(r), payload)
	if apperr != nil {
		apperr.Append("While handling '/webhook' endpoint")
		log.Warn(apperr.Error())
		httpcode, resp := httperrors.AppErrorToResponse(apperr, false)
		w.WriteHeader(httpcode)
		body, err := json.Marshal(resp)
		if err != nil {
			errjson := apperrors.Internal("Failed to marshal json response: %s", err)
			log.Warn(errjson)
			return
		}
		w.Write(body)
		return
	}

	switch e := event.(type) {
	case *github.IssuesEvent:
		log.Infof("%s has opened an issue: '%s'.",
			e.GetSender().GetLogin(), e.GetIssue().GetTitle())

	case *github.PullRequestReviewEvent:
		if e.GetAction() == "submitted" {
			log.Infof("%s has submitted a review on pull request: '%s'.",
				e.GetSender().GetLogin(), e.GetPullRequest().GetTitle())
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
		httpcode, resp := httperrors.AppErrorToResponse(apperr, false)
		w.WriteHeader(httpcode)
		body, err := json.Marshal(resp)
		if err != nil {
			errjson := apperrors.Internal("Failed to marshal json response: %s", err)
			log.Warn(errjson)
			return
		}
		w.Write(body)
		return
	}
	w.WriteHeader(http.StatusOK)
}
