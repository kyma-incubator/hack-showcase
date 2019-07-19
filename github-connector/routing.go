package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

//Validator is an interface used to allow mocking the github library methods
type Validator interface {
	ValidatePayload(*http.Request, []byte) ([]byte, error)
	ParseWebHook(string, []byte) (interface{}, error)
	GetToken() string
}

//WebHookHandler is a struct
type WebHookHandler struct {
	validator Validator
}

//NewWebHookHandler creates a new webhook handler with the passed interface
func NewWebHookHandler(v Validator) *WebHookHandler {
	return &WebHookHandler{validator: v}
}

func (wh *WebHookHandler) handleWebhook(w http.ResponseWriter, r *http.Request) {

	payload, err := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	//payload, err := github.ValidatePayload(r, []byte("my-secret-key"))
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
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
	case *github.PushEvent:
		log.Printf("push")
	case *github.WatchEvent:
		log.Printf("%s is watching repo \"%s\"\n", e.GetSender().GetLogin(), e.GetRepo().GetFullName())
	case *github.StarEvent:
		// someone starred our repository
		if e.GetAction() == "created" {
			log.Printf("repository starred\n")
		} else if e.GetAction() == "deleted" {
			log.Printf("repository unstarred\n")
		}
	default:
		log.Printf("unknown event type: \"%s\"\n", github.WebHookType(r))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Index")
}
