package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/eventparser"

	"github.com/google/go-github/github"
)

//Validator is an interface used to allow mocking the github library methods
type Validator interface {
	ValidatePayload(*http.Request, []byte) ([]byte, error)
	ParseWebHook(string, []byte) (interface{}, error)
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

func sendToKyma(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) {
	client := http.Client{}
	toSend, err := eventparser.GetEventRequestPayload(eventType, eventTypeVersion, eventID, sourceID,
		data)
	if err != nil {
		log.Printf("not converted")
	}

	jsonToSend, err := eventparser.GetEventRequestAsJSON(toSend)
	kymaRequest, _ := http.NewRequest(http.MethodPost, "http://event-bus-publish.kyma-system:8080/v1/events",
		bytes.NewReader(jsonToSend))
	response, err := client.Do(kymaRequest)
	log.Println(response)
	log.Println(err)
}

//HandleWebhook is a function that handles the /webhook endpoint.
func (wh *WebHookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {

	payload, err := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		log.Printf("request body: %s\n", r.Body)
		w.WriteHeader(http.StatusUnauthorized)

		return
	}
	defer r.Body.Close()

	event, err := wh.validator.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch e := event.(type) {
	case *github.IssuesEvent:

		sendToKyma("issuesevent.opened", "v1", "", "github-connector-app", payload)

	case *github.PullRequestReviewEvent:
		if e.GetAction() == "submitted" {
			sendToKyma("pullrequestreviewevent.submitted", "v1", "", "github-connector-app", payload)
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
