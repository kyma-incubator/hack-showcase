package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/githubwrappers"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"
)

func main() {

	log.Println("server started")
	webhook := handlers.NewWebHookHandler(
		githubwrappers.ReceivingEventsWrapper{},
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
