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
	log.Info("server started")

	id, err := registerservice.RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

	webhook := handlers.NewWebHookHandler(
		githubwrappers.ReceivingEventsWrapper{},
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
