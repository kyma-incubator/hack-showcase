package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte("my-secret-key"))
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	log.Println(event)

	switch e := event.(type) {
	case *github.WatchEvent:
		log.Printf(e.GetSender().GetLogin() + " has started watching your repo")
	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "testy")
}
