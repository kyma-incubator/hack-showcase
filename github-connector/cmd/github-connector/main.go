package main

import (
	"net/http"
	"os"

	registerservice "github.com/kyma-incubator/hack-showcase/github-connector/internal/application_registry"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/githubwrappers"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("server started")
	registerservice.RegisterService()

	webhook := handlers.NewWebHookHandler(
		githubwrappers.ReceivingEventsWrapper{},
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
