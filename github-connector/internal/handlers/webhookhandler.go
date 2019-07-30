package handlers

import (
	"net/http"

	"github.com/google/go-github/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/httperrors"
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

	payload, appError := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if appError != nil {
		log.Printf("%v", appError) //TODO
		w.WriteHeader(http.StatusUnauthorized)

		return
	}
	defer r.Body.Close()

	event, err := wh.validator.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Warn(httperrors.AppErrorToResponse(appError, false)) //TODO
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch e := event.(type) {
	case *github.IssuesEvent:

		log.Printf("%s has opened an issue: \"%s\"",
			e.GetSender().GetLogin(), e.GetIssue().GetTitle())

	case *github.PullRequestReviewEvent:
		if e.GetAction() == "submitted" {
			log.Printf("%s has submitted a review on pull request: \"%s\"",
				e.GetSender().GetLogin(), e.GetPullRequest().GetTitle())
		}
	case *github.PushEvent:
		log.Printf("push")
	case *github.WatchEvent:
		log.Printf("%s is watching repo \"%s\"\n",
			e.GetSender().GetLogin(), e.GetRepo().GetFullName())
	case *github.StarEvent:
		if e.GetAction() == "created" {
			log.Printf("repository starred\n")
		} else if e.GetAction() == "deleted" {
			log.Printf("repository unstarred\n")
		}
	case *github.PingEvent:

	default:
		log.Printf("unknown event type: \"%s\"\n", github.WebHookType(r))

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
